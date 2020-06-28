package map_convert

import (
	"errors"
	"reflect"
	"strings"
	"time"
)

var (
	ErrorAssertion = errors.New("fail to assertion")
	TimeLayout     = "2006-01-02 15:04:05"
)

//key转换成驼峰行
//只针对第一个字母和"_"后的一个字母转换
func HumpKey(key string) string {
	keyArr := strings.Split(key, "_")
	for key, value := range keyArr {
		r := []rune(value)
		if len(r) > 0 {
			if r[0] >= 97 && r[0] <= 122 {
				r[0] -= 32
			}
		}
		keyArr[key] = string(r)
	}
	return strings.Join(keyArr, "")
}

//时间转换为时间对象
//t时间支持整型，长整型，字符串（2006-01-02 15:04:05）
func TimeConvert(t interface{}) (time.Time, error) {
	v := reflect.ValueOf(t)
	switch v.Type().Kind() {
	case reflect.String:
		if strT, oK := t.(string); oK {
			return time.ParseInLocation(TimeLayout, strT, time.Local)
		} else {
			return time.Now(), ErrorAssertion
		}
	case reflect.Int:
		if intT, oK := t.(int); oK {
			return time.Unix(int64(intT), 0), nil
		} else {
			return time.Now(), ErrorAssertion
		}
	case reflect.Int64:
		if int64T, oK := t.(int64); oK {
			return time.Unix(int64T, 0), nil
		} else {
			return time.Now(), ErrorAssertion
		}
	}
	return time.Now(), errors.New("not supper type")
}
