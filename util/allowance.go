package util

import (
	"github.com/ACHTIX/assessment-tax/model"
	"log"
)

// แยกประเภทค่าลดหย่อน
func CheckAllowanceType(input []model.Allowance) float64 {
	amount := 0.0
	for i := range input {
		allowance := input[i].Amount
		if input[i].AllowanceType == "donation" {
			amount += DonationAllowance(allowance)
		} else if input[i].AllowanceType == "k-receipt" {
			amount += KReceiptAllowance(allowance)
		}
	}

	log.Println("CheckAllowanceType amount", amount)

	return amount
}

// ค่าลดหย่อนภาษีจากการบริจาค ไม่เกิน 100,000
func DonationAllowance(amount float64) float64 {
	if amount >= 100001 {
		return 100000
	} else if amount <= 100000 && amount != 0 {
		return amount
	} else {
		log.Println("DonationAllowance amount", amount)
		return 0
	}
}

// ค่าลด k-receipt ต้องมีค่ามากกว่า 0 บาท
// ค่าลดหย่อนภาษีจาก k-receipt สูงสุดไม่เกิน 50,000
func KReceiptAllowance(amount float64) float64 {
	if amount >= 50001 {
		return 50000
	} else if amount <= 50000 && amount != 0 {
		return amount
	}

	log.Println("KReceiptAllowance amount", amount)

	return 0
}

// ค่าลดหย่อนส่วนตัวต้องมีค่ามากกว่า 10,000 บาท
// แอดมินสามารถกำหนดค่าลดหย่อนส่วนตัวได้โดยไม่เกิน 100,000
func AllowancePersonalAdmin(amount float64) float64 {
	if amount >= 100001 {
		return 100000
	} else if amount <= 10000 {
		return 10000
	} else {
		log.Println("AllowancePersonalAdmin amount", amount)
		return amount
	}
}

// ค่าลดหย่อนภาษีจาก k-receipt แอดมินกำหมดได้ สูงสุดไม่เกิน 100,000
func AllowanceKReceiptAdmin(amount float64) float64 {
	if amount >= 100001 {
		return 100000
	} else if amount < 0 {
		return 0
	} else {
		log.Println("AllowanceKReceiptAdmin is", amount)
		return amount
	}
}
