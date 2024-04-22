package util

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
