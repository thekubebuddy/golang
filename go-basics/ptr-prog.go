package main
import "fmt"


func swap(x,y *int){
	
	temp:=*x
	*x=*y
	*y=temp
}


func main(){
	m1:=3
	// ptr will store the address for m1 
	ptr:=&m1
	// Simply prints the addr of the m1
	fmt.Println("Address value for the \"m1\" variable: ",ptr)

	// Deferencing the the pointer, actuall value stored at the memory location
	fmt.Println("Value stored at the address pointed by the \"ptr\": ",*ptr)
	
	// Referencing the address value for the "ptr"
	fmt.Println("Address value for the \"ptr\": ",&ptr)
	
	fmt.Println("-----------------------------------")
	fmt.Println("Pass by referene example in go using POINTERS")
	fmt.Println("-----------------------------------")
	
	x,y:=5,8
	fmt.Println("Before swap call: ",x,y)
	swap(&x,&y)
	fmt.Println("After swap call:  ",x,y)
	fmt.Println("-----------------------------------")
	
}


