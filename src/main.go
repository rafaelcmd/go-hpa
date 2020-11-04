package main

import (
	"fmt"
	"math"
)

func raiz(num float64) string {
	x := num
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}

	return "Code.education Rocks!"
}

func main() {
	fmt.Printf(raiz(0.001))
}
