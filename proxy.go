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
		var packet []byte
		var payloadLength int64
		//first start pulling 1 byte at a time until VLQ is resolved
		var iterator int64
		for {
			data := make([]byte, 1)
			_, err := c.Read(data)
			if err != nil {
				axeman <- err
			}
			packet = append(packet, data[0:]...)

			if data[0] < 0x80 && iterator > 0 {
				break
			}
			iterator++
		}

		//register the size of the Payload
		payloadLength = Varint(packet[1:]) / 2

		//if Payload length is negative to indicate compression, make it positive and add 1
		if payloadLength < 0 {
			payloadLength = (-payloadLength) + 1
		}

		//register how many bytes remain to read
		var remaining = int(payloadLength)

		//loop and read the TCP stream until remaining = 0
		for {
			var data []byte

			//Here, max buffer size is 256, this should be programmatically
			//determined in the future to allow full control

			var maxBufferSize = 256

			if remaining < maxBufferSize {
				data = make([]byte, remaining)
			} else {
				data = make([]byte, maxBufferSize)
			}
			n, err := c.Read(data)
			if err != nil {
				axeman <- err
			}
			//if there is data, then add it to the new packet
			packet = append(packet, data[0:n]...)

			remaining -= n

			//break the loop if no more data from this packet
			if remaining == 0 {
				break
			}
		}

		//call packet handler
		var passthrough bool
		packet, passthrough = PacketHandler(packet, payloadLength)

		//send the packet across if passthrough is true
		if passthrough {
			c.Incoming <- packet
		}
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
