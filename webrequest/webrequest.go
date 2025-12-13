package webrequest

import (
	"fmt"
	"io"
	"net/http"
)

func Learnwebrequest() {
	fmt.Println("learing webrequets")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error got while getting req:", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("type of response %T", res)

	data, errs := io.ReadAll(res.Body)
	if errs != nil {
		fmt.Println("error got while reading req:", errs)
		return
	}
	fmt.Println(string(data))

}
