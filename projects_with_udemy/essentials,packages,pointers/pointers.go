package main

import "fmt"

func main() {
	age:= 32
    
    var agePointer *int

	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	// adultYears := getAdultsAge(agePointer)
	// fmt.Println(": ", adultYears)

	fmt.Println(age)

}

func getAdultsAge(age *int) int {
    *age = *age - 18
	return *age
}