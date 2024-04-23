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
	if len(taxInput.Allowances) == 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	if len(taxInput.Allowances) == 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	taxResult, taxLevel := util.TaxCalculation(taxInput)

	taxData := model.TaxData{
		Tax: taxResult,
		TaxLevel: []model.TaxLevel{
			{Level: "0-150,000", Tax: 0},
			{Level: "150,001-500,000", Tax: 0},
			{Level: "500,001-1,000,000", Tax: 0},
			{Level: "1,000,001-2,000,000", Tax: 0},
			{Level: "2,000,001 ขึ้นไป", Tax: 0},
		},
	}

	for i := range taxData.TaxLevel {
		if taxData.TaxLevel[i].Level == taxLevel.Level {
			taxData.TaxLevel[i].Tax = taxLevel.Tax
			break
		} else if taxData.TaxLevel[i].Level == "0-150,000" {
			taxData.TaxLevel[i].Tax = 0
		} else if taxData.TaxLevel[i].Level == "150,001-500,000" {
			taxData.TaxLevel[i].Tax = 35000.0
		} else if taxData.TaxLevel[i].Level == "500,001-1,000,000" {
			taxData.TaxLevel[i].Tax = 110000.0
		} else if taxData.TaxLevel[i].Level == "1,000,001-2,000,000" {
			taxData.TaxLevel[i].Tax = 310000.0
		} else if taxData.TaxLevel[i].Level == "2,000,001 ขึ้นไป" {
			taxData.TaxLevel[i].Tax = taxResult
		}
	}

	// Return the tax JSON
	return c.JSON(http.StatusOK, taxData)
}

func SeparateTaxLevel(netIncome float64, threshold float64) float64 {
	taxRateStr := util.TaxRateLevel(netIncome)

	if netIncome < 0 {
		return -1 // Error case for negative netIncome
	} else if netIncome <= 150000 {
		return 0 // No tax for netIncome ≤ 150,000
	} else if netIncome <= 500000 {
		sub := netIncome - 150000
		total := sub * taxRateStr
		return total
	} else if netIncome <= 1000000 {
		sub := netIncome - 500000
		total := sub * taxRateStr
		return total
	} else if netIncome <= 2000000 {
		sub := netIncome - 1000000
		total := sub * taxRateStr
		return total
	} else if netIncome > 2000000 {
		sub := netIncome - 2000000
		total := sub * taxRateStr
		return total
	}
	return 0
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
