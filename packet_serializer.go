package stargopher

import (
	"fmt"
	"reflect"
)

//SerializePacket takes a packet and returns a byte array for
//transport across TCP
func SerializePacket(p interface{}, padding int) []byte {
	var data []byte
	if padding > 0 {
		data = make([]byte, padding)
	}
	var pid byte
	values := reflect.ValueOf(p).Elem()
	for i := 0; i < values.NumField(); i++ {
		if i == 0 {
			pid = byte(values.Field(i).Interface().(PacketType))
			continue
		}
		field := values.Field(i)
		t := field.Type().String()
		fmt.Println(t)
		switch t {
		case "string":
			data = append(data, []byte(field.Interface().(string))...)
			break
		case "[]uint8":
			data = append(data, field.Bytes()...)
			break
		default:
			var holder []byte
			var num uint64
			if t[0] == 'i' {
				num = uint64(field.Int())
			} else {
				num = field.Uint()
			}
			if num == 0 {
				holder = []byte{0}
			}
			for num > 0 {
				holder = append([]byte{uint8(num & 0xff)}, holder...)
				num >>= 8
			}
			data = append(data, holder...)
			break
		}
	}
	length := WriteVarint(int64(len(data) * 2))
	fmt.Println(int64(len(data)))
	fmt.Println(length)
	data = append(length, data...)
	data = append([]byte{pid}, data...)
	return data
}
