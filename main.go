package main

import (
	"github.com/ACHTIX/assessment-tax/database"
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func handleTaxCalculation(c echo.Context) error {
	var taxInput model.TaxInput

	// Bind the incoming JSON to the struct
	if err := c.Bind(&taxInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Ensure there is at least one allowance to process
	if len(taxInput.Allowances) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No allowances provided"})
	}

	// Use the first allowance for the calculation
	firstAllowance := taxInput.Allowances[0]

	// Calculate the tax based on the provided input and the first allowance
	taxResult := util.TaxCalculation(taxInput, firstAllowance)

	// Format the result into JSON and return it
	return c.JSON(http.StatusOK, map[string]float64{"tax": taxResult})
}

func main() {
	e := echo.New()

	e.POST("/tax/calculations", handleTaxCalculation)

	// Initialize the database and connect
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	// Start the Echo web server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
