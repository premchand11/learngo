package loops

import "fmt"

func Learnloops() {
	Counter := 0
	for i := 0; i < 5; i++ {
		Counter++
		fmt.Println(Counter)
	}
	HELLO := 0
	for {
		fmt.Println("loop infi")
		HELLO++
		if HELLO > 3 {
			break
		}

	}

}

// no do-while
//range key  word

func Learnloops2() {
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("index of number:%d,and value of number %d\n", index, value)
	}
}
