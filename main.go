package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("HHH")
	url := ""
	client = &http.Client{}
	req,_=http.NewRequest("GET", )

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("get info error")
	}
	fmt.Print(resp.Body)

}