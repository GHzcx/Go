package main

import (
	"regexp"
	"spider"
)
func GetBrandId(str string) string {
	reg := regexp.MustCompile(`<a href="/chuzu/.*?/brand/(.*?)/`)
	return reg.FindAllString(str, -1)[0]
}

func main() {
	spider.GetBrandInfoList(`110000`, `bj`)
	//fmt.Println(len(info),info)

}