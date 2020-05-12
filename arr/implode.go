package arr

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

//把切片合并成一个字符串
func Implode(list interface{}, seq string) (string, error) {
	listValue := reflect.ValueOf(list)
	//不是切片返回空
	if listValue.Kind() != reflect.Slice {
		return "", errors.New("Data must be slice")
	}
	// 获取长度
	count := listValue.Len()
	listStr := make([]string, count, count)
	for i := 0; i < count; i++ {
		v := listValue.Index(i)
		if str, err := getValue(v); err == nil {
			listStr[i] = str
		}
	}
	return strings.Join(listStr, seq), nil
}

//获取值
func getValue(value reflect.Value) (res string, err error) {
	switch value.Kind() {
	//如果是指针获取指针的值
	case reflect.Ptr:
		res, err = getValue(value.Elem())
	default:
		res = fmt.Sprint(value.Interface())
	}
	return
}
