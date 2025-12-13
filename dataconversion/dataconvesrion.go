package dataconversion

import (
	"fmt"
	"strconv"
)

func Learndataconversion() {
	num := 42
	fmt.Println(num)
	fmt.Printf("type is:%T\n", num)

	var data float64 = float64(num)
	fmt.Println(data)
	fmt.Printf("type is:%T\n", data)

	num2 := 123
	str := strconv.Itoa(num2)
	fmt.Printf("type is:%T\n", str)

	num_string := "123456"
	num3, _ := strconv.Atoi(num_string)
	fmt.Println(num3)
	fmt.Printf("type is:%T\n", num3)


	num_string2:="2.678"
	num4,_:=strconv.ParseFloat(num_string2,64)
	fmt.Println(num4)
	fmt.Printf("type is:%T\n", num4)


}
