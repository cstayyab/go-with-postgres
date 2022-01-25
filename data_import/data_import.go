package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cstayyab/go-with-posgres/helpers"
	"github.com/jszwec/csvutil"
)

type CompanyCSV struct {
	ID   uint   `csv:"company_id"`
	Name string `csv:"company_name"`
}

type CustomerCSV struct {
	ID          string `csv:"user_id"`
	Username    string `csv:"login"`
	Password    string `csv:"password"`
	Name        string `csv:"name"`
	CompanyID   string `csv:"company_id"`
	CreditCards string `csv:"credit_cards"`
}

func main() {
	db := helpers.GetDBConnection()
	if db == nil {
		panic("Cannot Connect to DB")
	}
	fmt.Println("Importing Models into Database . . .")
	db.AutoMigrate(&helpers.Company{}, &helpers.Customer{}, &helpers.Order{}, &helpers.OrderItem{}, &helpers.Delivery{})

	fmt.Println("Importing Data . . .")

	fmt.Println("Reading customer_companies.csv . . .")
	companiesFile, err := os.ReadFile("customer_companies.csv")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v", companiesFile)

	var companies []CompanyCSV

	if err := csvutil.Unmarshal(companiesFile, &companies); err != nil {
		panic(err)
	}
	for _, company := range companies {
		fmt.Printf("%+v\n", company)
		companyRec := helpers.Company{ID: company.ID, Name: company.Name}
		if result := db.Create(&companyRec); result.Error != nil {
			panic(result.Error)
		} else {
			fmt.Println("Company Inserted: " + company.Name)
		}
	}

	fmt.Println("Reading customers.csv . . .")
	customersFile, err := os.ReadFile("customers.csv")
	if err != nil {
		panic(err)
	}

	var customers []CustomerCSV

	if err := csvutil.Unmarshal(customersFile, &customers); err != nil {
		panic(err)
	}

	for _, customer := range customers {
		var creditCardsJSON []string
		_ = json.Unmarshal([]byte(customer.CreditCards), &creditCardsJSON)
		var company helpers.Company
		db.First(&company, customer.CompanyID)
		customerRec := helpers.Customer{ID: customer.ID, Username: customer.Username, Name: customer.Name, Password: customer.Password, CreditCards: creditCardsJSON, Company: company}
		result := db.Create(&customerRec)
		if result.Error != nil {
			fmt.Println("Customer Inserted: " + company.Name + "=>" + customerRec.Name)
		} else {
			panic(result.Error)
		}
	}
}
