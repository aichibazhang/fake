package util

import "sync"

var FuncData map[string]Info
var mul sync.Mutex

type Info struct {
	Description string
	Call        func() (interface{}, error)
}

func AddFuncData(funcName string, info Info) {
	if FuncData == nil {
		FuncData = make(map[string]Info)
	}
	mul.Lock()
	FuncData[funcName] = info
	mul.Unlock()
}
func GetFuncLookup(functionName string) *Info {
	info, ok := FuncData[functionName]
	if !ok {
		return nil
	}

	return &info
}
func RemoveFuncData(funcName string) {
	_, ok := FuncData[funcName]
	if !ok {
		return
	}
	mul.Lock()
	delete(FuncData, funcName)
	mul.Unlock()
}
