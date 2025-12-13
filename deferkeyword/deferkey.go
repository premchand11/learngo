package deferkeyword

//defer basically returs at the last
//lifo last in first out
//each defer stcaked on to the list


import "fmt"

func add(a,b int)int{
	return a+b
}

func Learndefer() {
	fmt.Println("start of the prgm")
	data := add(5, 6)
	defer fmt.Println(data)
	defer fmt.Println("middle of the prgm")
	defer fmt.Println("end of the prgm")

}
