package main

import "fmt"

func main() {

	for i := 1; i < 100; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println(i, "is FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, "is Fizz")
		case i%5 == 0:
			fmt.Println(i, "is Buzz")
		}

	}
	fmt.Println()
}
