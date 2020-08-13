package model

import (
	"fake_data/util"
	"testing"
)

func init() {
	util.Seed(0)
	AddCompanyData()
	AddJobData()
	AddNameData()
	AddPhoneData()

}
func TestCompanyInfo(t *testing.T) {
	value, err := CompanyInfo()
	t.Log(value, err)
	value, err = JobInfo()
	t.Log(value, err)
	value, err = NameInfo()
	t.Log(value, err)
	value, err = PhoneInfo()
	t.Log(value, err)
}
func TestAddCompanyData(t *testing.T) {
	value, _ := util.GetFuncLookup("company").Call()
	t.Log(value)
	value, _ = util.GetFuncLookup("job").Call()
	t.Log(value)
	value, _ = util.GetFuncLookup("name").Call()
	t.Log(value)
	value, _ = util.GetFuncLookup("phone").Call()
	t.Log(value)
}
