package spider

type HouseInfo struct {
	GongyuName string
	GongyuType string
	GongyuId int
	HouseNum int
	HouseVrNum int
	ShopNum int
}

type AllData map[string][]HouseInfo

