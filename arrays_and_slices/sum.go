package main

func Sum(number []int) int {
	res := 0
	for _, num := range number {
		res += num
	}
	return res
}

func SumTrails(number []int) int {
	res := 0
	for i, num := range number {
		if i != 0 {
			res = res + num
		}
	}
	return res

}

func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, number := range numbers {
		sums[i] = Sum(number)
	}

	return sums

}

func SumAllTrails(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sumsOfTrails := make([]int, lengthOfNumbers)

	for i, number := range numbers {
		sumsOfTrails[i] = SumTrails(number)
	}
	return sumsOfTrails
}
