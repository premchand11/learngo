package maps

import (
	"fmt"
)

//unorderd pairs key value

func Learnmap() {
	//name<->garde
	studnetgrades := make(map[string]int)
	studnetgrades["prem"] = 23
	studnetgrades["surya"] = 24
	studnetgrades["hem"] = 25
	studnetgrades["pavan"] = 22

	fmt.Println("marks of pavan:", studnetgrades["pavan"])
	fmt.Println("marks of hem:", studnetgrades["hem"])
	delete(studnetgrades, "hem")
	fmt.Println("marks of hem:", studnetgrades["hem"])
	//checking key if exists or not
	grades, exists := studnetgrades["hem"]
	fmt.Println("garde of hem", grades)
	fmt.Println("hem exists?", exists)

	Grades, Exists := studnetgrades["surya"]
	fmt.Println("garde of surya", Grades)
	fmt.Println("surya exists?", Exists)

	for index, value := range studnetgrades {
		fmt.Printf("key is %s and marks are %d\n", index, value)
	}

}
