package database

func SetPersonalAllowanceDB(setNum float64) error {
	// Prepare the SQL statement
	stmt, err := DB.Prepare(`UPDATE tax."allowance" SET Amount = $1 WHERE Type = 'Personal'`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided value
	_, err = stmt.Exec(setNum)
	if err != nil {
		return err
	}

	return nil
}

func SetKReceiptAllowanceDB(setNum float64) error {
	// Prepare the SQL statement
	stmt, err := DB.Prepare(`UPDATE tax."allowance" SET Amount = $1 WHERE Type = 'k-receipt'`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided value
	_, err = stmt.Exec(setNum)
	if err != nil {
		return err
	}

	return nil
}
