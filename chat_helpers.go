package stargopher

//BroadcastMessage takes a message and broadcasts it to everyone in the server
func BroadcastMessage(m string) {
	// packet := &chatReceived{
	// 	basePacket{ID: ChatReceived},
	// 	[]byte{1, 0, 0, 0, 0},
	// 	1,
	// 	6,
	// 	"SYSTEM",
	// 	SVLQ(len(m)),
	// 	m,
	// }
	//
	// data := SerializePacket(packet, 0)
	// for _, c := range connectedClients {
	// 	c.Connection.Outgoing <- data
	// }
}
