package slice

import "fmt"

//slice is a flexible ds which is built on arrays

func Learnslice() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 7, 8, 9)
	fmt.Println(numbers)

	fmt.Println("----------------------------------------")

	check := []int{1, 2, 3}
	fmt.Println(check)
	fmt.Println(len(check))
	fmt.Println(cap(check))

	fmt.Println("---------------make---------------")
	yellowers := make([]int, 3, 5)
	fmt.Println(yellowers)
	fmt.Println(len(yellowers))
	fmt.Println(cap(yellowers))
	//capactity doubles for every exceeding cap
	//we cant declare empty slice without make functions

}

//dynamic array
//reference to undelying slice
//make function take len and capacity
