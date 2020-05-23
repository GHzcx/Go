package main

import (
	"fmt"
	"os"
	"spider"
	"strconv"
	"time"
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
	time.Sleep(60 * time.Second)
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
	totalCount := 0
	totalNum := 0
	fp, err := os.OpenFile("./result.txt", os.O_CREATE | os.O_APPEND, 6)
	if err != nil {
		fmt.Println("文件打开失败。")
		return
	}
	for cityCode,info := range alldata {
		fmt.Println(spider.CityList[cityCode], len(info))
		for _,value := range info {
			if value.HouseVrNum > 0 {
				totalNum++
				fp.WriteString( value.GongyuId + " " + value.GongyuName + " " + strconv.Itoa(value.HouseVrNum) + "\n")
			}
			totalCount += value.HouseVrNum
		}
	}
	fmt.Println(totalCount, totalNum)
	defer close(result)
	defer close(done)
	defer close(Queue)

}