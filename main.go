package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/ACHTIX/assessment-tax/database"
	_ "github.com/ACHTIX/assessment-tax/docs"
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func handleTaxCalculation(c echo.Context) error {
	var taxInput model.TaxInput

	// Bind the incoming JSON to the struct
	if err := c.Bind(&taxInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input data"})
	}
	if err := c.Validate(taxInput); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Ensure there is at least one allowance to process
	if len(taxInput.Allowances) <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No allowances provided"})
	}

	taxResult, taxLevel, _ := util.TaxCalculation(taxInput)

	log.Println("taxResult", taxResult, "taxLevel", taxLevel)

	taxData := model.TaxData{}

	if taxResult >= 0 {
		taxData = model.TaxData{Tax: taxResult}
	} else {
		taxData = model.TaxData{TaxRefund: math.Abs(taxResult)}
	}

	taxData.TaxLevel = []model.TaxLevel{
		{Level: "0-150,000", Tax: 0},
		{Level: "150,001-500,000", Tax: 0},
		{Level: "500,001-1,000,000", Tax: 0},
		{Level: "1,000,001-2,000,000", Tax: 0},
		{Level: "2,000,001 ขึ้นไป", Tax: 0},
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

	log.Println("taxData", taxData)

	// Return the tax JSON
	return c.JSON(http.StatusOK, taxData)
}

func handleAdminDeductionPersonal(c echo.Context) error {
	var input model.AdminRequestStruct

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
	if err := c.Validate(input); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Calculate the personal deduction based on the fetched allowance
	result := util.AllowancePersonalAdmin(input.Amount)

	err = database.SetPersonalAllowanceDB(result)
	if err != nil {
		panic(err)
	}

	log.Println("result", result)

	// Return the tax JSON
	return c.JSON(http.StatusOK, map[string]float64{"personalDeduction": result})
}

func handleUpload(c echo.Context) error {
	file, err := c.FormFile("taxFile")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get the file")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	reader := csv.NewReader(src)
	var taxes []model.TaxResponseCSVDataStruct

	if _, err = reader.Read(); err != nil {
		return err
	}
	var loopNumber = 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read the csv file")
		}

		totalIncome, err := strconv.ParseFloat(record[0], 64)
		wht, err := strconv.ParseFloat(record[1], 64)
		donation, err := strconv.ParseFloat(record[2], 64)

		log.Println("totalIncome", totalIncome, "wht", wht, "donation", donation)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid totalIncome value")
		}

		data := model.TaxInput{
			TotalIncome: totalIncome,
			Wht:         wht,
			Allowances: []model.Allowance{
				{
					AllowanceType: "donation",
					Amount:        donation,
				},
			},
		}

		log.Println("data", data)

		tax, _, _ := util.TaxCalculation(data)

		log.Println("tax", tax)

		if tax >= 0 {
			taxes = append(
				taxes,
				model.TaxResponseCSVDataStruct{
					TotalIncome: totalIncome,
					Tax:         tax,
				},
			)
		} else {
			taxes = append(
				taxes,
				model.TaxResponseCSVDataStruct{
					TotalIncome: totalIncome,
					TaxRefund:   math.Abs(tax),
				},
			)
		}

		loopNumber += 1
	}

	log.Println("taxes", taxes)

	return c.JSON(http.StatusOK, model.TaxResponseCSVStruct{Taxes: taxes})
}

func handleAdminDeductionKReceipt(c echo.Context) error {

	var input model.AdminRequestStruct

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
	if err := c.Validate(input); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Calculate the personal deduction based on the fetched allowance
	result := util.AllowanceKReceiptAdmin(input.Amount)

	log.Println("result", result)

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
		log.Println("DATABASE_URL", DATABASE_URL)
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Default port if not specified
		log.Printf("Defaulting to port %s", PORT)
	}
	ADMIN_USERNAME := os.Getenv("ADMIN_USERNAME")
	if ADMIN_USERNAME == "" {
		ADMIN_USERNAME = "adminTax" // Default port if not specified
		log.Printf("Defaulting to admin user %s", ADMIN_USERNAME)
	}
	ADMIN_PASSWORD := os.Getenv("ADMIN_PASSWORD")
	if ADMIN_PASSWORD == "" {
		ADMIN_PASSWORD = "admin!" // Default port if not specified
		log.Printf("Defaulting to admin password %s", ADMIN_PASSWORD)
	}

	database.InitDB(DATABASE_URL)

	e.POST("/tax/calculations", handleTaxCalculation)

	e.POST("/admin/deductions/personal", handleAdminDeductionPersonal)

	e.POST("/tax/calculations/upload-csv", handleUpload)

	e.POST("/admin/deductions/k-receipt", handleAdminDeductionKReceipt)

	e.GET("/swagger/*", echoswagger.WrapHandler)

	e.Validator = &util.CustomValidator{Validator: validator.New()}

	//// Start the Echo web server on port 8080
	go func() {
		if err := e.Start(":" + PORT); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	preShutdownTasks()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("error shutting down the server: ", err)
	}
	fmt.Println("Server has been shut down.")
}

func preShutdownTasks() {
	fmt.Println("Running pre-shutdown tasks...")
	database.CloseDB()
	time.Sleep(5 * time.Second)
	fmt.Println("Pre-shutdown tasks complete.")
}
