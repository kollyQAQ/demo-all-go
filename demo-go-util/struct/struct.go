package _struct

import "reflect"

func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		//这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
		dvalue.Set(value)
	}
}
