# Workbook for my Go Learning

Table of content
===================
1. Variables: declaration & Initialization 
2. Arrays & slices
3. maps
4. Conditional statements: if, if else
5. Looooop: for, for as "while"
6. Functions: returning fuctions, multi-return function 
8. **Structs 
8. Pointers
10. Error handlings
* Go cmds

#### 1. Variables declaration & initialization

```go
package main
import (
   f "fmt" // package aliasing
)

func main(){
    // var <name> <type> = <value>
    var x int = 100
    
    // implicit type assingment or short variable declarations
    y := 100
	a, b := 10,12
	
    //  all variables & arrays are initializes to 0
    var z int
    f.Printlln(z)

}
```

* Go is quite case-sensitive
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
#### 6. Funtions: Multiple returning function, "naked" return
```go
func main(){
	avg := get_avg(3,4,5)
	fmt.Printf("avg of 3 numbers: %f",avg)
	x,y := swap("hello","world")
	fmt.Printf("\n%s %s\n",x,y)
}

func get_avg(a,b,c float32) float32{
	return ((a+b+c)/3)
}

//  Pass by refrence
func swap(a,b string)(b,a string){
	return b,a
}
```


#### Go commands

```bash
# printing the GOPATH
go env GOPATH

# canonically rewriting the go files
go fmt

# build
go build main.go


```

Refrences:

https://medium.com/rungo/error-handling-in-go-f0125de052f0
https://gobyexample.com/string-formatting
https://tour.golang.org/moretypes/7
https://tour.golang.org/basics/5
