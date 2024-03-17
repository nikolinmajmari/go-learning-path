package arrays

func Sum(numbers []int) int {
	sum := 0
	/*for i := 0; i < 5; i++ {
		sum += numbers[i]
	}*/
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	/*length := len(nums)
	sums := make([]int, length)*/
	var sums []int
	for _, numbers := range nums {
		/*sums[i] = Sum(numbers)*/
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int
	for _, numbers := range nums {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
