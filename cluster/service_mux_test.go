package cluster

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/rqlite/rqlite/tcp"
	"github.com/rqlite/rqlite/testdata/x509"
)

func Test_NewServiceSetGetNodeAPIAddrMuxed(t *testing.T) {
	ln, mux := mustNewMux()
	go mux.Serve()
	tn := mux.Listen(1) // Could be any byte value.

	s := New(tn)
	if s == nil {
		t.Fatalf("failed to create cluster service")
	}

	if err := s.Open(); err != nil {
		t.Fatalf("failed to open cluster service")
	}

	s.SetAPIAddr("foo")

	addr, err := s.GetNodeAPIAddr(s.Addr())
	if err != nil {
		t.Fatalf("failed to get node API address: %s", err)
	}
	if addr != "http://foo" {
		t.Fatalf("failed to get correct node API address")
	}

	if err := ln.Close(); err != nil {
		t.Fatalf("failed to close Mux's listener: %s", err)
	}
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close cluster service")
	}
}

func Test_NewServiceSetGetNodeAPIAddrMuxedTLS(t *testing.T) {
	ln, mux := mustNewTLSMux()
	go mux.Serve()
	tn := mux.Listen(1) // Could be any byte value.

	s := New(tn)
	if s == nil {
		t.Fatalf("failed to create cluster service")
	}

	if err := s.Open(); err != nil {
		t.Fatalf("failed to open cluster service")
	}

	s.SetAPIAddr("foo")

	addr, err := s.GetNodeAPIAddr(s.Addr())
	if err != nil {
		t.Fatalf("failed to get node API address: %s", err)
	}
	if addr != "http://foo" {
		t.Fatalf("failed to get correct node API address")
	}

	if err := ln.Close(); err != nil {
		t.Fatalf("failed to close Mux's listener: %s", err)
	}
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close cluster service")
	}
}

func mustNewMux() (net.Listener, *tcp.Mux) {
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic("failed to create mock listener")
	}

	mux, err := tcp.NewMux(ln, nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create mux: %s", err))
	}

	return ln, mux
}

func mustNewTLSMux() (net.Listener, *tcp.Mux) {
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic("failed to create mock listener")
	}

	cert := x509.CertFile("")
	defer os.Remove(cert)
	key := x509.KeyFile("")
	defer os.Remove(key)

	mux, err := tcp.NewTLSMux(ln, nil, cert, key, "")
	if err != nil {
		panic(fmt.Sprintf("failed to create TLS mux: %s", err))
	}
	mux.InsecureSkipVerify = true

	return ln, mux
}
