package main

import (
	"github.com/ACHTIX/assessment-tax/database"
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	e.POST("/tax/calculations", func(c echo.Context) error {
		var taxInput model.TaxInput

		// Bind the incoming JSON to the struct
		if err := c.Bind(&taxInput); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
		}

		// Check if there are any allowances
		if len(taxInput.Allowances) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "No allowances provided"})
		}

		firstAllowance := taxInput.Allowances[0] // Use the first allowance

		// Calculate tax based on deductions
		result := util.TaxCalculation(taxInput, firstAllowance)

		// Respond with the calculated tax data
		return c.JSON(http.StatusOK, map[string]float64{"tax": result})
	})

	// Initialize the database and connect
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	// Start the Echo web server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
