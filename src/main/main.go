package main

import (
	"fmt"
	"regexp"
	"spider"
	"time"
)
func GetBrandId(str string) string {
	reg := regexp.MustCompile(`<a href="/chuzu/.*?/brand/(.*?)/`)
	return reg.FindAllString(str, -1)[0]
}

func main() {
	//spider.GetBrandInfoList(`110000`, `bj`)
	//fmt.Println(len(info),info)

	alldata := make(spider.AllData, 0)
	Queue := make (chan int, 2)
	result := make (chan spider.Set, 100)
	done := make(chan struct{}, len(spider.Abbr))
	go func(result chan spider.Set, data *spider.AllData) {
		retry := 10
		for retry > 0 {
			select {
			case s:= <-result:
				(*data)[s.CityCode] = append((*data)[s.CityCode], s.HouseInfo)
			default:
				time.Sleep(time.Duration(60)*time.Second)
				retry--
			}
		}
	}(result, &alldata)
	for city, abbr := range spider.CityAbbr {
		Queue <- 1
		go spider.GetBrandInfoList(Queue, done,result, city, abbr)
		fmt.Println("Queue", len(Queue))
	}
	for working := len(spider.CityAbbr); working > 0; {
		select {
			case <-done:
				working--
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
	time.Sleep(time.Duration(60)*time.Second)
	for cityCode,info := range alldata {
		fmt.Println(spider.CityList[cityCode])
		for i := 0; i < len(info); i++ {
			fmt.Println(info[i])
		}
	}

	//defer close(result)
	//defer close(done)

}