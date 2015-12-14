package stargopher

import (
	"log"
	"net"

	"github.com/corvuscrypto/stargopher"
)

//StarboundServer is the main server object that contains information for
//initializing a TCP server for starbound
type StarboundServer struct {
	Name    string
	Address string
	Port    string
}

//Init is the main method in StarboundServer that initializes the server and
//starts the connection listener associated with the server
func (s StarboundServer) Init() {
	listenAddr, err := net.ResolveTCPAddr("tcp", s.Address+":"+s.Port)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	go s.listen(listener)
}

func (s StarboundServer) listen(l *net.TCPListener) {
	for {
		conn, _ := l.AcceptTCP()
		testPipe = stargopher.NewConnection(conn)
	}
}
