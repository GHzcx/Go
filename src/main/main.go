package main

import (
	"fmt"
	"spider"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(spider.GetProxyUrl())
	fmt.Println(spider.List)
	url := "https://app.api.ke.com/Rentplat/v1/house/total?city_id=110000&condition=ht1ab200301001000&request_ts=1589371036&scene=list"
	//url := "https://m.ke.com/chuzu/bj/brand/pg1/?ajax=1"
	spider.Download(url)
	h := fmt.Sprintf("%s, %s", "hah", "HH")
	fmt.Print(h)
}