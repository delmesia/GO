package main

func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}
