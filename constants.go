package main

import (
	"fmt"
	"math"
)

const name = "John Doe"

func main() {

	fmt.Println("Const Name is => " + name)

	const n = 500000
	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
