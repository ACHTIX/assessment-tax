package database

import (
	"github.com/ACHTIX/assessment-tax/model"
)

func GetAllowance() (*model.Allowance, error) {
	var amount model.Allowance
	err := DB.QueryRow(`SELECT amount FROM tax."allowance"`).Scan(&amount.Amount)

	if err != nil {
		return nil, err
	}

	return &amount, nil
}

func GetPersonalAllowance() (float64, error) {
	var personalDeduct model.GetTaxDeductStruct
	err := DB.QueryRow(`SELECT amount_deduct,id,is_active,create_at FROM public."master_deduct" WHERE is_active = TRUE AND type_deduct = 'personal'`).Scan(&personalDeduct.PersonalDeduct, &personalDeduct.Id, &personalDeduct.Is_active, &personalDeduct.Create_at)
	return personalDeduct.PersonalDeduct, err
}
