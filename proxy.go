package stargopher

import (
	"io"
	"log"
	"net"
)

//This will be moved to a more organized config struct later
var starboundAddr, _ = net.ResolveTCPAddr("tcp", "localhost:21025")

//This will be changed later on
var connectedClients map[string]*Client

//Connection struct creates adds channels onto a TCP conn for intercept
//of data (See Pipe)
type Connection struct {
	*net.TCPConn
	Incoming chan []byte
	Outgoing chan []byte
}

func (c *Connection) handler(axeman chan error) {
	for {
		data := make([]byte, 512)

		//Always read stuff from the TCP connection
		n, err := c.Read(data)
		if err != nil {
			axeman <- err
		}
		//if there is data, then pass it off to the Incoming channel
		c.Incoming <- data[:n]
		//repeat
	}
}

//Client holds info and connection of the client
type Client struct {
	Connection
	//UID        string
	Attributes map[string]interface{}
}

//Server struct holds a connection to the actual starbound server program
type Server struct {
	Connection
}

//Pipe holds both a Client and a Server object for moderating transmission
//of data between them.
type Pipe struct {
	client *Client
	server *Server
	axeman chan error
}

func (pipe *Pipe) pipeRoutine() {
	go pipe.client.handler(pipe.axeman)
	go pipe.server.handler(pipe.axeman)
	for {
		//handle data immediately as it comes in.
		//Later there will be functions for analyzing and modifying packets
		select {
		case data := <-pipe.client.Incoming:
			pipe.server.Write(data)
			break
		case data := <-pipe.server.Incoming:
			pipe.client.Write(data)
			break
		case data := <-pipe.axeman:
			if data == io.EOF {
				pipe.client.Close()
				pipe.server.Close()
				log.Println("closed connection to", pipe.client.RemoteAddr().String())
				return
			}
			break
		}
	}
}

func newConnection(conn *net.TCPConn) *Pipe {

	toServer := make(chan []byte)
	toClient := make(chan []byte)

	nClient := &Client{
		Connection{conn, toServer, toClient},
		make(map[string]interface{}),
	}

	serverConn, _ := net.DialTCP("tcp", nil, starboundAddr)

	nServer := &Server{
		Connection{serverConn, toClient, toServer},
	}

	nPipe := &Pipe{
		nClient,
		nServer,
		make(chan error),
	}

	go nPipe.pipeRoutine()

	return nPipe
}
