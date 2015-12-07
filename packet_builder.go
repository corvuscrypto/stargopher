package stargopher

import (
	"log"
	"reflect"
	"strconv"
)

var primitiveLengths = map[string]uint8{
	"uint8":  1,
	"int8":   1,
	"uint16": 2,
	"int16":  2,
	"uint32": 4,
	"int32":  4,
	"uint64": 8,
	"int64":  8,
}

//Packet ... this may not be needed...
type Packet interface{}

//These are the maps that will contain the functions that modify
//the packets and performs actions before or after handling.
//Only the packetModHandlers should modify server/client sent data
var beforeHandlers = make(map[PacketType][]func())
var packetModHandlers = make(map[PacketType][]func(Packet) (Packet, bool))
var afterHandlers = make(map[PacketType][]func())

//PacketHandler will build an appropriate packet, then
//call any associated methods. The variable passthrough will be true if
//the data is meant to be sent through. Otherwise the packet will be dropped
func PacketHandler(data []byte, payloadLength int64) ([]byte, bool) {

	ptype := PacketType(data[0])

	var passthrough = true
	//first handle the before action if exists
	for _, f := range beforeHandlers[ptype] {
		f()
	}

	packet := PacketDecoder(data, payloadLength)
	//then do the packet modifying functions
	for _, f := range packetModHandlers[ptype] {
		var rb = false
		packet, rb = f(packet)
		passthrough = passthrough && rb
	}

	//then do the after handling functions
	for _, f := range afterHandlers[ptype] {
		f()
	}

	//passthrough for now
	return data, true
}

//PacketDecoder is responsible for turning packet data into an easily
//modifiable struct to be later reencoded for transport
func PacketDecoder(data []byte, payloadLength int64) Packet {

	var packet Packet
	var slicePointer = 0

	//isolate the payload
	payload := data[len(data)-int(payloadLength):]
	ptype := PacketType(data[0])

	defer func() {
		if r := recover(); r != nil {
			log.Println("failed to build packet:", ptype, slicePointer, len(payload), r)
		}
	}()

	//build packet based on the packetRegistry
	t := packetRegistry[data[0]]
	p := reflect.New(t).Elem()
	p.Field(0).Set(reflect.ValueOf(basePacket{ID: ptype, PayloadLength: payloadLength}))
	for i := 1; i < p.NumField(); i++ {
		var length int

		f, ft := p.Field(i), t.Field(i)

		l := ft.Tag.Get("length")

		if l != "" {
			length, _ = strconv.Atoi(l)
		}

		//if lengthPrefix, set local flag
		if ft.Tag.Get("lengthPrefix") != "" {
			tf := p.Field(i - 1)
			switch tf.Type().String() {
			case "stargopher.SVLQ":
				length = int(tf.Interface().(SVLQ))
				break
			case "stargopher.VLQ":
				length = int(tf.Interface().(VLQ))
				break
			case "uint8":
				length = int(tf.Uint())
				break
			case "int64":
				length = int(tf.Int())
				break
			}
		}

		switch f.Type().String() {
		case "uint8", "uint16", "uint32", "uint64":
			var x uint64
			for j := uint8(0); j < primitiveLengths[f.Type().String()]; j++ {
				x = (x << 8) | uint64(payload[slicePointer])
				slicePointer++
			}
			f.SetUint(x)
			break
		case "int8", "int16", "int32", "int64":
			var x int64
			for j := uint8(0); j < primitiveLengths[f.Type().String()]; j++ {
				x = (x << 8) | int64(payload[slicePointer])
				slicePointer++
			}
			f.SetInt(x)
			break
		case "string":
			var x string
			if length > 0 {
				x = string(payload[slicePointer : slicePointer+int(length)])
			} else if length == 0 {
				x = string(payload[slicePointer:])
			} else {
				x = string(payload[slicePointer : len(payload)-(int(length))])
			}
			slicePointer += len(x)
			f.SetString(x)
			break
		case "[]uint8":
			var x []byte
			if length > 0 {
				x = payload[slicePointer : slicePointer+int(length)]
			} else if length == 0 {
				x = payload[slicePointer:]
			} else {
				x = payload[slicePointer : len(payload)-(int(length))]
			}
			slicePointer += len(x)
			f.SetBytes(x)
			break
		case "bool":
			var x = true
			if payload[slicePointer] == 0 {
				x = false
			}
			f.SetBool(x)
			slicePointer++
			break
		case "stargopher.SVLQ":
			var x, y = ReadSVarint(payload[slicePointer:])
			f.Set(reflect.ValueOf(SVLQ(x)))
			slicePointer += y
		case "stargopher.VLQ":
			var x, y = ReadVarint(payload[slicePointer:])
			f.Set(reflect.ValueOf(VLQ(x)))
			slicePointer += y
		}

	}
	return Packet(packet)
}

/*
func finalizePacket(p interface{}, t PacketType) {
	var err error
	switch int(t) {
	case 0:
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
	case 11:
	case 12:
	case 13:
	case 14:
	case 15:
	case 16:
	case 17:
	case 18:
	case 19:
	case 20:
	case 21:
	case 22:
	case 23:
	case 24:
	case 25:
	case 26:
	case 27:
	case 28:
	case 29:
	case 30:
	case 31:
	case 32:
	case 33:
	case 34:
	case 35:
	case 36:
	case 37:
	case 38:
	case 39:
	case 40:
	case 41:
	case 42:
	case 43:
	case 44:
	case 45:
	case 46:
	case 47:
	case 48:
	case 49:
	case 50:
	case 51:
	case 52:
	case 53:
	case 54:
	case 55:
	case 56:

	}
}*/
