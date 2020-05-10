package calc

// Add return sum of two integers
func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum = sum + num
	}

	return sum
}

func substract(i int, j int) int {
	return j - i
}
