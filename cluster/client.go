package cluster

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/rqlite/rqlite/command"
	"github.com/rqlite/rqlite/tcp/pool"
)

const (
	initialPoolSize = 4
	maxPoolCapacity = 64
)

// Client allows communicating with a remote node.
type Client struct {
	dialer  Dialer
	timeout time.Duration

	mu    sync.RWMutex
	pools map[string]pool.Pool
}

// NewClient returns a client instance for talking to a remote node.
func NewClient(dl Dialer) *Client {
	return &Client{
		dialer:  dl,
		timeout: 30 * time.Second,
		pools:   make(map[string]pool.Pool),
	}
}

// GetNodeAPIAddr retrieves the API Address for the node at nodeAddr
func (c *Client) GetNodeAPIAddr(nodeAddr string) (string, error) {
	var pl pool.Pool
	var ok bool

	c.mu.RLock()
	pl, ok = c.pools[nodeAddr]
	c.mu.RUnlock()

	// Do we need a new pool for the given address?
	if !ok {
		if err := func() error {
			c.mu.Lock()
			defer c.mu.Unlock()
			pl, ok = c.pools[nodeAddr]
			if ok {
				return nil // Pool was inserted just after we checked.
			}

			// New pool is needed for given address.
			factory := func() (net.Conn, error) { return c.dialer.Dial(nodeAddr, c.timeout) }
			p, err := pool.NewChannelPool(initialPoolSize, maxPoolCapacity, factory)
			if err != nil {
				return err
			}
			c.pools[nodeAddr] = p
			pl = p
			return nil
		}(); err != nil {
			return "", err
		}
	}

	// Got pool, now get a connection.
	conn, err := pl.Get()
	if err != nil {
		return "", fmt.Errorf("pool get: %s", err)
	}
	defer conn.Close()

	handleConnError := func(c net.Conn) {
		if pc, ok := conn.(*pool.PoolConn); ok {
			pc.MarkUnusable()
		}
	}

	// Send the request
	command := &Command{
		Type: Command_COMMAND_TYPE_GET_NODE_API_URL,
	}
	p, err := proto.Marshal(command)
	if err != nil {
		return "", fmt.Errorf("command marshal: %s", err)
	}

	// Write length of Protobuf
	b := make([]byte, 4)
	binary.LittleEndian.PutUint16(b[0:], uint16(len(p)))

	_, err = conn.Write(b)
	if err != nil {
		handleConnError(conn)
		return "", fmt.Errorf("write protobuf length: %s", err)
	}
	_, err = conn.Write(p)
	if err != nil {
		handleConnError(conn)
		return "", fmt.Errorf("write protobuf: %s", err)
	}

	b, err = ioutil.ReadAll(conn)
	if err != nil {
		handleConnError(conn)
		return "", fmt.Errorf("read protobuf bytes: %s", err)
	}

	a := &Address{}
	err = proto.Unmarshal(b, a)
	if err != nil {
		return "", fmt.Errorf("protobuf unmarshal: %s", err)
	}

	return a.Url, nil
}

// Execute performs an Execute on a remote node.
func (c *Client) Execute(er *command.ExecuteRequest, nodeAddr string, timeout time.Duration) ([]*command.ExecuteResult, error) {
	conn, err := c.dialer.Dial(nodeAddr, c.timeout)
	if err != nil {
		return nil, fmt.Errorf("dial connection: %s", err)
	}
	defer conn.Close()

	// Create the request.
	command := &Command{
		Type: Command_COMMAND_TYPE_EXECUTE,
		Request: &Command_ExecuteRequest{
			ExecuteRequest: er,
		},
	}
	p, err := proto.Marshal(command)
	if err != nil {
		return nil, fmt.Errorf("command marshal: %s", err)
	}

	// Write length of Protobuf
	b := make([]byte, 4)
	binary.LittleEndian.PutUint16(b[0:], uint16(len(p)))

	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	_, err = conn.Write(b)
	if err != nil {
		return nil, err
	}
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	_, err = conn.Write(p)
	if err != nil {
		return nil, err
	}

	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	b, err = ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	a := &CommandExecuteResponse{}
	err = proto.Unmarshal(b, a)
	if err != nil {
		return nil, err
	}

	if a.Error != "" {
		return nil, errors.New(a.Error)
	}
	return a.Results, nil
}

// Query performs an Query on a remote node.
func (c *Client) Query(qr *command.QueryRequest, nodeAddr string, timeout time.Duration) ([]*command.QueryRows, error) {
	conn, err := c.dialer.Dial(nodeAddr, c.timeout)
	if err != nil {
		return nil, fmt.Errorf("dial connection: %s", err)
	}
	defer conn.Close()

	// Create the request.
	command := &Command{
		Type: Command_COMMAND_TYPE_QUERY,
		Request: &Command_QueryRequest{
			QueryRequest: qr,
		},
	}
	p, err := proto.Marshal(command)
	if err != nil {
		return nil, fmt.Errorf("command marshal: %s", err)
	}

	// Write length of Protobuf, the Protobuf
	b := make([]byte, 4)
	binary.LittleEndian.PutUint16(b[0:], uint16(len(p)))

	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	_, err = conn.Write(b)
	if err != nil {
		return nil, err
	}
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	_, err = conn.Write(p)
	if err != nil {
		return nil, err
	}

	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	b, err = ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	a := &CommandQueryResponse{}
	err = proto.Unmarshal(b, a)
	if err != nil {
		return nil, err
	}

	if a.Error != "" {
		return nil, errors.New(a.Error)
	}
	return a.Rows, nil
}

// Stats returns stats on the Client instance
func (c *Client) Stats() (map[string]interface{}, error) {
	return map[string]interface{}{
		"timeout": c.timeout,
	}, nil
}
