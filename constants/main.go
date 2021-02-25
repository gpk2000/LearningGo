package main

import (
	"fmt"
	"math"
)

const s string = "constant string"

func main() {
	fmt.Println(s)

	// numeric const has no type until it is explicitly specified or implicitly inferred
	const n = 50000000 // type is inferred
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	const A int32 = 100 // type is explicitly specified
	fmt.Println(int64(A))
	fmt.Printf("%T\n", n)

	const B = math.Pi
	fmt.Printf("%T\n", B)

	fmt.Println(math.Sin(n))
}
