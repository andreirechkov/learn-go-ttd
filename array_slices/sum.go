package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numberToSum := range numbersToSum {
		if len(numberToSum) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numberToSum[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}