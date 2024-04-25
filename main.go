package main

import (
	"github.com/ACHTIX/assessment-tax/database"
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func handleTaxCalculation(c echo.Context) error {
	var taxInput model.TaxInput

	// Bind the incoming JSON to the struct
	if err := c.Bind(&taxInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Ensure there is at least one allowance to process
	if len(taxInput.Allowances) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No allowances provided"})
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

func handleAdminDeductionPersonal(c echo.Context) error {
	var input model.Allowance

	// Fetch from the database
	allowance, err := database.GetAllowance()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching users")
	}

	// Check if the allowance is nil
	if allowance == nil {
		return c.JSON(http.StatusInternalServerError, "No allowance found in the database")
	}

	// Bind the incoming JSON to the struct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Calculate the personal deduction based on the fetched allowance
	result := util.AllowancePersonalAdmin(input.Amount)

	err = database.SetPersonalAllowanceDB(result)
	if err != nil {
		panic(err)
	}

	// Return the tax JSON
	return c.JSON(http.StatusOK, map[string]float64{"personalDeduction": result})
}

func handleUpload(echo.Context) error {

	return nil
}

func handleAdminDeductionKReceipt(c echo.Context) error {

	var input model.Allowance

	// Fetch from the database
	allowance, err := database.GetAllowance()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching users")
	}

	// Check if the allowance is nil
	if allowance == nil {
		return c.JSON(http.StatusInternalServerError, "No allowance found in the database")
	}

	// Bind the incoming JSON to the struct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}

	// Calculate the personal deduction based on the fetched allowance
	result := util.AllowanceKReceiptAdmin(input.Amount)

	err = database.SetKReceiptAllowanceDB(result)
	if err != nil {
		panic(err)
	}

	// Return the tax JSON
	return c.JSON(http.StatusOK, map[string]float64{"kReceipt": result})
}

func main() {
	e := echo.New()

	// Use environment variables for database connection
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable" // Default port if not specified
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Default port if not specified
	}
	ADMIN_USERNAME := os.Getenv("ADMIN_USERNAME")
	if ADMIN_USERNAME == "" {
		ADMIN_USERNAME = "adminTax" // Default port if not specified
	}
	ADMIN_PASSWORD := os.Getenv("ADMIN_PASSWORD")
	if ADMIN_PASSWORD == "" {
		ADMIN_PASSWORD = "admin!" // Default port if not specified
	}

	database.InitDB(DATABASE_URL)

	e.POST("/tax/calculations", handleTaxCalculation)

	e.POST("/admin/deductions/personal", handleAdminDeductionPersonal)

	e.POST("tax/calculations/upload - csv", handleUpload)

	e.POST("/admin/deductions/k-receipt", handleAdminDeductionKReceipt)

	e.Validator = &util.CustomValidator{Validator: validator.New()}

	//// Start the Echo web server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
