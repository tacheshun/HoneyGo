package ssh

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/tacheshun/honeygo/pkg/config"
	"github.com/tacheshun/honeygo/pkg/log"
	"golang.org/x/crypto/ssh"
)

// Server represents the SSH honeypot server
type Server struct {
	config *config.Config
	sshConfig *ssh.ServerConfig
	listener net.Listener
	logger *log.Logger
}

// NewServer creates a new SSH server
func NewServer(cfg *config.Config) (*Server, error) {
	logger, err := log.NewLogger(cfg.LogPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	server := &Server{
		config: cfg,
		logger: logger,
	}

	if err := server.configureSSH(); err != nil {
		return nil, err
	}

	return server, nil
}

// configureSSH sets up the SSH server configuration
func (s *Server) configureSSH() error {
	s.sshConfig = &ssh.ServerConfig{
		ServerVersion: s.config.Banner,
		
		// Always reject authentication, but log the attempt
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			s.logAuthAttempt(conn, "password", string(password))
			return nil, fmt.Errorf("password rejected")
		},
		
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			s.logAuthAttempt(conn, "publickey", string(key.Marshal()))
			return nil, fmt.Errorf("public key rejected")
		},
	}

	// Load the host key
	if s.config.HostKey != "" {
		privateBytes, err := ioutil.ReadFile(s.config.HostKey)
		if err != nil {
			return fmt.Errorf("failed to load host key: %w", err)
		}

		private, err := ssh.ParsePrivateKey(privateBytes)
		if err != nil {
			return fmt.Errorf("failed to parse host key: %w", err)
		}

		s.sshConfig.AddHostKey(private)
	} else {
		// Generate a temporary key
		private, err := generateTemporaryKey(s.config.HostKeyType)
		if err != nil {
			return fmt.Errorf("failed to generate temporary key: %w", err)
		}
		
		s.sshConfig.AddHostKey(private)
	}

	return nil
}

// Start starts the SSH server
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.config.ListenAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.config.ListenAddress, err)
	}
	s.listener = listener

	s.logger.Info("Server started on %s", s.config.ListenAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				continue
			}
			return err
		}

		go s.handleConnection(conn)
	}
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown() {
	if s.listener != nil {
		s.listener.Close()
	}
	s.logger.Info("Server shutdown complete")
	s.logger.Close()
}

// handleConnection handles an incoming connection
func (s *Server) handleConnection(nConn net.Conn) {
	remoteAddr := nConn.RemoteAddr().String()
	s.logger.Info("New connection from %s", remoteAddr)

	_, chans, reqs, err := ssh.NewServerConn(nConn, s.sshConfig)
	if err != nil {
		s.logger.Info("Failed handshake: %v", err)
		nConn.Close()
		return
	}

	// Discard all global requests
	go ssh.DiscardRequests(reqs)

	// Handle channels
	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		channel, requests, err := newChannel.Accept()
		if err != nil {
			s.logger.Error("Could not accept channel: %v", err)
			continue
		}

		// Handle channel requests
		go func(in <-chan *ssh.Request) {
			for req := range in {
				ok := false
				switch req.Type {
				case "shell":
					ok = true
					// Could provide a fake shell here
				case "exec":
					// Could handle exec commands here
				case "pty-req":
					ok = true
				case "env":
					ok = true
				}
				req.Reply(ok, nil)
			}
		}(requests)

		// Close the session after a short delay
		go func() {
			time.Sleep(2 * time.Second)
			channel.Close()
		}()
	}
}

// logAuthAttempt logs authentication attempts
func (s *Server) logAuthAttempt(conn ssh.ConnMetadata, method, credential string) {
	s.logger.Auth(conn.RemoteAddr().String(), conn.User(), method, credential)
}