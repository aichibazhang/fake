package parse

import (
	"fmt"
	"github.com/aichibazhang/fake/model"
	"github.com/aichibazhang/fake/util"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	funcRegex       = regexp.MustCompile(`func\(([a-zA-Z]*)`)
	randDateRegx    = regexp.MustCompile(`RandDate\(([a-zA-Z0-9,-]*)\)`)
	randIntRegx     = regexp.MustCompile(`RandIntRangeBetween\(([a-zA-Z0-9,-]*)\)`)
	randFloatRegx   = regexp.MustCompile(`RandFloatRangeRand\(([a-zA-Z0-9,-]*)\)`)
	randIntRandRegx = regexp.MustCompile(`RandIntRangeRand\(([a-zA-Z0-9,-]*)\)`)
	codeRegx        = regexp.MustCompile(`CodeInfo\(([^)]*)\)`)
	codeParamRegx   = regexp.MustCompile(`\{(.*)\}`)
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
			paramType := v.FieldByName(fieldInfo.Name)
			switch paramType.Kind() {
			case reflect.String:
				v.FieldByName(fieldInfo.Name).SetString(defaultTag)
			case reflect.Int64:
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
					retList = paramFakeFunc(randDateRegx, fakeTag, util.RandDate)
				case "RandIntRangeBetween":
					retList = paramFakeFunc(randIntRegx, fakeTag, util.RandIntRangeBetween)
				case "RandFloatRangeRand":
					retList = paramFakeFunc(randFloatRegx, fakeTag, util.RandFloatRangeRand)
				case "RandIntRangeRand":
					retList = paramFakeFunc(randIntRandRegx, fakeTag, util.RandFloatRangeRand)
				case "DistrictInfo":
					retList = noParamFakeFunc(model.DistrictInfo)
				case "CityInfo":
					retList = noParamFakeFunc(model.CityInfo)
				case "DetailAddress":
					retList = noParamFakeFunc(model.DetailAddress)
				case "CompanyInfo":
					retList = noParamFakeFunc(model.CompanyInfo)
				case "JobInfo":
					retList = noParamFakeFunc(model.JobInfo)
				case "NameInfo":
					retList = noParamFakeFunc(model.NameInfo)
				case "PhoneInfo":
					retList = noParamFakeFunc(model.PhoneInfo)
				case "CodeInfo":
					retList = codeInfoFunc(codeRegx, fakeTag, model.CodeInfo)
				}
				filed := v.FieldByName(fieldInfo.Name)
				if filed.CanSet() {
					switch filed.Kind() {
					case reflect.String:
						value := retList[0].Interface().(string)
						fmt.Println(value)
						v.FieldByName(fieldInfo.Name).SetString(value)
					case reflect.Int64:
						value := retList[0].Interface().(int64)
						v.FieldByName(fieldInfo.Name).SetInt(value)
					default:
						v.FieldByName(fieldInfo.Name).Set(retList[0])
					}
				}
			}
		}
		parse(t.Field(i).Type, v.Field(i))
	}
}
func noParamFakeFunc(i interface{}) []reflect.Value {
	funcValue := reflect.ValueOf(i)
	return funcValue.Call(nil)
}
func codeInfoFunc(regexp *regexp.Regexp, fakeTag string, i interface{}) []reflect.Value {
	funcValue := reflect.ValueOf(i)
	funcMatch := regexp.FindStringSubmatch(fakeTag)[1]
	index := strings.LastIndex(funcMatch, ",")
	var param [2]string
	paramString := funcMatch[0:index]
	paramMatch := codeParamRegx.FindStringSubmatch(paramString)[1]
	params := strings.Split(paramMatch, ",")
	param[1] = strings.TrimSpace(funcMatch[index+1:])
	var paramList []reflect.Value
	paramList = append(paramList, reflect.ValueOf(params))
	paramList = append(paramList, reflect.ValueOf(param[1]))
	return funcValue.Call(paramList)
}
func paramFakeFunc(regexp *regexp.Regexp, fakeTag string, i interface{}) []reflect.Value {
	funcValue := reflect.ValueOf(i)
	funcMatch := regexp.FindStringSubmatch(fakeTag)[1]
	param := strings.SplitN(funcMatch, ",", reflect.TypeOf(i).NumIn())
	var paramList []reflect.Value
	for k, v := range param {
		switch reflect.TypeOf(i).In(k).Kind() {
		case reflect.Int64:
			value, _ := strconv.Atoi(v)
			paramList = append(paramList, reflect.ValueOf(int64(value)))
		default:
			paramList = append(paramList, reflect.ValueOf(v))
		}
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
