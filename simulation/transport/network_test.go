package transport

import (
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	peerA PeerID = iota
	peerB
	peerC
	peerD
)

var testPacket = []byte("TEST")

func TestNetwork_Listen(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	defer a.Close()

	if assert.IsType(t, new(net.UDPAddr), a.LocalAddr()) {
		assert.EqualValues(t, 0, a.LocalAddr().(*net.UDPAddr).Port)
	}
}

func TestNetwork_ListenTwice(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	defer a.Close()

	_, err = network.Listen(peerA, 0)
	assert.Truef(t, errors.Is(err, ErrOpen), "unexpected error: %v", err)
}

func TestConn_Close(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	assert.NoError(t, a.Close())
	assert.NoError(t, a.Close())
}

func TestConn_ReadClosed(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	require.NoError(t, a.Close())

	_, _, err = a.ReadFromUDP(nil)
	assert.Truef(t, errors.Is(err, ErrClosed), "unexpected error: %v", err)
}

func TestConn_ReadWrite(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	defer a.Close()
	b, err := network.Listen(peerB, 0)
	require.NoError(t, err)
	defer b.Close()

	n, err := a.WriteToUDP(testPacket, b.LocalAddr().(*net.UDPAddr))
	require.NoError(t, err)
	require.Equal(t, len(testPacket), n)

	buf := make([]byte, len(testPacket)+1)
	n, addr, err := b.ReadFromUDP(buf)
	require.NoError(t, err)
	require.Equal(t, len(testPacket), n)

	assert.Equal(t, buf[:n], testPacket)
	assert.Equal(t, addr, a.LocalAddr())
}

func TestConn_ReadByte(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	defer a.Close()
	b, err := network.Listen(peerB, 0)
	require.NoError(t, err)
	defer b.Close()

	n, err := a.WriteToUDP(testPacket, b.LocalAddr().(*net.UDPAddr))
	require.NoError(t, err)
	require.Equal(t, len(testPacket), n)

	buf := make([]byte, 1)
	for i := 0; i < len(testPacket); i++ {
		n, _, err := b.ReadFromUDP(buf)
		require.NoError(t, err)
		require.Equal(t, 1, n)
		assert.Equal(t, testPacket[i], buf[0])
	}
}

func TestConn_WriteInvalid(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	defer a.Close()

	addr := &net.UDPAddr{
		IP:   a.LocalAddr().(*net.UDPAddr).IP,
		Port: 1,
	}
	n, err := a.WriteToUDP(testPacket, addr)
	assert.Equal(t, 0, n)
	assert.Truef(t, errors.Is(err, ErrInvalidTarget), "unexpected error: %v", err)
}

func TestConn_WriteClosed(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	defer a.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < queueSize; i++ {
			_, err := a.WriteToUDP(testPacket, a.LocalAddr().(*net.UDPAddr))
			assert.NoError(t, err)
		}
		_, err := a.WriteToUDP(testPacket, a.LocalAddr().(*net.UDPAddr))
		assert.Truef(t, errors.Is(err, ErrClosed), "unexpected error: %v", err)
	}()

	time.Sleep(100 * time.Millisecond)
	require.NoError(t, a.Close())
	wg.Wait()
}

func TestConn_Port(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	defer a.Close()
	b0, err := network.Listen(peerB, 0)
	require.NoError(t, err)
	defer b0.Close()
	b1, err := network.Listen(peerB, 1)
	require.NoError(t, err)
	defer b1.Close()

	_, err = a.WriteToUDP([]byte("TEST0"), b0.LocalAddr().(*net.UDPAddr))
	require.NoError(t, err)
	_, err = a.WriteToUDP([]byte("TEST1"), b1.LocalAddr().(*net.UDPAddr))
	require.NoError(t, err)

	buf := make([]byte, 255)

	n, _, err := b0.ReadFromUDP(buf)
	require.NoError(t, err)
	assert.Equal(t, buf[:n], []byte("TEST0"))

	n, _, err = b1.ReadFromUDP(buf)
	require.NoError(t, err)
	assert.Equal(t, buf[:n], []byte("TEST1"))
}

func TestConn_WriteConcurrent(t *testing.T) {
	network := NewNetwork()

	a, err := network.Listen(peerA, 0)
	require.NoError(t, err)
	defer a.Close()
	b, err := network.Listen(peerB, 0)
	require.NoError(t, err)
	defer b.Close()
	c, err := network.Listen(peerC, 0)
	require.NoError(t, err)
	defer c.Close()
	d, err := network.Listen(peerD, 0)
	require.NoError(t, err)
	defer d.Close()

	var wg sync.WaitGroup
	const (
		numSender  = 3
		numPackets = 1000
	)

	// reader
	wg.Add(1)
	go func() {
		defer wg.Done()

		buf := make([]byte, len(testPacket)+1)
		for i := 0; i < numSender*numPackets; i++ {
			_, _, err := d.ReadFromUDP(buf)
			assert.NoError(t, err)
		}
	}()

	burst := func(c *conn, addr *net.UDPAddr) {
		defer wg.Done()
		for i := 0; i < numPackets; i++ {
			n, err := c.WriteToUDP(testPacket, addr)
			require.NoError(t, err)
			require.Equal(t, len(testPacket), n)
		}
	}

	wg.Add(numSender)
	burst(a, d.LocalAddr().(*net.UDPAddr))
	burst(b, d.LocalAddr().(*net.UDPAddr))
	burst(c, d.LocalAddr().(*net.UDPAddr))

	// wait for everything to finish
	wg.Wait()
}
