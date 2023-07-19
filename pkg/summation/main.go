package summation

func Summation(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	return total
}
