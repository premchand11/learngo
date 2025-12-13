package switchcase

import "fmt"

func Learnswitch() {
	day := 3

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	default:
		fmt.Println("doesnotexist")
	}
}
