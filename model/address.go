package model

import (
	"fake_data/util"
	"strings"
)

type Address struct {
	Province  string
	City      string
	Districts string
	Streets   string
}

func AddAddressFunc() {
	util.AddFuncData("address", util.Info{
		Description: "获取街道小区信息",
		Call:        DistrictInfo,
	})
	util.AddFuncData("city", util.Info{
		Description: "获取省市名称",
		Call:        CityInfo,
	})
	util.AddFuncData("detailAddress", util.Info{
		Description: "获取详细地址",
		Call:        DetailAddress,
	})
}
func DistrictInfo() (interface{}, error) {
	var value string
	info := []string{"districts", "streets"}
	for k, _ := range info {
		info, err := util.GetRandValue([]string{"address", info[k]})
		if err != nil {
			panic(err)
		}
		info = util.ReplaceWithNumbers(info)
		value += info
	}
	return value, nil
}

// 通过代码传递参数的形式获取自己想要的东西,比如只想要districts,传第一个参数为districts,第二个参数为address函数名称
func CodeInfo(info []string, funcName string) (interface{}, error) {
	var value string
	for k, _ := range info {
		info, err := util.GetRandValue([]string{funcName, info[k]})
		if err != nil {
			return "", err
		}
		info = util.ReplaceWithNumbers(info)
		value += info
	}
	return value, nil
}

func CityInfo() (interface{}, error) {
	value, err := util.GetRandValue([]string{"city", "cityInfo"})
	info := strings.Split(value, ",")
	if len(info) == 2 {
		return info[0] + info[1], nil
	}
	return "", err
}
func DetailAddress() (interface{}, error) {
	addr, err := DistrictInfo()
	if err != nil {
		return "", err
	}
	city, err := CityInfo()
	if err != nil {
		return "", err
	}
	value := city.(string) + addr.(string)
	return value, nil
}

func AddressInfo() (Address, error) {
	var address Address
	value, err := util.GetRandValue([]string{"city", "cityInfo"})
	if err != nil {
		return address, err
	}
	info := strings.Split(value, ",")
	if len(info) == 2 {
		address.Province = info[0]
		address.City = info[1]
	}
	districts, err := CodeInfo([]string{"districts"}, "address")
	if err != nil {
		return address, err
	}
	address.Districts = districts.(string)
	streets, err := CodeInfo([]string{"streets"}, "address")
	if err != nil {
		return address, err
	}
	address.Streets = streets.(string)
	return address, err
}
