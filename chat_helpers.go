package stargopher

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

var commandHandlers = make(map[string]func(string, []string) bool)

//BroadcastMessage takes a message and broadcasts it to everyone in the server
func BroadcastMessage(m string) {
	packet := &chatReceived{
		basePacket{ID: ChatReceived},
		[]byte{1, 0, 0, 0, 0},
		1,
		14,
		"^FF0000;SYSTEM",
		VLQ(len(m)),
		m,
	}
	data := SerializePacket(packet, 0)
	for _, c := range connectedClients {
		c.send(data)
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
		15,
		"^#ff0000;SYSTEM",
		VLQ(len(msg)),
		msg,
	}

	client.sendPacket(packet)

}

func checkCommands(uid string, packet Packet) (Packet, bool) {
	castPacket := packet.(reflect.Value).Interface().(chatSent)
	message := castPacket.Message
	if message[0] == '/' {
		var command, param, param2 string
		spaceIndex := strings.Index(message, " ")
		pipeIndex := strings.Index(message, "|")
		if spaceIndex > 0 && spaceIndex < len(message)-1 {
			command = message[1:spaceIndex]
			param = message[spaceIndex+1:]
		} else {
			command = message[1:]
		}
		if param != "" && pipeIndex > 0 && pipeIndex < len(message)-1 {
			param, param2 = message[spaceIndex+1:pipeIndex], message[pipeIndex+1:]
		}

		params := []string{param, param2}

		fmt.Println(command, params)
		if f, ok := commandHandlers[command]; ok {
			return packet, f(uid, params)
		}
	}
	return packet, false
}

func whisperCommand(uid string, params []string) bool {
	if len(params) != 2 {
		return false
	}
	me := "^#BF5FFF;" + GetClient(uid).Attributes["username"].(string)
	m := &chatReceived{basePacket: basePacket{}}
	m.ID = PacketType(ChatReceived)
	m.UserName = me
	m.UserNameLength = VLQ(len(me))
	m.Message = "^#AEEEEE;" + params[1]
	m.MessageLength = VLQ(len(m.Message))
	m.UserID = 123
	m.Channel = []byte{1, 0, 0, 0, 0}
	cmap := GetClientsByAttributes(map[string]interface{}{"username": params[0]})
	cmap[uid] = GetClient(uid)
	for _, v := range cmap {
		v.sendPacket(m)
	}

	return false
}

func AddChatCommand(cmd string, f func(string, []string) bool) {
	commandHandlers[cmd] = f
}
