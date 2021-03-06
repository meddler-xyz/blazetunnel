package server

import (
	"log"
	"net"

	"github.com/lucas-clemente/quic-go"
)

// Server struct contains the runtime configuration
// for the server to function
type Server struct {
	hostmap        *HostMap
	domain         string
	idleTimeout    uint
	publicListener net.Listener
	tunnelListener quic.Listener
}

// NewServer is used to create a new Server instance
// and initialize all the fields
func NewServer(domain string, idleTimeout uint) *Server {
	log.Println("Starting blazetunnel server")
	log.Println("Domain", domain, idleTimeout)
	log.Println("Idle Timeout", domain, idleTimeout)

	return &Server{
		domain:      domain,
		idleTimeout: idleTimeout,
		hostmap:     NewHostMap(),
	}
}

func (s *Server) init() error {
	err := s.initPublic()
	if err != nil {
		return err
	}

	return s.initTunnel()
}

// Start starts the server by listening to the public instance and the
// tunneling instance0.0.
func (s *Server) Start() error {
	err := s.init()
	if err != nil {
		return err
	}

	go s.startPublic()
	s.startTunnel()
	return nil
}
