// Package httpconn establish http server.
package httpconn

import (
	"fmt"
	"net"
	"net/http"
)

// Server creates a new http server.
type server struct {
	*http.Server
	port int
}

// Serve interface.
type Serve interface {
	Close()
	Start() error
	StartServer() (net.Listener, error)
}

// New http server which runs in give port.
func New(port int) Serve {

	return &server{
		port: port,
	}
}

func (s *server) Start() error {

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil); err != nil {
		return fmt.Errorf("Cannot establish http server connection:%v", err)
	}

	fmt.Printf("Http server listening on :%d\n", s.port)

	return nil
}

// StartServer start new net server.
func (s *server) StartServer() (net.Listener, error) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return nil, fmt.Errorf("Cannot establish http server connection:%v", err)
	}

	fmt.Printf("Server listening on port:%d", s.port)

	return listener, nil
}

// Close the server.
func (s *server) Close() {
	s.Close()
}
