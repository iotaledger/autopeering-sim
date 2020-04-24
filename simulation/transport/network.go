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
	sync.Mutex
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

	c         chan transfer
	closeOnce sync.Once
	closing   chan struct{}
}

type transfer struct {
	pkt  []byte
	addr *net.UDPAddr
}

var (
	queueSize = 1
	ipNet     = net.IPNet{
		IP:   net.IPv4(10, 0, 0, 0),
		Mask: net.IPv4Mask(0, 0, 0xff, 0xff),
	}
)

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

func newConn(addr *net.UDPAddr, network *Network) *conn {
	c := &conn{
		addr:    addr,
		network: network,
		c:       make(chan transfer, queueSize),
		closing: make(chan struct{}),
	}
	// create empty buffer
	c.buf.Buffer = bytes.NewBuffer(nil)
	return c
}

// ReadFrom implements the Transport ReadFrom method.
func (c *conn) ReadFromUDP(b []byte) (int, *net.UDPAddr, error) {
	for {
		if c.buf.Len() > 0 {
			n, err := c.buf.Read(b)
			return n, c.buf.addr, err
		}

		if err := c.fillBuffer(); err != nil {
			return 0, nil, err
		}
	}
}

// WriteTo implements the Transport WriteTo method.
func (c *conn) WriteToUDP(pkt []byte, addr *net.UDPAddr) (int, error) {
	// determine the receiving peer
	peer, ok := c.network.connections[addr.String()]
	if !ok {
		return 0, ErrInvalidTarget
	}

	// nothing to write
	if len(pkt) == 0 {
		return 0, nil
	}
	// clone the packet before sending, just to make sure...
	req := transfer{pkt: append([]byte{}, pkt...), addr: c.addr}

	select {
	case peer.c <- req:
		return len(pkt), nil
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
	case res := <-c.c:
		c.buf.Buffer = bytes.NewBuffer(res.pkt)
		c.buf.addr = res.addr
		return nil
	case <-c.closing:
		return ErrClosed
	}
}

func ipFromID(id uint16) net.IP {
	nIP := ipNet.IP.To4()
	return net.IPv4(nIP[0], nIP[1], byte(id>>8), byte(id))
}
