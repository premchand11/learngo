package functions

import "fmt"

func LearnFunctions() {
	fmt.Println("learning functions")
}

func Add(a, b int) int {
	return a + b
}

func Subtraction(a, b int) (ans int) {
	ans = a - b
	return ans
}
