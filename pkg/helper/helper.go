package helper

import (
	"encoding/json"
	"os"
	"reflect"
)

func GetStructFields(obj any, tag string) (fileds []string, err error) {
	reflection := reflect.ValueOf(obj)

	for i := 0; i < reflection.Type().NumField(); i++ {
		field := reflection.Type().Field(i)
		switch val := field.Tag.Get(tag); val {
		case "-", "":
			continue
		default:
			fileds = append(fileds, val)
		}
	}

	return
}

func PrintJson(obj any) {
	_ = json.NewEncoder(os.Stdout).Encode(obj)
}
