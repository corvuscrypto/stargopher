package stargopher

import "time"

type funcType int

const APPEND_BEFORE = funcType(0)
const APPEND_ON = funcType(1)
const APPEND_AFTER = funcType(2)

//These are the maps that will contain the functions that modify
//the packets and performs actions before or after handling.
//Only the packetModHandlers should modify server/client sent data
var beforeHandlers = make(map[PacketType][]func(string))
var packetModHandlers = make(map[PacketType][]func(string, Packet) (Packet, bool))
var afterHandlers = make(map[PacketType][]func(string))

//AddPluginHandlers sets an array of functions to a map key
//to the appropriate handler map according to the behavior wanted.
//immediately invoked functions can also be set here.
func AddPluginHandlers(c funcType, p PacketType, f interface{}) {
	/*
	   E.g. this would mean add actions before each ChatSent packet type is
	   handled:

	  beforeHandlers[ChatSent] = append(beforeHandlers[ChatSent], myCustomFunction)
	*/
	switch c {
	case APPEND_BEFORE:
		beforeHandlers[p] = append(beforeHandlers[p], f.(func(string)))
		break
	case APPEND_ON:
		packetModHandlers[p] = append(packetModHandlers[p], f.(func(string, Packet) (Packet, bool)))
		break
	case APPEND_AFTER:
		afterHandlers[p] = append(afterHandlers[p], f.(func(string)))
		break
	}
}

//AddScheduledRoutine allows you to add scheduled tasks to the server, such as
//message of the day, or other periodic notifications. You can also schedule
//logging or routine restarts with the server.
func AddScheduledRoutine(d time.Duration, f func()) {

	//initiate goroutine to handle scheduled function
	go func(d time.Duration, f func()) {
		for {
			<-time.After(d)
			f()
		}
	}(d, f)

}

func addDefaultHandlers() {
	packetModHandlers[ClientConnect] = append(packetModHandlers[ClientConnect], addDefaultClientAttributes)
	packetModHandlers[ChatSent] = append(packetModHandlers[ChatSent], checkCommands)
	AddChatCommand("whisper", whisperCommand)
}
