package iteration

func repeat(input string, time int) string {
	result := input
	for i := 0; i < time-1; i++ {
		result += input
	}
	return result
}
