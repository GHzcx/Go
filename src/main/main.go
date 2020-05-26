package main

import (
	"fmt"
	"os"
	"spider"
	"strconv"
)

func main() {
	alldata := make(spider.AllData, 0)
	vrData := make(spider.AllData, 0)
	Queue := make (chan int, 2)
	result := make (chan spider.Set, 10)
	done := make(chan struct{}, len(spider.CityAbbr))

	for city, abbr := range spider.CityAbbr {
		Queue <- 1
		go spider.GetBrandInfoList(Queue, done, result, city, abbr)
		fmt.Println("Queue", len(Queue))
	}
	vrData = spider.GetVrInfoList("310000")
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
	fp, err := os.OpenFile("./result.txt", os.O_CREATE | os.O_APPEND, 6)
	if err != nil {
		fmt.Println("文件打开失败。")
		return
	}
	for cityCode,info := range alldata {
		vrInfo := vrData[cityCode]
		for _, value := range vrInfo {
			flag := 0
			for _, item := range info {
				if value.GongyuId == item.GongyuId {
					flag = 1
					if value.HouseVrNum != item.HouseVrNum {
						fp.WriteString( value.GongyuId + " " + value.GongyuName + " " + strconv.Itoa(value.HouseVrNum) + "\n")
						break
					}
				}
			}
			if flag == 0 {
				fp.WriteString( value.GongyuId + " " + value.GongyuName + " " + strconv.Itoa(value.HouseVrNum) + "\n")
			}
		}

	}
	defer close(result)
	defer close(done)
	defer close(Queue)

}