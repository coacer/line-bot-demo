package utils

import "reflect"

func GetKeys(s interface{}) (keys []string) {
	t := reflect.TypeOf(s)
	elem := reflect.ValueOf(s)
	cnt := elem.NumField()
	keys = make([]string, cnt)
	for i := 0; i < cnt; i++ {
		keys[i] = t.Field(i).Name
	}
	return
}
