package stargopher

type funcType int

const APPEND_BEFORE = funcType(0)
const APPEND_ON = funcType(1)
const APPEND_AFTER = funcType(2)

//AddPluginHandlers sets an array of functions to a map key
//to the appropriate handler map according to the behavior wanted.
//immediately invoked functions can also be set here.
func AddPluginHandlers(c funcType, p packetType, f interface{}) {
	/*
	   E.g. this would mean add actions before each ChatSent packet type is
	   handled:

	  beforeHandlers[ChatSent] = append(beforeHandlers[ChatSent], myCustomFunction)
	*/
	switch c {
	case APPEND_BEFORE:
		beforeHandlers[p] = append(beforeHandlers[p], f.(func()))
		break
	case APPEND_ON:
		packetModHandlers[p] = append(packetModHandlers[p], f.(func(Packet) Packet))
		break
	case APPEND_AFTER:
		afterHandlers[p] = append(afterHandlers[p], f.(func()))
		break
	}
}
