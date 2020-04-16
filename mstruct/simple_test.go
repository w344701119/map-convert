package mstruct

import (
	"testing"
)


type DataType struct {
	Id int
	Name string
	IsTest bool
	CreateTime string
	UpdateTime string
	TimeStamp  int

}

func TestConvert(t *testing.T){

	var m =map[string]interface{}{"Id":1,"Name":"测试111","IsTest":true,"CreateTime":"2020-04-16 17:00:34","UpdateTime":"2020-04-16 17:00:34","TimeStamp":1587027634}
	var c =NewConvert( m)

	var dest DataType
	err:=c.Convert(&dest)
	t.Log(err)
	t.Log(dest)

}


func BenchmarkConvert(b *testing.B) {

	for i:=0;i<b.N ;i++  {
		var m =map[string]interface{}{"Id":1,"Name":"测试111","IsTest":true,"CreateTime":"2020-04-16 17:00:34","UpdateTime":"2020-04-16 17:00:34","TimeStamp":1587027634}
		var c =NewConvert( m)
		var dest DataType
		_=c.Convert(&dest)
	}

}