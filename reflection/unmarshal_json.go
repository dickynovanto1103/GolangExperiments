package main

import (
	"encoding/json"
	"log"
	"reflect"
)

// oh, we cannot get numField from ptr value, it will panic
// oh I learnt that if we want to get field name, we cannot do reflect.ValueOf first, since it will return
// Value...we must do reflect.TypeOf such that it can return type
func UnmarshalJSON(data []byte, v interface{}) {
	typ := reflect.TypeOf(v)
	log.Printf("typ: %+v", typ)
	nonPtrTyp := typ.Elem()
	log.Printf("elem: %+v numField: %v", nonPtrTyp, nonPtrTyp.NumField())

	fieldNameToValueMap := make(map[string]reflect.Value)

	valueElem := reflect.ValueOf(v).Elem()

	for i := 0; i < nonPtrTyp.NumField(); i++ {
		field := nonPtrTyp.Field(i)
		jsonTag := field.Tag.Get("json")
		fieldNameToValueMap[jsonTag] = valueElem.Field(i)

		log.Printf("field: %+v", field)
	}

	log.Printf("fieldName to valueMap: %+v", fieldNameToValueMap)

	tempMap := make(map[string]interface{})
	if err := json.Unmarshal(data, &tempMap); err != nil {
		log.Panicf("error in unmarshaling data, err: %v", err)
	}

	log.Printf("tempMap: %+v", tempMap)

	//for fieldName, val := range fieldNameToValueMap {
	//	if _, ok := tempMap[fieldName]; !ok {
	//		continue
	//	}
	//
	//	if val.IsValid() && val.CanSet() {
	//		val.Set(reflect.ValueOf(tempMap[fieldName]).Convert(val.Type()))
	//	}
	//
	//}

	for key, val := range tempMap {
		fieldValue, ok := fieldNameToValueMap[key]
		if !ok {
			continue
		}

		value := reflect.ValueOf(val)
		log.Printf("key: %v val: %v, valueType: %v, fieldValue Type: %v", key, val, value.Type(), fieldValue.Type())

		if value.Type().ConvertibleTo(fieldValue.Type()) {
			fieldValue.Set(value.Convert(fieldValue.Type()))
		}
	}
}
