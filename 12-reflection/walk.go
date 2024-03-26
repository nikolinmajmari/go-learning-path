package wakling

import (
	"reflect"
)

type Person struct {
	Name    string
	Profile []Profile
}

type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(string)) {
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			v, ok := val.Recv()
			if ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valCallResult := val.Call(nil)
		for _, res := range valCallResult {
			walkValue(res)
		}

	default:
		//fmt.Printf("Kind: %s", val.Kind())
		//panic("unhandled default case")
	}
	//fmt.Printf("%s\n", val.Interface())
	//if val.Kind() == reflect.Slice {
	//	for i := 0; i < val.Len(); i++ {
	//		walk(val.Index(i).Interface(), fn)
	//	}
	//	return
	//}
	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)
	//	switch field.Kind() {
	//	case reflect.String:
	//		fn(field.String())
	//	case reflect.Struct:
	//		walk(field.Interface(), fn)
	//	case reflect.Slice:
	//		walk(field.Interface(), fn)
	//	default:
	//		//panic("unhandled default case")
	//	}
	//}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
