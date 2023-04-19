package jsonx

import (
	"encoding/json"
	"reflect"
)

func Equal(jsonStr1, jsonStr2 string) (bool, error) {
	var i interface{}
	var i2 interface{}
	err := json.Unmarshal([]byte(jsonStr1), &i)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal([]byte(jsonStr2), &i2)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(i, i2), nil
}

func Equals(jsonStr1, jsonStr2 string, moreJsonStr ...string) (bool, error) {
	equal, err := Equal(jsonStr1, jsonStr2)
	if err != nil {
		panic(err)
		return false, err
	}

	if !equal {
		return false, nil
	}
	if len(moreJsonStr) >= 1 {
		for _, v := range moreJsonStr {
			eq, err := Equal(v, jsonStr1)
			if err != nil {
				panic(err)
				return false, err
			}
			if !eq {
				return false, nil
			}
		}
	}
	return true, nil
}
