package main
import "fmt"


func Sum(x,y int)(int){
	return x+y
}

func MultiReturn(x,y int)(string,int){
	return "Sum is",Sum(x,y)
}


func main(){
	fmt.Println("Hello from main funtion.!")
	
	// funtion with the arguments 
	x,y := 3,4
	fmt.Println(Sum(x,y))

	// Funtion with different return muliple type
	x,y=4,5
	fmt.Println(MultiReturn(x,y))	

}