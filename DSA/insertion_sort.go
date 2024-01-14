package main

import "fmt"

func main() {
	array := [5]int{5, 3, 4, 1, 2}

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if array[j-1] > array[j] {
				temp := array[j]
				array[j] = array[j-1]
				array[j-1] = temp
			} else {
				break
			}
		}
	}
	PrintArray(array)
}

func PrintArray(array [5]int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
