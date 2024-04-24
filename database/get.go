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
