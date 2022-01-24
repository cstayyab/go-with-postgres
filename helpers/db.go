package helpers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBConnection() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Australia/Melbourne",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else if db != nil {
		return db
	}
	return nil
}
