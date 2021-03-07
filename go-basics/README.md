# Workbook for my Go Learning

Table of content
===================
1. Variables: declaration & Initialization 
2. Arrays & slices
3. maps
4. Conditional statements: if, if else
5. Looooop: for, for as "while"
6. Functions: returning fuctions, multi-return function
7. **Structs 
8. Pointers



* Go is quite casesensitive
for e.g
```go
func main(){
}
```
is not equal to 
```go
func main()
{

}
```

##### 1. Variables declaration & initialization

```go
package main
import (
    "fmt"
)

func main(){
    // var <name> <type> = <value>
    var x int = 100
    
    // shorthand
    y := 100
    
    //  all variables & arrays are initializes to 0
    var z int
    fmt.Printlln(z)

}
```

#### 2. Golang Slices: Arrays of dynamic length

```go
func main(){
    x := []int{1,22,34,554,33,45,1}
    fmt.Println(x)
    x = append(x,556,6,6)
    fmt.Println(x)
}

```


#### 3. Map
```go
func main(){
	// map[type-of-key]type-of-values-in-map -> make(map[string]int)
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["rectangle"] = 4
	vertices["hexagon"] = 5
	vertices["square"] = 5
	fmt.Println(vertices)
	fmt.Println(vertices["rectangle"])
	//  deleting the keys from the map
	delete(vertices, "square")

	fmt.Println(vertices)
}
```

#### 4. Conditional statement: if, if else
* Really case sensitive
```go
func main(){
	x := 10
	y := 100
	if x > y {
		fmt.Println("x is greater than y")
	}else{
		fmt.Println("x is smaller than y")
	}

	percentage := 75
	if percentage > 60 && percentage < 75{
		fmt.Println("You are B graded.!")
	}else if percentage >= 75{
		fmt.Println("You are distincted.!")
	}else if percentage<35 {
		fmt.Println("Sorry.! You are legend.!")
	}

}
```

#### 5. for loop, for as while

* for loop
```go

func main(){
    for i:= 0; i<10; i++{
        fmt.Println(i)
    }
}
```

* "for" as "while"
```go
func main(){
    x := []int{12,34,554,33,45,1}
    i := len(x)-1
    for i>=0{
	fmt.Println(x[i])
	i--
	}
}
```






Refrences:

https://tour.golang.org/moretypes/7
https://medium.com/rungo/error-handling-in-go-f0125de052f0
