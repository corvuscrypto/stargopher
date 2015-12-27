package stargopher

import "reflect"

//This will be changed later on
var connectedClients = make(map[string]*Client)

//Client holds info and connection of the client
type Client struct {
	Connection
	UID        string
	Attributes map[string]interface{} //holds all
}

//convenience method
func (c Client) sendPacket(packet Packet) {
	data := SerializePacket(packet, 0)
	c.Connection.Outgoing <- data
}

func (c Client) send(data []byte) {
	c.Connection.Outgoing <- data
}

//GetClient takes a uid string argument and returns the reference to that client
//if present, or nil if the client doesn't exist.
func GetClient(uid string) *Client {
	return connectedClients[uid]
}

//GetClientsByAttributes takes a map of attributes and returns a list of clients
//that have attributes equal to those in the map
func GetClientsByAttributes(attrs map[string]interface{}) map[string]*Client {
	//copy clients in map to array
	var clients = make(map[string]*Client)
	for k, v := range connectedClients {
		clients[k] = v
	}
	for k, a := range attrs {
		for i, c := range clients {
			if c.Attributes[k] != a {
				delete(clients, i)
			}
		}
	}
	return clients
}

//GetReadyClients returns a list of clients
//that are fully connected and ready
func GetReadyClients() map[string]*Client {
	//copy clients in map to array
	var clients = make(map[string]*Client)
	for k, v := range connectedClients {
		clients[k] = v
	}

	for i, c := range clients {
		if c.Attributes["isReady"] != true {
			delete(clients, i)
		}
	}
	return clients
}

//default function for attaching attributes to a client
func addDefaultClientAttributes(uid string, packet Packet) (Packet, bool) {
	c := GetClient(uid)
	castPacket := packet.(reflect.Value).Interface().(clientConnect)
	c.Attributes["username"] = castPacket.Name
	c.Attributes["species"] = castPacket.Species
	c.Attributes["isReady"] = true
	return Packet(packet), false
}
