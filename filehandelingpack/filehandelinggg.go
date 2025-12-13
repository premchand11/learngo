package filehandelingpack

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Learnfilehandel() {
	writefile()
	// readandgiveoutput()
	shortcut()

}

func writefile() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("error came while creating file:", err)
		return
	}
	defer file.Close()
	contet := "hello world"
	byte, error := io.WriteString(file, contet+"\n")
	if error != nil {
		fmt.Println("error came while creating file:", error)
		return
	}
	fmt.Println(byte)
	fmt.Println("sucessfully created file")
}

func readandgiveoutput() {

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error came while opening file:", err)
		return
	}

	defer file.Close()

	//create buffer
	buffer := make([]byte, 1024)
	//read the file content into buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error came while reading  file:", err)
			return

		}
		//process the content
		fmt.Println(string(buffer[:n]))
	}
}

//shortcut methof io-util

func shortcut() {
	conetet, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("got error during reading", err)
		return
	}
	fmt.Println(string(conetet))

}

//iofile is the total file will be on memory
//depricated this
//we can just use os.readfile
//but by using shortcu there will be memory issues