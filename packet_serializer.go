package stargopher

import (
	"bytes"
	"compress/zlib"
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
			var holder = make([]byte, primitiveLengths[t])
			var num uint64
			if t[0] == 'i' {
				num = uint64(field.Int())
			} else {
				num = field.Uint()
			}
			if num == 0 {
				holder = []byte{}
			}

			for i := range holder {
				holder[len(holder)-i-1] = uint8(num & 0xff)
				num >>= 8
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
	return data
}
