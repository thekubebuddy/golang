# Workbook for my Go Learning

Table of content
===================
1. Variables: declaration & Initialization, constant 
2. Arrays 
3. Slices, Slices literal, appending to a slices
4. maps
5. Conditional statements: if, if else
6. Looooop: for, for as "while"
7. for range: with map, arrays, blank operator
8. Functions: returning fuctions, multi-return function, naked function, "defer" function
9. **Structs 
10. Pointers
11. Error handlings
12. Switch
13. **Go methods on structs, pointer receivers 
15. **Go interfaces(very important)
16. Go cmds


* statically typed: makes it efficient
* compiled language
* 25 keywords rather that 
* **Concurrency support**
* OOP phylosical, though not having inheritence
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



#### 2. Arrays

```go
func main(){
	x := [7]int{12, 3, 4, 33, 4, 45, 2}
	for i := 0; i < len(x); i++ {
		f.Println(x[i])
	}
	f.Println(x[1:5])
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


func main(){
	// for ever
	for {

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


#### 7. for range

```go

a := []int{12,34,334,44,44}
for _,v := range a{
	f.Println(v)
}
 
vertex := make(map[string]int)
b["rectange"] = 3
b["square"] = 4
b["triangle"] = 3
```




#### 7.switch
The break statement that is needed at the end of each case is provided automatically in Go.




#### 8.	Pointers
* Golang has no pointers arithmatic, unlike C,C++. 

```go
func main(){
	x := 9
	y := &x
	f.Println(x)
	f.Println(y)
	f.Println(*y)
	*y = 10
	f.Println(x)
	// printing y's address
	f.Println(&y)

}
```

#### 13. Go method & method-with-pointers(pointers receiver)

```go
type Vertex struct {
	X, Y float32
}
func (v Vertex) Max() float32 {
	if v.x > v.y {
		return v.x
	}
	return v.y
}
func (v *Vertex) Scale(factor float32) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func main() {
	v1 := Vertex{1, 2}
	f.Println(v1)
	v1.Scale(10)
	f.Println(v1)
	f.Println(v1.Max())
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


References:
```
https://medium.com/rungo/error-handling-in-go-f0125de052f0
https://www.hackerearth.com/practice/notes/introduction-to-golang/
https://golang.org/doc/effective_go#blank
https://gobyexample.com/string-formatting
https://golang.org/pkg/fmt/
https://tour.golang.org/moretypes/7
https://www.youtube.com/watch?v=YEKjSzIwAdA
***https://golang.org/doc/faq#unused_variables_and_imports
```