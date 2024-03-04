package fakebinary

// FakeBinary converts a string of decimal digits to binary digits.
//
// It takes a string `input` as input and returns a string of binary digits.
func FakeBinary(input string) string {
	result := ""
	for _, char := range input {
		if char < '5' {
			result += "0"
		} else {
			result += "1"
		}
	}

	return result
}
