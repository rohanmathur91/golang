package main

/**
- Every go project will have main.go
- main.go only pulls all the packages and other packages cannot import anything from main.go
*/

import (
	"fmt"
)

func main() {
	var title string = "Hello world"
	title1 := "Hello world" // declaring and initializing, type inference
	fmt.Println(title)
	fmt.Println(title1)
	var (
		isEmpty bool = true
		isNew   bool = true
		value   int  = 100
	)

	fmt.Println(isEmpty, isNew, value)

	sum, product := getSumAndProduct(4, 4)
	fmt.Println(sum, product)

}

// not exported
func getSumAndProduct(a int, b int) (int, int) {
	return a + b, a * b
}
