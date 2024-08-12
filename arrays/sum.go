package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllUsingMake(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for index, slice := range numbersToSum {
		sums[index] = Sum(slice)
	}

	return sums
}

func SumAllUsingAppend(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func SumAllHeads(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		head := slice[:1]
		sums = append(sums, Sum(head))
	}

	return sums
}
