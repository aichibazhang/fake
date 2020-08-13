package util

import (
	"errors"
	"fake_data/data"
	"math/rand"
	"reflect"
)

const hashtag = '#'

var paramErr = errors.New("传入的参数不正确")

// Get Random Value
func GetRandValue(dataVal []string) (string, error) {
	if !dataExist(dataVal) {
		return "", paramErr
	}
	var randValue string
	switch val := data.Data[dataVal[0]].(type) {
	case map[string][]string:
		randValue = val[dataVal[1]][rand.Intn(len(val[dataVal[1]]))]
	case map[string]map[string][]string:
		mapKeys := reflect.ValueOf(val[dataVal[1]]).MapKeys()
		province := mapKeys[rand.Intn(len(mapKeys))].String()
		randValue = province + "," + val[dataVal[1]][province][rand.Intn(len(val[dataVal[1]][province]))]
	}
	return randValue, nil
}

func dataExist(dataVal []string) bool {
	var exists bool
	switch val := data.Data[dataVal[0]].(type) {
	case map[string][]string:
		if len(dataVal) == 2 {
			_, exists = data.Data[dataVal[0]]
			if exists {
				_, exists = val[dataVal[1]]
			}
		}
	case map[string]map[string][]string:
		if len(dataVal) == 2 {
			_, exists = data.Data[dataVal[0]]
			if exists {
				_, exists = val[dataVal[1]]
			}
		}
	}
	return exists
}
func ReplaceWithNumbers(str string) string {
	if str == "" {
		return str
	}
	byteStr := []byte(str)
	for i := 0; i < len(byteStr); i++ {
		if byteStr[i] == hashtag {
			byteStr[i] = randDigit()
		}
	}
	return string(byteStr)
}
func randDigit() byte {
	return byte(rand.Intn(10)) + '0'
}
