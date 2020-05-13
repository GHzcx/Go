package main

import (
	"fmt"
	"spider"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(spider.GetProxyUrl())
	fmt.Println(spider.List)
}