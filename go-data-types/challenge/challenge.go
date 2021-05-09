package main

import (
	"fmt"
	"strings"
)

func main() {
	var val string
	val = "MCLX"
	data := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	result := 0

	slice := strings.Split(val, "")

	fmt.Println(slice)

	for _, c := range slice {
		num, exist := data[c]
		if exist {
			result = result + num
			fmt.Printf("とおったよ")
		} else {
			fmt.Printf("Error: %s is bad string", c)
			return
		}

	}
	fmt.Println("Result is", result)
}
