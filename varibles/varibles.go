package varibles

import "fmt"

func VariblesTry() {
	var version string = "1.00"
	var money int = 100
	var decimall float32 = 23.45
	var isTrue bool = true
	const cantchnage int = 56666
	const yesyou string = "erri"
	prem := "name"
	fmt.Println(version)
	fmt.Println(money)
	fmt.Println(decimall)
	fmt.Println(isTrue)
	fmt.Print(cantchnage)
	fmt.Println(prem)
	fmt.Printf(yesyou)
	fmt.Println("---------------------")

	//println vs printf
	// print ln adds new line after the completion and after parametes also add space
	//printf fomramt specifier %d-integer,%f-float, %s-string,%T-type vale      and doest add any lin eor space after line or parameter
}
