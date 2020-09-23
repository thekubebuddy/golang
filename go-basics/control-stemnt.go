package main

import "fmt"

func main() {
	x, y := 1, 11

	// if .. else ..
	if x<y {
		fmt.Println("x is smaller than y")
	}else{
		fmt.Println("y is smaller than x")
	}

	// if else if 
	if x>y {
		fmt.Println("x is greater than y")
	}else if x<y {
		fmt.Println("x is smaller than y")
	}else{
		fmt.Println("x and y both are equal.!")
	}

	// for loop 
	for i :=1; i<10; i++ {
		fmt.Println(i)
	}
	
	// Infinite for loop
	// for {
	// 	fmt.Println("Infinity..!!")
	// }

	// Traversing an array with the for loop 
	arr := []int{1,2,34,76}

	for i,value := range arr{
		fmt.Println("arr[",i,"]->",value)
	}
}
