package stargopher

import (
	"fmt"
	"log"
)

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
	//log for debugging for now
	//fmt.Println(PacketType(data[0]).String(), data[1:])

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
	//isolate the payload
	payload := data[len(data)-int(payloadLength):]
	pt := PacketType(data[0])
	//define special cases for building each packet type
	switch pt {

	case ChatSent:
		log.Println(payloadLength, data)
		text := string(payload[1 : len(payload)-1])
		if text[0] == '/' {
			//handleCommand()
			//ignore all commands for now to test behavior
			return nil
		}
		break

	case ChatReceived:
		log.Println(payloadLength, data)
		channel := data[:5]
		var i = 5
		userID := int(payload[i])
		i++
		userNameLength := int(payload[i])
		i++
		userName := string(payload[i : i+userNameLength])
		i += userNameLength
		messageLength := int(payload[i])
		i++
		message := string(payload[i : i+messageLength])
		fmt.Println("Client sent message: ", message)
		packet = &chatReceived{pt, channel, userID, userNameLength, userName, messageLength, message}
	}
	return packet
}
