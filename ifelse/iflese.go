package ifelse

import "fmt"

func Learnifelse() {
	x := 10
	if x == 10 {
		fmt.Println("this is 10 yayyayayaya")
	}
	z := 9
	if z > 9 {
		fmt.Println("z is greate2 than 9")
	} else if z < 9 {
		fmt.Println("z is smaller than 9")
	} else {
		fmt.Println("z is 9")
	}

	if x > 9 && z == 9 {
		fmt.Println("double")
	}
	//we can use or || condition as well
}
