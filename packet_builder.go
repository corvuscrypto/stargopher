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

//PacketHandler will build an appropriate packet, then
//call any associated methods. The variable passthrough will be true if
//the data is meant to be sent through. Otherwise the packet will be dropped
func PacketHandler(uid string, pc chan []byte, data []byte, payloadLength int64) {

	ptype := PacketType(data[0])

	if int64(ptype)|payloadLength == 0 {
		pc <- data
		return
	}

	var passthrough = true
	//first handle the before action if exists
	for _, f := range beforeHandlers[ptype] {
		f(uid)
	}
	packet := PacketDecoder(data, payloadLength)
	//then do the packet modifying functions
	for _, f := range packetModHandlers[ptype] {
		var rb = false
		packet, rb = f(packet)
		passthrough = passthrough && rb
	}
	//passthrough for now

	pc <- SerializePacket(packet, 0)
	return
	//then do the after handling functions
	for _, f := range afterHandlers[ptype] {
		f(uid)
	}

}

//PacketDecoder is responsible for turning packet data into an easily
//modifiable struct to be later reencoded for transport
func PacketDecoder(data []byte, payloadLength int64) Packet {

	var slicePointer = 0

	//isolate the payload
	payload := data[len(data)-int(payloadLength):]
	ptype := PacketType(data[0])

	defer func() {
		if r := recover(); r != nil {
			log.Println("failed to build packet:", data, payloadLength, ptype, slicePointer, len(payload), r)
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
		case "[]uint8", "[]byte":
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

	return p
}

/*
func finalizePacket(p interface{}, t PacketType) interface{} {
	switch int(t) {
	case 0:
		return p.(protocolVersion)
	case 1:
		return p.(serverDisconnect)
	case 2:
		return p.(connectSuccess)
	case 3:
		return p.(connectFailure)
	case 4:
		return p.(handshakeChallenge)
	case 5:
		return p.(chatReceived)
	case 6:
		return p.(timeUpdate)
	case 7:
		return p.(celestialResponse)
	case 8:
		return p.(playerWarpResult)
	case 9:
		return p.(clientConnect)
	case 10:
		return p.(clientDisconnectRequest)
	case 11:
		return p.(handshakeResponse)
	case 12:
		return p.(playerWarp)
	case 13:
		return p.(flyShip)
	case 14:
		return p.(chatSent)
	case 15:
		return p.(celestialRequest)
	case 16:
		return p.(clientContextUpdate)
	case 17:
		return p.(worldStart)
	case 18:
		return p.(worldStop)
	case 19:
		return p.(centralStructureUpdate)
	case 20:
		return p.(tileArrayUpdate)
	case 21:
		return p.(tileUpdate)
	case 22:
		return p.(tileLiquidUpdate)
	case 23:
		return p.(tileDamageUpdate)
	case 24:
		return p.(tileModificationFailure)
	case 25:
		return p.(giveItem)
	case 26:
		return p.(containerSwapResult)
	case 27:
		return p.(environmentUpdate)
	case 28:
		return p.(entityInteractResult)
	case 29:
		return p.(updateTileProtection)
	case 30:
		return p.(modifyTileList)
	case 31:
		return p.(damageTileGroup)
	case 32:
		return p.(collectLiquid)
	case 33:
		return p.(requestDrop)
	case 34:
		return p.(spawnEntity)
	case 35:
		return p.(entityInteract)
	case 36:
		return p.(connectWire)
	case 37:
		return p.(disconnectAllWires)
	case 38:
		return p.(openContainer)
	case 39:
		return p.(closeContainer)
	case 40:
		return p.(containerSwap)
	case 41:
		return p.(itemApplyInContainer)
	case 42:
		return p.(containerStartCrafting)
	case 43:
		return p.(containerStopCrafting)
	case 44:
		return p.(burnContainer)
	case 45:
		return p.(clearContainer)
	case 46:
		return p.(worldClientStateUpdate)
	case 47:
		return p.(entityCreate)
	case 48:
		return p.(entityUpdate)
	case 49:
		return p.(entityDestroy)
	case 50:
		return p.(hitRequest)
	case 51:
		return p.(damageRequest)
	case 52:
		return p.(damageNotification)
	case 53:
		return p.(entityMessage)
	case 54:
		return p.(entityMessageResponse)
	case 55:
		return p.(updateWorldProperties)
	case 56:
		return p.(heartbeatUpdate)
	}
	return nil
}
*/
