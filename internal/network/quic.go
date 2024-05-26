package network

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/369guang/tg-im/internal/logs"
	"github.com/quic-go/quic-go"
	"go.uber.org/zap"
	"io"
)

type QuicServer struct {
	listener *quic.Listener
}

type QuicClient struct {
	session quic.Connection
	stream  quic.Stream
}

func NewQuicServer(addr, certFile, keyFile string) (*QuicServer, error) {
	tlsCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-chat"},
	}

	listener, err := quic.ListenAddr(addr, tlsConfig, nil)
	if err != nil {
		return nil, err
	}

	return &QuicServer{listener: listener}, nil
}

func (s *QuicServer) Serve() error {
	for {
		session, err := s.listener.Accept(context.Background())
		if err != nil {
			logs.Logger.Error("Error accepting session", zap.Error(err))
			continue
		}

		go s.handleSession(session)
	}
}

func (s *QuicServer) handleSession(session quic.Connection) {
	for {
		stream, err := session.AcceptStream(context.Background())
		if err != nil {
			logs.Logger.Error("Error accepting stream", zap.Error(err))
			return
		}

		go s.handleStream(stream)
	}
}

func (s *QuicServer) handleStream(stream quic.Stream) {
	defer stream.Close()

	buf := make([]byte, 1024)
	for {
		n, err := stream.Read(buf)
		if err != nil {
			logs.Logger.Error("Error reading from stream", zap.Error(err))
			return
		}

		logs.Logger.Info("Received message", zap.String("message", string(buf[:n])))

		_, err = stream.Write(buf[:n])
		if err != nil {
			logs.Logger.Error("Error writing to stream", zap.Error(err))
			return
		}
	}
}

func NewQuicClient(addr, serverName string) (*QuicClient, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-chat"},
	}

	session, err := quic.DialAddr(context.Background(), addr, tlsConfig, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("session: ", session)

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Println("stream: ", stream)

	return &QuicClient{session: session, stream: stream}, nil
}

func (c *QuicClient) Send(msg string) error {
	_, err := c.stream.Write([]byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func (c *QuicClient) Receive() (string, error) {
	buf := make([]byte, 1024)
	n, err := c.stream.Read(buf)
	if err != nil {
		if err != io.EOF {
			return "", err
		}
	}

	return string(buf[:n]), nil
}

func (c *QuicClient) Close() error {
	err := c.stream.Close()
	if err != nil {
		return err
	}

	err = c.session.CloseWithError(0, "client shutdown")
	if err != nil {
		return err
	}

	return nil
}
