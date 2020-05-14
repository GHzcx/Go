package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type HouseInfo struct {
	GongyuName string
	GongyuType string
	GongyuId string
	HouseNum int
	HouseVrNum int
	ShopNum int
}

type AllData map[string][]HouseInfo //映射关系 城市 =》公寓信息

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

//填充每个房源下的字段信息
func (h HouseInfo)UpdateOnlineInfo (abbr string) {

}


//获取urlBody
func Download (urlLink string) string {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(GetProxyUrl())
	}
	client := &http.Client{
		Transport: &http.Transport{Proxy: proxy},
		Timeout: 3 * time.Second,
	}
	req,_ := http.NewRequest(http.MethodGet, urlLink, nil)
	req.Header.Set("Cookie", "lianjia_uuid=D54C855E-409C-4938-BE3C-9C0923F8AF58; lianjia_ssid=6EC4B1B6-BBA8-4396-8DA6-7C966DF2C877; lianjia_udid=A253C17E-5837-4AF8-8523-06F3AF7C5851; lianjia_token=2.003a4d930e46e84eaf2be0ba3fb4ad3b3f")
	req.Header.Add("Lianjia-Access-Token", "2.003a4d930e46e84eaf2be0ba3fb4ad3b3f")
	req.Header.Add("Lianjia-Device-Id", "A253C17E-5837-4AF8-8523-06F3AF7C5851")
	req.Header.Add("Authorization", "MjAxODAxMTFfaW9zOjk4MDVjNGUwZDI1MWZhNWE3ZjU2NTE0NTNlNjM0OWM3MTg2MzllODY=")
	resp, err := client.Do(req)
	retry := 3
	for err != nil && retry != 0 {
		resp, err = client.Do(req)
		fmt.Println(err)
		retry--
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	retry = 3
	for err != nil && retry > 0 {
		body, err = ioutil.ReadAll(resp.Body)
		fmt.Println("read body error")
	}
	return string(body)
}
// 获取公寓ID 公寓名称
func GetBrandIdAndName (str string) (string, string) {
	reg := regexp.MustCompile(`<a href="/chuzu/.*?/brand/(.*?)/`)
	brandId := reg.FindStringSubmatch(str)[1]
	reg = regexp.MustCompile(`<p class="oneline content__item__title">([\s\S]*?)</p>`)
	brandName := reg.FindStringSubmatch(str)[1]
	return strings.Fields(brandId)[0],strings.Fields(brandName)[0]
}
//获取公寓列表信息
func GetBrandInfos(str string) ([]HouseInfo, int) {
	h := make([]HouseInfo, 0)
	reg := regexp.MustCompile(`<div class="content__item ">([\s\S]*?)</div>`)
	for _, item := range reg.FindAllString(str, -1) {
		brandId, brandName := GetBrandIdAndName(item)
		h = append(h, HouseInfo{GongyuId: brandId, GongyuName: brandName})
	}
	return h, len(h)
}

//获取城市公寓信息
func GetBrandInfoList(cityCode, abbr string) []HouseInfo {
	h := make([]HouseInfo, 0)
	//默认提取数量为20
	list := 20
	for page :=0; ; page++ {
		detailUrl := fmt.Sprintf("https://m.ke.com/chuzu/%s/brand/pg%d/?ajax=1", abbr, page)
		info, num := GetBrandInfos(Download(detailUrl))
		retry := 3
		if list == 20 && num ==0 {
			for retry > 0 {
				info, num = GetBrandInfos(Download(detailUrl))
				if num > 0 {break}
				retry--
			}
		}
		list = num
		fmt.Printf("detailUrl:%s num:%d list: %d \n", detailUrl, num, list)
		//可以在此增加通道
		for i :=0; i < len(info); i++ {
			h = append(h, info[i])
		}
		if list == 0 {break}
	}
	return h
}








