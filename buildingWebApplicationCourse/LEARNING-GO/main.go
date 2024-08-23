package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool `json:"has_dog"`
}

func main() {
	myJson := `
	[
	   {
	       "first_name": "Clark",
		   "last_name": "Kent",
		   "hair_color": "green",
		   "has_dog": "true"
		},

		 {
	       "first_name": "Estemis",
		   "last_name": "Kara",
		   "hair_color": "black",
		   "has_dog": "false"
		}
	]
	`

	var unmarshelled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshelled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Println("unmarshalled: %v", unmarshelled)


	var mySlice []Person

	var m1 Person
	m1.FirstName = "Hello"
	m1.LastName = "World"
	m1.HairColor = "red"
	m1.HasDog = false

   mySlice = append(mySlice, m1)


   newJson, err := json.MarshalIndent(mySlice, "", " ")

   if err != nil {
	log.Println("error", err)
   }

   fmt.Println((string(newJson)))
}
