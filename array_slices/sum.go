package main

func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbers ...[]int) (total []int) {
	for _, slice := range numbers {
		total = append(total, Sum(slice))
	}
	return
}
