package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"water-spot-enricher/internal/handler"
)

func main() {
	router := gin.Default()

	connection := getConnection()
	router.POST("/incidents", handler.Handle(connection))

	router.Run("localhost:8070")
}

func getConnection() *sql.DB {
	connStr := "postgresql://pguser:pgpassword@localhost:5432/water-spots?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
