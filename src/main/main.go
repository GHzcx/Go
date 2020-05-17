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
	//spider.GetBrandInfoList(`110000`, `bj`)
	//fmt.Println(len(info),info)
	alldata := make(spider.AllData, 0)
	result := make (chan spider.Set, 0)
	done := make(chan struct{}, len(spider.Abbr))
	for city, abbr := range spider.Abbr {
		go spider.GetBrandInfoList(done,result, city, abbr)
	}
	for working := len(spider.Abbr); working > 0; {
		select {
			case <-done:
				working--
			case s:= <-result:
				alldata[s.CityCode] = append(alldata[s.CityCode], s.HouseInfo)
		}
	}
	//处理goroutine结束以后通道中存在的数据
DONE:
	for {
		select {
			case s:= <-result:
				alldata[s.CityCode] = append(alldata[s.CityCode], s.HouseInfo)
			default:
				break DONE
		}
	}
	for cityCode,info := range alldata {
		fmt.Println(spider.List[cityCode])
		for i := 0; i < len(info); i++ {
			fmt.Println(info[i])
		}
	}

	defer close(result)
	defer close(done)

}