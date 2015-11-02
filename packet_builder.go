package stargopher

import "fmt"

//Packet ... this may not be needed...
type Packet interface{}

//These are the maps that will contain the functions that modify
//the packets and performs actions before or after handling.
//Only the packetModHandlers should modify server/client sent data
var beforeHandlers = make(map[packetType][]func())
var packetModHandlers = make(map[packetType][]func(Packet) Packet)
var afterHandlers = make(map[packetType][]func())

//PacketHandler will build an appropriate packet, then
//call any associated methods. The variable passthrough will be true if
//the data is meant to be sent through. Otherwise the packet will be dropped
func PacketHandler(data []byte, payloadLength int64) ([]byte, bool) {
	//log for debugging for now
	fmt.Println(packetType(data[0]).String(), data[1:])

	ptype := packetType(data[0])
	packet := PacketDecoder(data, payloadLength)

	//first handle the before action if exists
	for _, f := range beforeHandlers[ptype] {
		f()
	}

	//then do the packet modifying functions
	for _, f := range packetModHandlers[ptype] {
		packet = f(packet)
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

	//isolate the payload
	payload := data[len(data)-int(payloadLength):]

	//define special cases for building each packet type
	switch packetType(data[0]) {

	case chatSent:
		fmt.Println(payload)
		text := string(payload[1 : len(payload)-1])
		if text[0] == '/' {
			//handleCommand()
			//ignore all commands for now to test behavior
			return nil
		}
		fmt.Println("Client sent message: ", text)
		break

	case chatReceived:
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
		return &ChatReceived{userID, userNameLength, userName, messageLength, message}
	}
	return nil
}
