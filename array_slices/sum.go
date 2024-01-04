package main

func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAllTails(numbers ...[]int) (total []int) {
	for _, slice := range numbers {
		if len(slice) == 0 {
			total = append(total, 0)
		} else {
			tail := slice[1:]
			total = append(total, Sum(tail))
		}
	}
	return
}
