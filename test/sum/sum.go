package sum

//Sum will add up all the  numbers of the array
func Sum(numbers []int) int {
	sumA := 0
	for _, value := range numbers {
		sumA += value
	}
	return sumA
}
