package main

/**
- Every go project will have main.go
- main.go only pulls all the packages and other packages cannot import anything from main.go
*/

import (
	"fmt"
)

func main()  {
	fmt.Println("Hello world")
}