package network

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/369guang/tg-im/internal/logs"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/webtransport-go"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type WebTransportServer struct {
	Server   *webtransport.Server
	CertFile string
	KeyFile  string
}

type WebTransportClient struct {
	Session *webtransport.Session
	Stream  webtransport.Stream
}

func NewWebTransportServer(addr, certFile, keyFile string) (*WebTransportServer, error) {

	tlsCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		logs.Logger.Error("Error loading cert", zap.Error(err))
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{http3.NextProtoH3},
	}

	// enable webtransport
	server := webtransport.Server{
		H3: http3.Server{
			Addr:            addr,
			TLSConfig:       tlsConfig,
			EnableDatagrams: true,
			QUICConfig: &quic.Config{
				EnableDatagrams: true,
			},
		},
	}

	return &WebTransportServer{Server: &server, CertFile: certFile, KeyFile: keyFile}, nil
}

func (s *WebTransportServer) Serve() error {
	fmt.Println("WebTransport server started")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := s.Server.Upgrade(writer, request)
		if err != nil {
			fmt.Println("Error upgrading connection: ", err)
			logs.Logger.Error("Error upgrading connection", zap.Error(err))
			return
		}
		fmt.Println("new connection: ", conn)
		go s.handleSession(conn)
	})

	logs.Logger.Info("WebTransport server started")
	//err := s.Server.H3.ListenAndServeTLS(s.CertFile, s.KeyFile)
	err := s.Server.ListenAndServeTLS(s.CertFile, s.KeyFile)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		logs.Logger.Error("Error starting server", zap.Error(err))
		return err
	} else {
		fmt.Println("WebTransport server ListenAndServeTLS")
	}
	return nil
}

func (s *WebTransportServer) handleSession(conn *webtransport.Session) {
	for {
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			fmt.Println("Error accepting stream: ", err)
			logs.Logger.Error("Error accepting stream", zap.Error(err))
			return
		}

		go s.handleStream(stream)
	}
}

func (s *WebTransportServer) handleStream(stream webtransport.Stream) {

	defer stream.Close()

	buf := make([]byte, 1024)
	for {
		n, err := stream.Read(buf)
		if err != nil {
			fmt.Println("Error reading from stream: ", err)
			logs.Logger.Error("Error reading from stream", zap.Error(err))
			return
		}

		logs.Logger.Info("Received message", zap.String("message", string(buf[:n])))

		_, err = stream.Write(buf[:n])
		if err != nil {
			fmt.Println("Error writing to stream: ", err)
			logs.Logger.Error("Error writing to stream", zap.Error(err))
			return
		}
	}
}

func NewWebTransportClient(addr string) (*WebTransportClient, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{http3.NextProtoH3},
	}
	dialer := &webtransport.Dialer{
		TLSClientConfig:         tlsConfig,
		StreamReorderingTimeout: 5 * time.Second,
	}
	var headers http.Header
	fmt.Println("addr: ", addr)
	rsp, session, err := dialer.Dial(context.Background(), addr, headers)
	if err != nil {
		fmt.Println("Error dialing: ", err)
		logs.Logger.Error("Error dialing", zap.Error(err))
		return nil, err
	}
	defer rsp.Body.Close()

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		fmt.Println("Error opening stream: ", err)
		logs.Logger.Error("Error opening stream", zap.Error(err))
		return nil, err
	}

	return &WebTransportClient{Session: session, Stream: stream}, nil
}

func (c *WebTransportClient) Send(message string) error {
	if c.Stream == nil {
		return fmt.Errorf("stream is nil")
	}

	fmt.Println("Sending message", message)
	_, err := c.Stream.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to stream: ", err)
		return err
	}

	return nil
}

func (c *WebTransportClient) Receive() (string, error) {
	if c.Stream == nil {
		return "", fmt.Errorf("stream is nil")
	}
	buf := make([]byte, 1024)
	n, err := c.Stream.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func (c *WebTransportClient) Close() error {
	if c.Session == nil {
		return fmt.Errorf("session is nil")
	}
	return c.Session.CloseWithError(0, "client shutdown")
}
