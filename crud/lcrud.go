package crud

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todos struct {
	UserID    int64  `json:"userId"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("got error during get request", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("the probelm:", res.StatusCode)
	}
	// data, err1 := io.ReadAll(res.Body)
	// if err1 != nil {
	// 	fmt.Println("got error during reading the data", err1)
	// 	return
	// }
	var todo Todos
	err2 := json.NewDecoder(res.Body).Decode(&todo)
	if err2 != nil {
		fmt.Println("some kind of error in decoding", err2)
	}
	fmt.Println(todo)

}

func PerformPostRequest() {
	todo := Todos{
		UserID:    23,
		Title:     "tilte",
		Completed: true,
	}
	jsondata, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error while converting", err)
		return
	}
	//convert data to string

	jsonstr := string(jsondata)

	//convert string to reader
	jsonreader := strings.NewReader(jsonstr)
	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "application/json", jsonreader)
	if err != nil {
		fmt.Println("error wheile post request", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("got an error during io read",err)
	}
	fmt.Println(string(data))

}

func Learncrudd() {
	fmt.Println("learning crud ")
	//PerformGetRequest()
	PerformPostRequest()
}
