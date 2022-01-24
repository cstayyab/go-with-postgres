package main

import (
	"github.com/cstayyab/go-with-posgres/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := helpers.GetDBConnection()
	db.AutoMigrate()
	router := gin.Default()

	router.Run("localhost:8080")
}
