package transport

import (
	"bytes"
	"errors"
	"net"
	"sync"
)

// Errors returned by the transport package
var (
	ErrClosed        = errors.New("use of closed network connection")
	ErrOpen          = errors.New("connection already open")
	ErrInvalidTarget = errors.New("invalid target address")
)

// Network offers in-memory transfers between an arbitrary number of clients.
type Network struct {
	sync.RWMutex
	connections map[string]*conn
}

// A PeerID specifies the ID of a peer to which we can connect.
type PeerID = uint16

type conn struct {
	network *Network
	addr    *net.UDPAddr

	buf struct {
		*bytes.Buffer
		addr *net.UDPAddr
	}

	queue     chan packet
	closeOnce sync.Once
	closing   chan struct{}
}

type packet struct {
	data []byte
	addr *net.UDPAddr
}

var queueSize = 1

// NewNetwork creates a new in-memory transport transport.
// For each provided address a corresponding client is created.
func NewNetwork() *Network {
	return &Network{
		connections: make(map[string]*conn),
	}
}

// Listen announces that peer and port on the local transport address.
func (n *Network) Listen(id PeerID, port int) (*conn, error) {
	n.Lock()
	defer n.Unlock()

	addr := &net.UDPAddr{IP: ipFromID(id), Port: port}
	if _, contains := n.connections[addr.String()]; contains {
		return nil, ErrOpen
	}

	c := newConn(addr, n)
	n.connections[addr.String()] = c
	return c, nil
}

func (n *Network) close(c *conn) {
	n.Lock()
	defer n.Unlock()

	delete(n.connections, c.addr.String())
}

func (n *Network) peer(addr string) *conn {
	n.RLock()
	defer n.RUnlock()
	return n.connections[addr]
}

func newConn(addr *net.UDPAddr, network *Network) *conn {
	c := &conn{
		addr:    addr,
		network: network,
		queue:   make(chan packet, queueSize),
		closing: make(chan struct{}),
	}
	// create empty buffer
	c.buf.Buffer = bytes.NewBuffer(nil)
	return c
}

// ReadFrom implements the Transport ReadFrom method.
func (c *conn) ReadFromUDP(b []byte) (int, *net.UDPAddr, error) {
	for {
		// read from buffer
		if c.buf.Len() > 0 {
			n, err := c.buf.Read(b)
			return n, c.buf.addr, err
		}
		// load from queue
		if err := c.fillBuffer(); err != nil {
			return 0, nil, err
		}
	}
}

// WriteTo implements the Transport WriteTo method.
func (c *conn) WriteToUDP(b []byte, addr *net.UDPAddr) (int, error) {
	// determine the receiving to
	to := c.network.peer(addr.String())
	if to == nil {
		return 0, ErrInvalidTarget
	}

	// nothing to write
	if len(b) == 0 {
		return 0, nil
	}
	// copy the data before sending, just to make sure...
	pkt := packet{data: append([]byte{}, b...), addr: c.addr}

	select {
	case to.queue <- pkt:
		return len(b), nil
	case <-c.closing:
		return 0, ErrClosed
	}
}

// Close closes the transport layer.
func (c *conn) Close() error {
	c.closeOnce.Do(func() {
		close(c.closing)
		c.network.close(c)
	})
	return nil
}

// LocalAddr returns the local transport address.
func (c *conn) LocalAddr() net.Addr {
	return c.addr
}

func (c *conn) fillBuffer() error {
	select {
	case pkt := <-c.queue:
		c.buf.Buffer = bytes.NewBuffer(pkt.data)
		c.buf.addr = pkt.addr
		return nil
	case <-c.closing:
		return ErrClosed
	}
}

func ipFromID(id uint16) net.IP {
	// use 10.0.0.0/16 subnet
	return net.IPv4(10, 0, byte(id>>8), byte(id))
}
