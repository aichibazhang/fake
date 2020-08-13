package model

import (
	"fake_data/util"
)

func AddCompanyData() {
	util.AddFuncData("company", util.Info{
		Description: "获取公司名称",
		Call:        CompanyInfo,
	})
}
func AddJobData() {
	util.AddFuncData("job", util.Info{
		Description: "获取职位名称",
		Call: func() (i interface{}, err error) {
			return Info("job")
		},
	})
}
func AddNameData() {
	util.AddFuncData("name", util.Info{
		Description: "获取人名",
		Call:        NameInfo,
	})
}
func AddPhoneData() {
	util.AddFuncData("phone", util.Info{
		Description: "获取手机号",
		Call:        PhoneInfo,
	})
}
func CompanyInfo() (interface{}, error) {
	value, err := Info("company")
	if err != nil {
		return "", err
	}
	return value.(string), nil
}
func JobInfo() (interface{}, error) {
	value, err := Info("job")
	if err != nil {
		return "", err
	}
	return value.(string), nil
}
func NameInfo() (interface{}, error) {
	value, err := Info("name")
	if err != nil {
		return "", err
	}
	return value.(string), nil
}
func PhoneInfo() (interface{}, error) {
	value, err := Info("phone")
	if err != nil {
		return "", err
	}
	return value.(string), nil
}
func Info(funcName string) (interface{}, error) {
	var value string
	var info []string
	switch funcName {
	case "company":
		info = []string{"prefix", "middle", "suffix"}
	case "job":
		info = []string{"job"}
	case "name":
		info = []string{"first_name", "last_name_male"}
	case "phone":
		info = []string{"prefix","middle", "suffix"}
	}
	for k, _ := range info {
		info, err := util.GetRandValue([]string{funcName, info[k]})
		if err != nil {
			panic(err)
		}
		info = util.ReplaceWithNumbers(info)
		value += info
	}
	return value, nil
}
