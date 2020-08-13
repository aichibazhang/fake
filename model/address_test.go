package model

import (
	"fake_data/util"
	"fmt"
	"testing"
)

func init() {
	util.Seed(0)
	AddAddressFunc()
}
func TestAddress(t *testing.T) {
	value1, _ := AddressInfo()
	fmt.Println(value1)
	value2, err := PhoneInfo()
	fmt.Println(value2, err)
	//for i := 0; i < 300; i++ {
	//	value3, err := DetailAddress()
	//	fmt.Println(value3, err)
	//}
	value3, _ := CodeInfo([]string{"districts", "streets"}, "address")
	t.Log(value3)
	value4, err := CodeInfo([]string{"cityInfo"}, "city")
	t.Log(value4,err)
	value5, err := CodeInfo([]string{"prefix","middle"}, "company")
	t.Log(value5,err)
	value6, err := CodeInfo([]string{"middle"}, "phone")
	t.Log(value6,err)
}
func TestAddDataFunc(t *testing.T) {
	value, _ := util.GetFuncLookup("detailAddress").Call()
	t.Log(value)
	value, _ = util.GetFuncLookup("city").Call()
	t.Log(value)
	value, _ = util.GetFuncLookup("address").Call()
	t.Log(value)
	value, _ = util.GetFuncLookup("address").Call()
	t.Log(value)
}
func BenchmarkAddAddressFunc(b *testing.B) {
	b.ReportAllocs()
	for i:=0;i<b.N;i++{
		AddressInfo()
	}
}
