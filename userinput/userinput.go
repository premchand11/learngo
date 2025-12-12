package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func Userinput1() {
	fmt.Println("hey what is your name?")
	var name string
	fmt.Scan(&name)
	fmt.Println("hello,mr my name is:", name)
}

// here its a coorect way but still fails to read  after space
// so we need to use buffio in order for better input
// stores data in buffer(temp stroage)
// standarsd input

func Userinput2() {
	fmt.Println("Hello what is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("my name is ", name)

}
