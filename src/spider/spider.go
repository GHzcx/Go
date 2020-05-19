package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
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

type Set struct {
	CityCode string
	CityAbbr string
	HouseInfo
}
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
func getSleep() {
	time.Sleep(time.Duration(rand.Int63n(500) + 100)*time.Millisecond)

}
//填充每个房源下的字段信息
func (s Set)UpdateOnlineInfo() Set {
	//获取公寓类型 门店数量
	Url := fmt.Sprintf("https://m.ke.com/chuzu/%s/brand/%s/info/", s.CityAbbr, s.GongyuId)
	detailStr := Download(Url)
	s.GongyuType = GetGongyuType(detailStr)
	s.ShopNum = GetShopNum(detailStr)
	v := url.Values{
		"city_id": []string{s.CityCode},
		"condition": []string{"ab"+s.GongyuId},
		"request_ts": []string{strconv.FormatInt(time.Now().Unix(), 10)},
		"scene": []string{"list"},
	}
	Url = fmt.Sprintf("https://app.api.ke.com/Rentplat/v1/house/total?%s", v.Encode())
	fmt.Printf("HouseInfo url:%s", Url)
	s.HouseNum = GetBrandHouseNum(Download(Url))
	v.Set("condition","vr1ab"+s.GongyuId)
	v.Set("request_ts", strconv.FormatInt(time.Now().Unix(), 10))
	Url = fmt.Sprintf("https://app.api.ke.com/Rentplat/v1/house/total?%s", v.Encode())
	fmt.Printf("HouseVrInfo url:%s", Url)
	s.HouseVrNum = GetBrandHouseNum(Download(Url))
	return s
}
//获取urlBody
func Download (urlLink string) string {
	getSleep()
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
	REDO:
	resp, err := client.Do(req)
	retry := 10
	for err != nil && retry != 0 {
		fmt.Println(err)
		getSleep()
		resp, err = client.Do(req)
		retry--
	}
	if err != nil {
		fmt.Println(err)
		goto REDO
	}
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(err1)
		resp.Body.Close()
		goto REDO
	}

	defer resp.Body.Close()
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
//获取公寓类型
func GetGongyuType(str string) string {
	gongyuType := 0
	jizhongshiStr := `id="brandApartment"`
	fensanshiStr := `id="brandHouse"`
	if strings.Contains(str, jizhongshiStr) == true {gongyuType |= 1}
	if strings.Contains(str, fensanshiStr) == true {gongyuType |= 2}

	return Enum[string(gongyuType)]
}
//获取门店数量
func GetShopNum(str string) int {
	reg := regexp.MustCompile(`<p class="brand__content--info--desc">(.*?)</p>`)
	desc := reg.FindAllStringSubmatch(str, -1)
	reg = regexp.MustCompile(`<p class="brand__content--info--num">(.*?)</p>`)
	num := reg.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(desc) ; i++ {
		if desc[i][1] == "门店数量" {
			n, err := strconv.Atoi(num[i][1])
			if err != nil { return 0}
			return n
		}
	}
	return 0
}

func GetBrandHouseNum(str string) int {
	fmt.Println(str)
	var f interface{}
	err := json.Unmarshal([]byte(str),&f)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	m := f.(map[string]interface{})["data"].(map[string]interface{})["total"]
	return int(m.(float64))
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
func GetBrandInfoList(Queue chan int, done chan struct{},result chan<- Set,  cityCode, abbr string) {
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
		for i := 0; i < len(info); i++ {
			go func(result chan<- Set, s Set) {
				h := s.UpdateOnlineInfo()
				result <- h
			}(result, Set{CityAbbr: abbr, CityCode: cityCode, HouseInfo:info[i]})
		}
		if list == 0 {break}
		fmt.Println("Queue", len(Queue), "")
	}
	done <- struct{}{}
	<-Queue
}








