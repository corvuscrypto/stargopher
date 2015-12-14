package stargopher

import (
	"bytes"
	"compress/zlib"
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
	var values reflect.Value
	if reflect.TypeOf(p).String() != "reflect.Value" {
		values = reflect.ValueOf(p).Elem()
	} else {
		values = p.(reflect.Value)
	}
	for i := 0; i < values.NumField(); i++ {
		if i == 0 {
			pid = byte(values.Field(i).Interface().(basePacket).ID)
			continue
		}
		field := values.Field(i)
		t := field.Type().String()
		switch t {
		case "string":
			data = append(data, []byte(field.String())...)
			break
		case "[]uint8":
			data = append(data, field.Bytes()...)
			break
		case "stargopher.SVLQ":
			var num = field.Int() * 2
			if num < 0 {
				num--
			}
			data = append(data, WriteVarint(num)...)
			break
		case "stargopher.VLQ":
			var num = int64(field.Uint())
			if num < 0 {
				num--
			}
			data = append(data, WriteVarint(num)...)
			break
		default:
			fmt.Println(t)
			var holder []byte
			var num uint64
			var i uint8
			if t[0] == 'i' {
				num = uint64(field.Int())
			} else {
				num = field.Uint()
			}
			if num == 0 {
				i++
				holder = []byte{0}
			}

			for num > 0 {
				holder = append([]byte{uint8(num & 0xff)}, holder...)
				num >>= 8
				i++
			}
			for primitiveLengths[t] > i {
				holder = append([]byte{0}, holder...)
				i++
			}
			data = append(data, holder...)
			break
		}
	}

	var b bytes.Buffer
	zw := zlib.NewWriter(&b)
	zw.Write(data)
	zw.Close()
	var _sub = 0
	if len(data) > b.Len() {
		data = b.Bytes()
		_sub = 1
	}

	length := WriteVarint(int64((len(data) * 2) - _sub))
	data = append(length, data...)
	data = append([]byte{pid}, data...)
	fmt.Println(data)
	return data
}
