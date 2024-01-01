package integers

import "fmt"

func Add(a, b int) (sum int) {
	sum = a + b
	return
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
}
