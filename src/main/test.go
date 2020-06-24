package main

import (
	"fmt"
	"os"
	"spider"
)
func test(sig chan int) {
	sig <- 1
}

func main () {
	fp, err := os.OpenFile("./公寓列表.txt", os.O_CREATE | os.O_APPEND, 6)
	if err != nil {
		fmt.Println("文件打开失败。")
		return
	}
	list := 20
	abbr := "sh"
	for page :=0; ; page++ {
		detailUrl := fmt.Sprintf("https://m.ke.com/chuzu/%s/brand/pg%d/?ajax=1", abbr, page)
		info, num := spider.GetBrandInfos(spider.Download(detailUrl))
		retry := 3
		if list == 20 && num ==0 {
			for retry > 0 {
				info, num = spider.GetBrandInfos(spider.Download(detailUrl))
				if num > 0 {break}
				retry--
			}
		}
		list = num
		fmt.Printf("detailUrl:%s num:%d list: %d \n", detailUrl, num, list)
		for i := 0; i < len(info); i++ {
			fp.WriteString(info[i].GongyuId + " " + info[i].GongyuName + "\n")
		}
		if list == 0 {break}
	}
}