package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	dbHandle := db.getDatabase()
	router := gin.Default()

	router.Run("localhost:8080")
}
