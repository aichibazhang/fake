package parse

import (
	"fmt"
	"github.com/aichibazhang/fake/model"
	"testing"
)

type DrugOrder struct {
	Order
	model.Address `fake:"func(AddressInfo)"`
	OrderAmt string `fake:"func(RandDate(2016-05-06,2020-05-02,datetime))"`
}
type Order struct {
	OrderInfo
	OrderCnt  string `default:"25"`
	OrderProv string `default:"25"`
}
type OrderInfo struct {
	OrderAmt string `fake:"func(RandDate(2016-05-06,2020-05-02,date))"`
}
func TestFuncParse(t *testing.T) {
	order:=DrugOrder{
		Order:    Order{},
		OrderAmt: "",
	}
	TagParse(&order)
	fmt.Println(order)
}
