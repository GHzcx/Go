package main

import (
	"fmt"
	"spider"
)

func main() {
	alldata := make(spider.AllData, 0)
	Queue := make (chan int, 2)
	result := make (chan spider.Set, 10)
	done := make(chan struct{}, len(spider.CityAbbr))

	for city, abbr := range spider.CityAbbr {
		Queue <- 1
		go spider.GetBrandInfoList(Queue, done, result, city, abbr)
		fmt.Println("Queue", len(Queue))
	}
	for working := len(spider.CityAbbr); working > 0; {
		select {
			case s:= <-result:
				alldata[s.CityCode] = append(alldata[s.CityCode], s.HouseInfo)
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
	for cityCode,info := range alldata {
		fmt.Println(spider.CityList[cityCode], len(info))
	}

	defer close(result)
	defer close(done)
	defer close(Queue)

}