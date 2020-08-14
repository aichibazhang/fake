package parse

import (
	"errors"
	"github.com/aichibazhang/fake/model"
	"github.com/aichibazhang/fake/util"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	funcRegex    = regexp.MustCompile(`func\(([a-zA-Z]*)`)
	randDateRegx = regexp.MustCompile(`RandDate\(([a-zA-Z0-9,-]*)\)`)
	randIntRegx  = regexp.MustCompile(`RandIntRangeBetween\(([a-zA-Z0-9,-]*)\)`)
	funcError    = errors.New("参数填写错误")
)

// 解析标签中自定义tag:func对应函数,default对应默认值
func TagParse(in interface{}) {
	parse(reflect.TypeOf(in), reflect.ValueOf(in))
}

func parse(inType reflect.Type, inValue reflect.Value) {
	switch inType.Kind() {
	case reflect.Ptr:
		ptr(inType, inValue)
	case reflect.Struct:
		pStruct(inType, inValue)
	}
}

func pStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		if defaultTag, ok := tag.Lookup("default"); ok {
			paramType := reflect.TypeOf(fieldInfo.Name)
			switch paramType.Kind() {
			case reflect.String:
				v.FieldByName(fieldInfo.Name).SetString(defaultTag)
			case reflect.Int:
				defaultInt, _ := strconv.Atoi(defaultTag)
				v.FieldByName(fieldInfo.Name).SetInt(int64(defaultInt))
			}
		} else if fakeTag, ok := tag.Lookup("fake"); ok {
			if strings.Contains(fakeTag, "func") {
				tagMatch := funcRegex.FindStringSubmatch(fakeTag)[1]
				var retList []reflect.Value
				switch tagMatch {
				case "AddressInfo":
					retList = noParamFakeFunc(model.AddressInfo)
				case "RandDate":
					retList = paramFakeFunc( randDateRegx, fakeTag, util.RandDate)
				case "RandIntRangeBetween":
					retList = paramFakeFunc(randDateRegx, fakeTag, util.RandIntRangeBetween)
				}
				v.FieldByName(fieldInfo.Name).Set(retList[0])
			}
		}
		parse(t.Field(i).Type, v.Field(i))
	}
}
func noParamFakeFunc(i interface{}) []reflect.Value {
	funcValue := reflect.ValueOf(i)
	return funcValue.Call(nil)
}
func paramFakeFunc(regexp *regexp.Regexp, fakeTag string, i interface{}) []reflect.Value {
	funcValue := reflect.ValueOf(i)
	funcMatch := regexp.FindStringSubmatch(fakeTag)[1]
	param := strings.Split(funcMatch, ",")
	var paramList []reflect.Value
	for _, v := range param {
		paramList = append(paramList, reflect.ValueOf(v))
	}
	return funcValue.Call(paramList)
}

func ptr(inType reflect.Type, value reflect.Value) {
	ele := inType.Elem()
	if value.IsNil() {
		nv := reflect.New(ele)
		parse(ele, value.Elem())
		value.Set(nv)
	} else {
		parse(ele, value.Elem())
	}
}
