package database

//func SetPersonalAllowanceDB(setNum float64) (*model.Allowance, error) {
//	// Prepare the SQL statement
//	err := DB.QueryRow(`UPDATE tax."allowance" SET Amount = ? WHERE Type = 'Personal'`)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return err, nil
//}

//&model.Allowance{AllowanceType: "Personal", Amount: setNum}

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
