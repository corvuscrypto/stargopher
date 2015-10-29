package stargopher

import "fmt"

//Packet ... this may not be needed...
type Packet struct {
	ID            packetType
	PayloadLength int64
}

//PacketHandler will build an appropriate packet, then
//call any associated methods. The variable passthrough will be true if
//the data is meant to be sent through. Otherwise the packet will be dropped
func PacketHandler(data []byte, payloadLength int64) ([]byte, bool) {

	return data, true
}

//PacketDecoder is responsible for turning packet data into an easily
//modifiable struct to be later reencoded for transport
func PacketDecoder(data []byte, payloadLength int64) interface{} {

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
