package util

import "github.com/ACHTIX/assessment-tax/model"

// รับเงินทั้งหมดมาเพื่อหักค่าลดหย่อนภาษี = เงินสุทธิ netIncome
func IncomeCalculation(input model.TaxInput, allowance model.Allowance) float64 {
	allowanceTotal := 60000 + checkAllowanceType(allowance.AllowanceType, allowance.Amount) //ยอดรวมลดหย่อนภาษี
	netIncome := input.TotalIncome - allowanceTotal                                         //เงินทั้งหมดหักค่าลดหย่อนแล้ว
	return netIncome
}

// คำนวนภาษีที่ต้องจ่าย (หักwhtแล้ว)
func TaxCalculation(input model.TaxInput, allowance model.Allowance) float64 {
	netIncome := IncomeCalculation(input, allowance) //จาก IncomeCalculation(เงินที่หักค่าลดหย่อนแล้ว)
	deduction := DeductionTaxLevel(netIncome)        //จาก DeductionTaxLevel(ภาษีที่ต้องจ่าย)

	subWHT := deduction - input.Wht

	return subWHT
}

// ระบุเปอร์เซ้นที่ต้องนำมาคำนวน
func taxRateLevel(netIncome float64) float64 {
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
	} else if netIncome > 2000001 {
		return 0.35 // 35% tax for netIncome > 2,000,000
	}
	return 0
}

// คำนวนภาษีก่อนหักwht
func DeductionTaxLevel(netIncome float64) float64 {
	taxRateStr := taxRateLevel(netIncome)

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
		total := (sub * taxRateStr) + 35000
		return total
	} else if netIncome <= 2000000 {
		sub := netIncome - 1000000
		total := (sub * taxRateStr) + 110000
		return total
	} else if netIncome > 2000000 {
		sub := netIncome - 2000000
		total := (sub * taxRateStr) + 310000
		return total
	}
	return 0
}

// ตรวจสอบขั้นของภาษี
func checkTaxLevel(netIncome float64) string {
	switch {
	case netIncome < 0:
		return "Error: negative income"
	case netIncome <= 150000:
		return "0 - 150,000"
	case netIncome <= 500000:
		return "150,001 - 500,000"
	case netIncome <= 1000000:
		return "500,001 - 1,000,000"
	case netIncome <= 2000000:
		return "1,000,001 - 2,000,000"
	default:
		return "2,000,001 and up"
	}
}
