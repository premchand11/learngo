package structs

import "fmt"

//composite data type  that groups togethere varibles

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

type contact struct {
	Email       string
	phonenumber int
}

type Address struct {
	House int
	Area  string
	State string
}

type Emplyoee struct {
	Person_details Person
	Person_contact contact
	Person_Address Address
}

func LearnStructs() {
	var prem Person
	fmt.Println("prem is a person", prem) //" "" "o
	prem.Firstname = "premchand"
	prem.Lastname = "panku"
	prem.Age = 20
	fmt.Println("prem's details", prem)

	surya := Person{
		Firstname: "surya charan",
		Lastname:  "reddy",
		Age:       23,
	}
	fmt.Println("surya's details:", surya)

	var hem = new(Person)
	hem.Firstname = "hemchander"
	hem.Lastname = "vadyavath"
	hem.Age = 23

	fmt.Println("hem's details", hem)

	employee1 := Emplyoee{
		Person_details: Person{
			Firstname: "Purvaja",
			Lastname:  "pammi",
			Age:       22,
		},
		Person_contact: contact{
			Email:       "pammipurvaja@gmail.com",
			phonenumber: 9988877755,
		},
		Person_Address: Address{
			House: 23,
			Area:  "miyapur",
			State: "Telangana",
		},
	}

	fmt.Println(employee1)
}
