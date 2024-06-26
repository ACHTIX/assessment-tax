package util

import (
	"errors"
	"github.com/ACHTIX/assessment-tax/model"
	"log"
)

// IncomeCalculation รับเงินทั้งหมดมาเพื่อหักค่าลดหย่อนภาษี = เงินสุทธิ netIncome
func IncomeCalculation(input model.TaxInput) float64 {
	var personal float64

	amount := model.Allowance{Amount: CheckAllowanceType(input.Allowances)}
	allowance := model.Allowance{AllowanceType: "Personal"}

	if allowance.Amount <= 0 {
		personal = 60000
	} else if allowance.Amount < 10000 {
		personal = 10000
	} else {
		personal = allowance.Amount
	}

	allowanceTotal := personal + amount.Amount      //ยอดรวมลดหย่อนภาษี
	netIncome := input.TotalIncome - allowanceTotal //เงินทั้งหมดหักค่าลดหย่อนแล้ว

	log.Println("netIncome", netIncome)

	return netIncome
}

// TaxCalculation คำนวนภาษีที่ต้องจ่าย (หักwhtแล้ว)
func TaxCalculation(input model.TaxInput) (float64, model.TaxLevel, error) {
	netIncome := IncomeCalculation(input)               //จาก IncomeCalculation(เงินที่หักค่าลดหย่อนแล้ว)
	deduction, taxlevel := DeductionTaxLevel(netIncome) //จาก DeductionTaxLevel(ภาษีที่ต้องจ่าย)
	subWHT := deduction - input.Wht

	if netIncome < 0 {
		return 0, model.TaxLevel{}, errors.New("\"Key: 'TaxInput.TotalIncome' Error:Field validation for 'TotalIncome' failed on the 'gte' tag\"")
	}

	return subWHT, taxlevel, nil
}

// ระบุเปอร์เซ้นที่ต้องนำมาคำนวน
func TaxRateLevel(netIncome float64) float64 {
	if netIncome < 0 {
		return -1 // Error case for negative netIncome
	} else if netIncome <= 150000 {
		return 0 // No tax for netIncome ≤ 150,000
	} else if netIncome <= 500000 {
		return 0.10 // 10% tax for netIncome ≤ 500,000
	} else if netIncome <= 1000000 {
		return 0.15 // 15% tax for netIncome ≤ 1,000,000
	} else if netIncome <= 2000000 {
		return 0.20 // 20% tax for netIncome ≤ 2,000,000
	} else if netIncome >= 2000001 {
		return 0.35 // 35% tax for netIncome > 2,000,000
	}

	log.Println("netIncome", netIncome)

	return 0
}

// คำนวนภาษีก่อนหักwht
func DeductionTaxLevel(netIncome float64) (tax float64, level model.TaxLevel) {
	taxRateStr := TaxRateLevel(netIncome)

	if netIncome < 0 && netIncome != 0 {
		return 0, model.TaxLevel{} // Error case for negative netIncome
	} else if netIncome <= 150000 || netIncome == 0 {
		return 0, model.TaxLevel{Level: "0-150,000", Tax: 0.0} // No tax for netIncome ≤ 150,000
	} else if netIncome <= 500000 {
		sub := netIncome - 150000
		total := sub * taxRateStr
		return total, model.TaxLevel{Level: "150,001-500,000", Tax: total}
	} else if netIncome <= 1000000 {
		sub := netIncome - 500000
		total := (sub * taxRateStr) + 35000
		return total, model.TaxLevel{Level: "500,001-1,000,000", Tax: total}
	} else if netIncome <= 2000000 {
		sub := netIncome - 1000000
		total := (sub * taxRateStr) + 110000
		return total, model.TaxLevel{Level: "1,000,001-2,000,000", Tax: total}
	} else {
		sub := netIncome - 2000000
		total := (sub * taxRateStr) + 310000
		return total, model.TaxLevel{Level: "2,000,001 ขึ้นไป", Tax: total}
	}
}
