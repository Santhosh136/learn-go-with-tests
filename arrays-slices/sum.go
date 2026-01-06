package arraysslices

func Sum(numbers []int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(slices ...[]int) []int {

	// length := len(slices)
	// result := make([]int, length)

	// for i, sl := range slices {
	// 	result[i] = Sum(sl)
	// }

	var result []int

	for _, numbers := range slices {
		sum := Sum(numbers)
		result = append(result, sum)
	}

	return result
}

func SumAllTails(slices ...[]int) []int {
	result := []int{}
	for _, numbers := range slices {
		if len(numbers) == 0 {
			result = append(result, 0)
			continue
		}
		sum := Sum(numbers[1:])
		result = append(result, sum)
	}
	return result
}
