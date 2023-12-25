package utils

import "reflect"

func GetStructField(s any, field string) (any, bool) {
	sType := reflect.TypeOf(s)
	sVal := reflect.ValueOf(s)

	for i := 0; i < sType.NumField(); i++ {
		f := sType.Field(i)
		if f.Name == field {
			return sVal.FieldByName(field).Interface(), true
		}
	}

	return nil, false
}
