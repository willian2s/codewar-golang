package cat_years_dog_years

func CalculateYears(years int) (result [3]int) {
	catYears := 0
	dogYears := 0

	for i := 1; i <= years; i++ {
		switch i {
		case 1:
			catYears += 15
			dogYears += 15
		case 2:
			catYears += 9
			dogYears += 9
		default:
			catYears += 4
			dogYears += 5
		}
	}

	return [3]int{years, catYears, dogYears}
}
