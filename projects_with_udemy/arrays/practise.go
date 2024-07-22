package main

import "fmt"

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

type Product struct{
id string
title string
number float64
}

func main() {
// 1
	hobbies := [3]string{"judo", "tabletennis", "basketball"}
	fmt.Println(hobbies) 

// 2 
    tasktwo := hobbies[0]
	fmt.Println(tasktwo)

	tasktwotwo := hobbies[1:3]
	fmt.Println(tasktwotwo)

// 3
    mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)


// 4
    fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:]
	fmt.Println(mainHobbies)

// 5
courseGoals := []string{"Laern go!", "Learn all the basics"}
fmt.Println(courseGoals)

// 6 
courseGoals[1] = "Learn all the details"
courseGoals = append(courseGoals, "Learn all the basics!")
fmt.Println(courseGoals)

// 7 
products := []Product{
    {
		"first-product",
		"A first Product",
		12.99,
	},

	{
		"second-product",
		"A second Product",
		129.99,
	},
}

fmt.Println(products)

newProduct := Product{
	"third-product",
	"A third product",
	15.99,
}

products = append(products, newProduct)

fmt.Println(products)
}


