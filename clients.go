package stargopher

//This will be changed later on
var connectedClients = make(map[string]*Client)

//Client holds info and connection of the client
type Client struct {
	Connection
	UID        string
	Attributes map[string]interface{} //holds all
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
	var clients map[string]*Client
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
