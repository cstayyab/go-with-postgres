package helpers

import (
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name string `gorm:"column:company_name"`
}

func (Company) TableName() string {
	return "customer_companies"
}

type Customer struct {
	gorm.Model
	Username    string         `gorm:"unique;not null;column:login"`
	Password    string         `gorm:"column:password"`
	Name        string         `gorm:"column:name"`
	Company     Company        `gorm:"embedded;column:company_id"`
	CreditCards pq.StringArray `gorm:"type:text[];column:credit_cards"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Created  int64    `gorm:"autoCreateTime;column:created_at"`
	Name     string   `gorm:"column:order_name"`
	Customer Customer `gorm:"embedded;column:customer_id"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	gorm.Model
	Order       Order   `gorm:"embedded;column:order_id"`
	PPU         float64 `gorm:"type:decimal(9,4);column:price_per_unit"`
	Quantity    uint    `gorm:"column:quantity"`
	ProductName string  `gorm:"column:product"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

type Delivery struct {
	gorm.Model
	OrderItem OrderItem `gorm:"embedded;column:order_item_id"`
	Quantity  uint      `gorm:"column:delivered_quantity"`
}

func (Delivery) TableName() string {
	return "deliveries"
}

func GetDBConnection() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Australia/Melbourne",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else if db != nil {
		return db
	}
	return nil
}
