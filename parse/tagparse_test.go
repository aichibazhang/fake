package parse

import (
	"fmt"
	"github.com/aichibazhang/fake/model"
	"github.com/aichibazhang/fake/util"
	"testing"
)

//注解优先级:
//	defalut>fake
type DrugOrder struct {
	Order
	model.Address `fake:"func(AddressInfo)"`
	OrderAmt      string `fake:"func(RandDate(2016-05-06,2020-05-02,datetime))"`
	City          string `fake:"func(CityInfo)"`
	DistrictInfo  string `fake:"func(DistrictInfo)"`
	DetailAddress string `fake:"func(DetailAddress)"`
	Code
}
type Order struct {
	OrderInfo
	OrderCnt  string `default:"father"`
	OrderProv int64  `default:"25"`
}
type OrderInfo struct {
	OrderAmt    string `fake:"func(RandDate(2016-05-06,2020-05-02,date))"`
	Amt         int64  `fake:"func(RandIntRangeBetween(30000,50000))"`
	AmtPercent  int64  `fake:"func(RandIntRangeRand(11,99))"`
	CompanyInfo string `fake:"func(CompanyInfo)"`
	JobInfo     string `fake:"func(JobInfo)"`
	NameInfo    string `fake:"func(NameInfo)"`
	PhoneInfo   string `fake:"func(PhoneInfo)"`
}
type Code struct {
	CodeInfo  string `fake:"func(CodeInfo([]string{prefix,middle}, company))"`
	PhoneInfo string `fake:"func(CodeInfo([]string{middle}, phone))"`
}
type Number struct {
	Count   int64   `fake:"func(RandIntRangeBetween(30000,50000))"`
	Percent float64 `fake:"func(RandFloatRangeRand(4))"`
	Code    string  `fake:"func(RandEnum(110000,120000,130000,150000,170000))"`
}

func init() {
	util.Seed(0)
}
func TestFuncParse(t *testing.T) {
	//order := DrugOrder{}
	//TagParse(&order)
	//fmt.Println(order)
	number := Number{}
	for i := 0; i < 20; i++ {
		TagParse(&number)
		fmt.Println(number)
	}
}
