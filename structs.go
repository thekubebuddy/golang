package main
import "fmt"


// Declaring a Car Structs

type Car struct{
	Name string
	Age int
	ModelNo int
}


func main(){
	// Declaring and initialzing the Car structs
	// var car1 Car
	car1:=Car{
		Name: "Nano",
		Age: 2,
		ModelNo: 2020,
	}
	fmt.Println(car1)
}
