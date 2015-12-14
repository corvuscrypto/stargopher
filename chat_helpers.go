package stargopher

import (
	"log"
)

//BroadcastMessage takes a message and broadcasts it to everyone in the server
func BroadcastMessage(m string) {
	packet := &chatReceived{
		basePacket{ID: ChatReceived},
		[]byte{1, 0, 0, 0, 0},
		1,
		6,
		"SYSTEM",
		VLQ(len(m)),
		m,
	}

	data := SerializePacket(packet, 0)
	for _, c := range connectedClients {
		c.Connection.Outgoing <- data
	}
}

//MessageClient provides a direct method for the system to send a message to
//the client.
func MessageClient(uid string, msg string) {
	client, ok := connectedClients[uid]
	if !ok {
		log.Printf("Tried to send message to client %s, but client was not found!\n", uid)
		return
	}

	packet := &chatReceived{
		basePacket{ID: ChatReceived},
		[]byte{1, 0, 0, 0, 0},
		1,
		6,
		"SYSTEM",
		VLQ(len(msg)),
		msg,
	}

	data := SerializePacket(packet, 0)

	client.Connection.Outgoing <- data

}
