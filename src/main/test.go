package main

import (
	"fmt"
	"net/url"
	"spider"
	"strconv"
	"time"
)
func test(sig chan int) {
	sig <- 1
}

func main () {

	v := url.Values{
		"city_id": []string{"310000"},
		"condition": []string{"ht1vr1"},
		"request_ts": []string{strconv.FormatInt(time.Now().Unix(), 10)},
		"scene": []string{"list"},
	}
	Url := fmt.Sprintf("https://app.api.ke.com/Rentplat/v2/house/list?%s", v.Encode())
	fmt.Println(spider.Download(Url))
}