package tap

import (
	"net"

	"github.com/go-gost/gost/v3/pkg/listener"
	"github.com/go-gost/gost/v3/pkg/logger"
	md "github.com/go-gost/gost/v3/pkg/metadata"
	"github.com/go-gost/gost/v3/pkg/registry"
	tap_util "github.com/go-gost/x/internal/util/tap"
)

func init() {
	registry.ListenerRegistry().Register("tap", NewListener)
}

type tapListener struct {
	saddr  string
	addr   net.Addr
	cqueue chan net.Conn
	closed chan struct{}
	logger logger.Logger
	md     metadata
}

func NewListener(opts ...listener.Option) listener.Listener {
	options := &listener.Options{}
	for _, opt := range opts {
		opt(options)
	}
	return &tapListener{
		saddr:  options.Addr,
		logger: options.Logger,
	}
}

func (l *tapListener) Init(md md.Metadata) (err error) {
	if err = l.parseMetadata(md); err != nil {
		return
	}

	l.addr, err = net.ResolveUDPAddr("udp", l.saddr)
	if err != nil {
		return
	}

	ifce, ip, err := l.createTap()
	if err != nil {
		if ifce != nil {
			ifce.Close()
		}
		return
	}

	itf, err := net.InterfaceByName(ifce.Name())
	if err != nil {
		return
	}

	addrs, _ := itf.Addrs()
	l.logger.Infof("name: %s, mac: %s, mtu: %d, addrs: %s",
		itf.Name, itf.HardwareAddr, itf.MTU, addrs)

	l.cqueue = make(chan net.Conn, 1)
	l.closed = make(chan struct{})

	conn := tap_util.NewConn(l.md.config, ifce, l.addr, &net.IPAddr{IP: ip})

	l.cqueue <- conn

	return
}

func (l *tapListener) Accept() (net.Conn, error) {
	select {
	case conn := <-l.cqueue:
		return conn, nil
	case <-l.closed:
	}

	return nil, listener.ErrClosed
}

func (l *tapListener) Addr() net.Addr {
	return l.addr
}

func (l *tapListener) Close() error {
	select {
	case <-l.closed:
		return net.ErrClosed
	default:
		close(l.closed)
	}
	return nil
}
