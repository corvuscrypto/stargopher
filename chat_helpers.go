package stargopher

import "log"

//BroadcastMessage takes a message and broadcasts it to everyone in the server
func BroadcastMessage(m string) {
	packet := &chatReceived{
		ChatReceived,
		[]byte{1, 0, 0, 0, 0},
		1,
		6,
		"SYSTEM",
		len(m),
		m,
	}

	data := SerializePacket(packet, 0)
	log.Println("sending", data)
	for _, c := range connectedClients {
		log.Println(c.UID)
		c.Connection.Outgoing <- data
	}
}
