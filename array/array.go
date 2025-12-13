package array

import "fmt"

func Learnarray() {
	var name [5]string

	name[0] = "prem"
	name[1] = "chand"

	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(name[0])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	var trail [7]int
	fmt.Println(trail)
	//always default values is zero
	//array is not dynamic

}
