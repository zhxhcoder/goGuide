package main

import (
	"fmt"
	"package-demo/cal"
	"package-demo/cal/multi"
)

func main() {

	result := multi.Multi(1, 2)
	fmt.Printf("%d\n", result)

	result2 := cal.Add(1, 2)
	fmt.Printf("%d\n", result2)

	result3 := cal.Subtract(1, 2)
	fmt.Printf("%d\n", result3)

}
