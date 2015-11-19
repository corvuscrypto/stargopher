package stargopher

import "reflect"

func serializePacket(p interface{}) []byte {
	var data []byte
	values := reflect.ValueOf(p).Elem()
	for i := 1; i < values.NumField(); i++ {
		data = append(data, values.Field(i).Interface().(byte))
	}

	return []byte("")
}
