package main

import "fmt"

func main() {
	// using var Go can infer its type from initialized variables
	// the variables get zero initialized when you define a type like in line 17-18
	// := is a short hand notation of declaring and defining a variable and type will
	// be inferred from the initialzation literal
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
