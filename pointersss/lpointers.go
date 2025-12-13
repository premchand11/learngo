package pointersss

import "fmt"

// stores the memory adress
// vlaue stored in varible rather than the value
// effiecnt memeory mangement
// we use * this symbol

func modifyvaluebyref(a *int) {
	*a = *a * 2

}

func Learnpointers() {
	var num int
	num = 2

	var ptr *int
	ptr = &num

	fmt.Println("num has vale:", num)
	fmt.Println("pointer has:", ptr)

	num2 := 3
	ptr2 := &num2

	fmt.Println("num has vale:", num2)
	fmt.Println("pointer has:", ptr2)
	fmt.Println("data contain through pointe:", *ptr2)

	value := 10
	modifyvaluebyref(&value)
	fmt.Println("value contains", value)
}
