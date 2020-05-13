package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type HouseInfo struct {
	GongyuName string
	GongyuType string
	GongyuId int
	HouseNum int
	HouseVrNum int
	ShopNum int
}

type AllData map[string][]HouseInfo

//获取代理IP
func GetProxyUrl() string {
	proxyUrl := "http://10.132.66.8:1700/proxy.php"
	resp, err := http.Get(proxyUrl)
	for err != nil {
		resp, err = http.Get(proxyUrl)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read faild")
	}
	var f []map[string]string
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Println("json decode faild")
	}
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(len(f) - 1))
	str := strings.ToLower(f[num]["proxy_type"]) +"://" + f[num]["ip"] + ":" + f[num]["port"]
	return str
}




