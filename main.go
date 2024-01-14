package main

import "fmt"

func main() {
	array := [5]int{5, 3, 4, 1, 2}
	temp := 0

	for i := 0; i < len(array)-1; i++ {
		for j := 1; j < len(array)-1; j++ {
			if array[j-1] > array[j] {
				temp = array[j]
				array[j-1] = array[j]
				array[j] = temp
			} else {
				break
			}
		}
	}
	printArray(array)
}

func printArray(array [5]int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array)
	}
}
