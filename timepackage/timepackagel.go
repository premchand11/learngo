package timepackage

import (
	"fmt"
	"time"
)

func Learntimepack() {
	//2006-01-02 15:04:05 => keep in mind this cuz this is the tome when go relases yearmonth date hours minutes seconds
	//24 hour fomat

	currenttime := time.Now()
	fmt.Printf("time type %T\n", currenttime)
	formatted := currenttime.Format("02-01-2006,Monday")
	fmt.Println(formatted)
}
