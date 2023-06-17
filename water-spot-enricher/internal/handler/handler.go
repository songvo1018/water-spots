package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type WaterSpotData struct {
	Identifier    uint8   `json:"identifier,omitempty"`
	Model         string  `json:"model,omitempty"`
	LastService   string  `json:"last_service,omitempty"`
	Owner         string  `json:"owner,omitempty"`
	IsUnsecured   bool    `json:"is_unsecured,omitempty"`
	OzonPressure  float32 `json:"ozon_pressure,omitempty"`
	WaterPressure float32 `json:"water_pressure,omitempty"`
}

func Handle(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newIncident WaterSpotData

		if err := c.BindJSON(&newIncident); err != nil {
			log.Printf("An error occured: %v", err)
			return
		}

		fmt.Println("Incident caught")
		save(&newIncident, db)
	}
}

func save(newIncident *WaterSpotData, db *sql.DB) {
	fmt.Printf("saving: %v", newIncident)
	// todo data validation

	if newIncident.Identifier > 0 {
		insertString := `insert into "water-spots"
							("identifier", "Model", "LastService", "Owner", "IsUnsecured", "OzonPressure","WaterPressure") 
							values ($1, $2, $3, $4, $5, $6, $7)`
		result, err := db.Exec(insertString, newIncident.Identifier, newIncident.Model, newIncident.LastService, newIncident.Owner, newIncident.IsUnsecured, newIncident.OzonPressure, newIncident.WaterPressure)
		if err != nil {
			log.Printf("An error occured while executing query: %v", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}

		fmt.Println("result of saving ok:", rowsAffected > 0)
	}

}
