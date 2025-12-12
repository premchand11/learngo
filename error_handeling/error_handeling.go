package errorhandeling

import "fmt"

func divison(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cant be zero")
	}
	return a / b, nil

}

func Errorhandel() {
	ans, _ := divison(101, 2)
	fmt.Println("answer is", ans)
	secans, err := divison(235, 0)
	fmt.Println("answer is", secans, err)

}

// _ is a underscore that ignorieng the error or the varible coming
