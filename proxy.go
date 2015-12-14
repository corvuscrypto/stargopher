package stargopher

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"net"
)

//This will be moved to a more organized config struct later
var starboundAddr, _ = net.ResolveTCPAddr("tcp", "localhost:21025")

//Connection struct creates adds channels onto a TCP conn for intercept
//of data (See Pipe)
type Connection struct {
	*net.TCPConn
	Incoming chan []byte
	Outgoing chan []byte
}

func (c *Connection) handler(axeman chan error, uid string) {
	pc := make(chan []byte)
	for {
		var packet []byte
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
		payloadLength, _ := ReadSVarint(packet[1:])

		//register how many bytes remain to read
		var remaining = int(payloadLength)
		if remaining < 0 {
			remaining = -remaining
		}

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
		var packetSend = make([]byte, len(packet))
		copy(packetSend, packet)
		if payloadLength < 0 {
			b := bytes.NewReader(packetSend[iterator+1:])
			zr, _ := zlib.NewReader(b)
			g, _ := ioutil.ReadAll(zr)
			zr.Close()
			packetSend = append(packetSend[:iterator+1], g...)
			payloadLength = int64(len(g))
		}
		if payloadLength < 0 {
			payloadLength = -payloadLength
		}
		if packet[0] == 2 {
			fmt.Println(packetSend)
			fmt.Println(payloadLength)
		}
		go PacketHandler(uid, pc, packetSend, payloadLength)
		<-pc
		c.Incoming <- packet
	}
}

//Server struct holds a connection to the actual starbound server program
type Server struct {
	Connection
}

func (s Server) sendPacket(packet Packet) {
	data := SerializePacket(packet, 0)
	s.Connection.Outgoing <- data
}

func (s Server) send(data []byte) {
	s.Connection.Outgoing <- data
}

//Pipe holds both a Client and a Server object for moderating transmission
//of data between them.
type Pipe struct {
	client *Client
	server *Server
	axeman chan error
}

func (pipe *Pipe) pipeRoutine() {
	go pipe.client.handler(pipe.axeman, pipe.client.UID)
	go pipe.server.handler(pipe.axeman, pipe.client.UID)
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
				//remove client from map
				delete(connectedClients, pipe.client.UID)
				return
			}
			break
		}
	}
}

//NewConnection creates a new connection to the starbound server and returns a
//full duplex TCP connection represented as type Pipe
func NewConnection(conn *net.TCPConn) *Pipe {

	toServer := make(chan []byte)
	toClient := make(chan []byte)

	nClient := &Client{
		Connection{conn, toServer, toClient},
		newUUID(),
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
	connectedClients[nClient.UID] = nClient
	go nPipe.pipeRoutine()

	return nPipe
}
