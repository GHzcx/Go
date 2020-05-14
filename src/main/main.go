package main

import (
	"fmt"
	"regexp"
	"spider"
)
func GetBrandId(str string) string {
	reg := regexp.MustCompile(`<a href="/chuzu/.*?/brand/(.*?)/`)
	return reg.FindAllString(str, -1)[0]
}

func main() {
	//url := "https://app.api.ke.com/Rentplat/v1/house/total?city_id=110000&condition=ht1ab200301001000&request_ts=1589371036&scene=list"
	url := "https://m.ke.com/chuzu/bj/brand/pg1/?ajax=1"
	str := spider.Download(url)
	//fmt.Println(spider.GetBrandIdAndName(str))
	info, len := spider.GetBrandInfos(str)
	fmt.Println(len, info)
}