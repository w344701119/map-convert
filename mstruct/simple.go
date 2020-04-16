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

//
//func (s *SimpleConvert)Convert(dest interface{})(error){
//	var t=reflect.TypeOf(dest)
//	//判断是否是指针
//	if t.Kind()!=reflect.Ptr{
//		return errors.New("is not ptr")
//
//	}
//	var e=t.Elem();
//	if e.Kind()!=reflect.Struct{
//		return errors.New("is not a struct")
//	}
//
//	if(len(s.TargetMap)==0){
//		return errors.New("nothing to set")
//	}
//	var va =reflect.New(e).Elem()
//
//	for key,value:=range s.TargetMap{
//		if _,b:=e.FieldByName(key); !b{
//			return fmt.Errorf("can not find %s from %s",key,e.Name())
//		}
//		ns:=va.FieldByName(key)
//		//判断是否可以设置值
//		if ca:=ns.CanSet();ca{
//			ns.Set(reflect.ValueOf(value))
//		}else{
//			return fmt.Errorf("%s can't  set",key)
//
//		}
//
//	}
//	reflect.ValueOf(dest).Elem().Set(va)
//	return nil
//}


func (s *SimpleConvert)Convert(dest interface{})(error){
	var t=reflect.TypeOf(dest)
	//判断是否是指针
	if t.Kind()!=reflect.Ptr{
		return errors.New("is not ptr")
	}
	var e=t.Elem();
	if e.Kind()!=reflect.Struct{
		return errors.New("is not a struct")
	}
	if(len(s.TargetMap)==0){
		return errors.New("nothing to set")
	}
	for key,value:=range s.TargetMap{
		if _,b:=e.FieldByName(key); !b{
			return fmt.Errorf("can not find %s from %s",key,e.Name())
		}
		ns:=reflect.ValueOf(dest).Elem().FieldByName(key)
		//判断是否可以设置值
		if ca:=ns.CanSet();ca && ns.Kind()==reflect.ValueOf(value).Kind(){
			ns.Set(reflect.ValueOf(value))
		}else{
			return fmt.Errorf("%s can't  set",key)
		}
	}
	return nil
}
