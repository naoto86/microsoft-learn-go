package main

import "fmt"

func main() {
	var sqroot float64 = 25

	currguess := 1.0
	prevguess := 0.0

	for count := 1; count <= 10; count++ {
		prevguess = currguess
		currguess = prevguess - (prevguess*prevguess-sqroot)/(2*prevguess)
		if currguess == prevguess {
			break
		}
		fmt.Println("A guess for square root is ", currguess)
	}
	fmt.Println("Square root is:", currguess)
}
