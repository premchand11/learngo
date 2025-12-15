package jsonn

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	IsAdult bool   `josn:"IsAdult"`
}

func Learnjsongo() {
	fmt.Println("learning json")
	person := Person{
		Name:    "jhon",
		Age:     24,
		IsAdult: true,
	}

	fmt.Println("person datan is:", person)
	//convert in json encoding marshal

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("got error during encoding", err)
		return
	}
	fmt.Println("person data aftre encoding", string(jsonData))

	//decoding
	var personData Person
	errr := json.Unmarshal(jsonData, &personData)
	if errr != nil {
		fmt.Println("got error during encoding", errr)
		return
	}
	fmt.Println(personData)

}
