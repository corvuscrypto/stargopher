package stargopher

import "fmt"

//BroadcastMessage takes a message and broadcasts it to everyone in the server
func BroadcastMessage(m string) {
	packet := &ChatReceivedPacket{
		0,
		6,
		"SYSTEM",
		len(m),
		m,
	}
	fmt.Println(packet)
}
