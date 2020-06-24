package main

import (
	"fmt"
	"net/http"
)

func getError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer)
		fmt.Println(request)
		fmt.Fprint(writer, "hello")
	})
	err := http.ListenAndServe(":1550", nil)
	getError(err)
}