package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := []string{}

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Alex")
    
    fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	fmt.Println(courseRatings)

	for i, values := range userNames {
		fmt.Println(i)
		fmt.Println(values)

	}
} 