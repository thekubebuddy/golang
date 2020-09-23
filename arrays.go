package main
import "fmt"

func todo() {
	// integer array
	int_arr:=[]int{1,2,24,33}
	// string array
	str_arr:=[]string{"hello","world","..!"}
	fmt.Println(int_arr)
	fmt.Println(str_arr)
	str_arr2:=[]string{"This","is","an","diffecrent string array.!"}
	
	// Appending additional elements to the string array
	str_arr2=append(str_arr2,"hellow")
	fmt.Println(str_arr2)

}

func main(){
	todo()
	// All of the variables initialized with 0
	var x int
	fmt.Println(x)
}