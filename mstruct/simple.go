package mstruct

import (
	map_convert "convert"
	"fmt"
	"reflect"
	"time"
)

type SimpleConvert struct {
	TargetMap map[string]interface{}
}

func NewConvert(m map[string]interface{}) *SimpleConvert {
	return &SimpleConvert{TargetMap: m}
}

var (
	ErrorNotPtr    = fmt.Errorf("is not ptr")
	ErrorNotStruct = fmt.Errorf("is not a struct")
	ErrorNothing   = fmt.Errorf("nothing to set")
)

//把简单的map转换成结构
func (s *SimpleConvert) Convert(dest interface{}) error {
	var t = reflect.TypeOf(dest)
	//判断是否是指针
	if t.Kind() != reflect.Ptr {
		return ErrorNotPtr
	}
	var e = t.Elem()
	if e.Kind() != reflect.Struct {
		return ErrorNotStruct
	}
	if len(s.TargetMap) == 0 {
		return ErrorNothing
	}
	for key, value := range s.TargetMap {
		key = map_convert.HumpKey(key)
		if _, b := e.FieldByName(key); !b {
			return fmt.Errorf("can not find %s from %s", key, e.Name())
		}
		ns := reflect.ValueOf(dest).Elem().FieldByName(key)
		vt := reflect.ValueOf(value).Kind()
		if ns.Kind() == reflect.Struct {
			if ns.Type() == reflect.TypeOf(time.Time{}) {
				if tm, err := map_convert.TimeConvert(value); err == nil {
					ns.Set(reflect.ValueOf(tm))
					continue
				} else {
					return err
				}
			}
		}
		if ns.Kind() != vt {
			return fmt.Errorf("cann't use the %s as %s in the %s field", ns.Kind(), vt, key)
		}
		//判断是否可以设置值
		if ns.CanSet() {
			ns.Set(reflect.ValueOf(value))
		} else {
			return fmt.Errorf("%s can't  set", key)
		}
	}
	return nil
}
