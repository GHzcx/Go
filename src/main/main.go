package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"
)
func GetBrandId(str string) string {
	reg := regexp.MustCompile(`<a href="/chuzu/.*?/brand/(.*?)/`)
	return reg.FindAllString(str, -1)[0]
}

func main() {
	//url := "https://app.api.ke.com/Rentplat/v1/house/total?city_id=110000&condition=ht1ab200301001000&request_ts=1589371036&scene=list"
	//url := "https://m.ke.com/chuzu/bj/brand/pg1/?ajax=1"
	//url := "https://m.ke.com/chuzu/bj/brand/200302019000/info/"
	//str := spider.Download(url)
	//https://app.api.ke.com/Rentplat/v1/house/total?city_id=110000&condition=ht1ab200306000695&request_ts=1589468641&scene=list
	v := url.Values{
		"city_id": []string{"110000"},
		"condition": []string{"ab200306000695"},
		"request_ts": []string{strconv.FormatInt(time.Now().Unix(), 10)},
		"scene": []string{"list"},
	}
	fmt.Println(v.Encode())
	//fmt.Println(time.Now().Unix())
	//fmt.Println(spider.GetShopNum(str))

	//info, len := spider.GetBrandInfos(str)
	//info := spider.GetBrandInfoList(`110000`, `bj`)
	//fmt.Println(len(info),info)

}