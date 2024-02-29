package main

import (
	"fmt"

	"github.io/thekubebuddy/golang/calc/internal/utils"
)

func main() {
	fmt.Println("Calculator program called..!!")
	x, y := 5, 10
	fmt.Printf("Sum of %d & %d: %d\n", x, y, utils.Add(x, y))
	fmt.Printf("Subraction of %d & %d: %d\n", x, y, utils.Sub(x, y))
	fmt.Printf("Division of %d & %d: %.2f\n", x, y, utils.Div(x, y))
}
