package main
import (
	"fmt"
	"strings"
)
func main(){
	str1:="hello"
	str2:="world"
	fmt.Println(str1+" "+str2)
	// strings in go are mutable
	str2="Wowdw w..!" // assinging new world value to the str2
	fmt.Println("Mutable string: "+str1+" "+str2)
	
	// String funtion for the string
	str3:=str1+" "+str2 
	fmt.Println(strings.Split(str3,"w"))

}
