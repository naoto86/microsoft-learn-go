package main

import "fmt"

func main() {
	// studentsAge := map[string]int{
	// 	"john": 32,
	// 	"bob":  31,
	// }
	// fmt.Println(studentsAge)

	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)

	fmt.Println("Bob's age is", studentsAge["bob"])

	age, exist := studentsAge["christy"]

	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

	// delete(studentsAge, "john")
	// fmt.Println(studentsAge)

	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
