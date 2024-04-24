package util

// แยกประเภทค่าลดหย่อน
func checkAllowanceType(allowanceType string, allowanceAmount float64) float64 {
	if allowanceType == "donation" {
		amount := donationAllowance(allowanceAmount)
		return amount
	} else if allowanceType == "k-receipt" {
		amount := kReceiptAllowance(allowanceAmount)
		return amount
	}
	return 0
}

// ค่าลดหย่อนภาษีจากการบริจาค ไม่เกิน 100,000
func donationAllowance(amount float64) float64 {
	if amount >= 100001 {
		return 100000
	} else if amount <= 100000 && amount != 0 {
		return amount
	}
	return 0
}

// ค่าลด k-receipt ต้องมีค่ามากกว่า 0 บาท
// ค่าลดหย่อนภาษีจาก k-receipt สูงสุดไม่เกิน 50,000
func kReceiptAllowance(amount float64) float64 {
	if amount >= 50001 {
		return 50000
	} else if amount <= 50000 && amount != 0 {
		return amount
	}
	return 0
}

// ค่าลดหย่อนภาษีจาก k-receipt แอดมินกำหมดได้ สูงสุดไม่เกิน 100,000
// ค่าลดหย่อนส่วนตัวต้องมีค่ามากกว่า 10,000 บาท
// แอดมินสามารถกำหนดค่าลดหย่อนส่วนตัวได้โดยไม่เกิน 100,000
func AllowancePersonalAdmin(amount float64) float64 {
	if amount >= 100001 {
		return 100000
	} else {
		return amount
	}
}
