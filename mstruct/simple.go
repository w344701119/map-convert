package mstruct

import (
	"errors"
	"fmt"
	"reflect"
)

type SimpleConvert struct {
	TargetMap map[string]interface{}
}

func NewConvert(m map[string]interface{}) *SimpleConvert{
	return &SimpleConvert{TargetMap:m}
}





func (s *SimpleConvert)Convert(dest interface{})(err error){
	var destType=reflect.TypeOf(dest)
	//判断是否是指针
	if destType.Kind()!=reflect.Ptr{
		err=errors.New("is not ptr")
		return
	}
	var e=destType.Elem();
	if e.Kind()!=reflect.Struct{
		err=errors.New("is not a struct")
	}
	var va =reflect.New(e).Elem()
	if(len(s.TargetMap)>0){
		for key,value:=range s.TargetMap{
			if _,b:=e.FieldByName(key); b{
				nameStruct:=va.FieldByName(key)
				if canSet:=nameStruct.CanSet();canSet{
					nameStruct.Set(reflect.ValueOf(value))
				}
			}else{
				err=fmt.Errorf("can not find %s from %s",key,e.Name())
				return
			}
		}
	}
	reflect.ValueOf(dest).Elem().Set(va)
	return
}
