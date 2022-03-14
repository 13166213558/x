package mws

import (
	"net/http"
	"time"

	mdata "github.com/go-gost/gost/v3/pkg/metadata"
)

const (
	defaultPath    = "/ws"
	defaultBacklog = 128
)

type metadata struct {
	path    string
	backlog int
	header  http.Header

	handshakeTimeout  time.Duration
	readHeaderTimeout time.Duration
	readBufferSize    int
	writeBufferSize   int
	enableCompression bool

	muxKeepAliveDisabled bool
	muxKeepAliveInterval time.Duration
	muxKeepAliveTimeout  time.Duration
	muxMaxFrameSize      int
	muxMaxReceiveBuffer  int
	muxMaxStreamBuffer   int
}

func (l *mwsListener) parseMetadata(md mdata.Metadata) (err error) {
	const (
		path    = "path"
		backlog = "backlog"
		header  = "header"

		handshakeTimeout  = "handshakeTimeout"
		readHeaderTimeout = "readHeaderTimeout"
		readBufferSize    = "readBufferSize"
		writeBufferSize   = "writeBufferSize"
		enableCompression = "enableCompression"

		muxKeepAliveDisabled = "muxKeepAliveDisabled"
		muxKeepAliveInterval = "muxKeepAliveInterval"
		muxKeepAliveTimeout  = "muxKeepAliveTimeout"
		muxMaxFrameSize      = "muxMaxFrameSize"
		muxMaxReceiveBuffer  = "muxMaxReceiveBuffer"
		muxMaxStreamBuffer   = "muxMaxStreamBuffer"
	)

	l.md.path = mdata.GetString(md, path)
	if l.md.path == "" {
		l.md.path = defaultPath
	}

	l.md.backlog = mdata.GetInt(md, backlog)
	if l.md.backlog <= 0 {
		l.md.backlog = defaultBacklog
	}

	l.md.handshakeTimeout = mdata.GetDuration(md, handshakeTimeout)
	l.md.readHeaderTimeout = mdata.GetDuration(md, readHeaderTimeout)
	l.md.readBufferSize = mdata.GetInt(md, readBufferSize)
	l.md.writeBufferSize = mdata.GetInt(md, writeBufferSize)
	l.md.enableCompression = mdata.GetBool(md, enableCompression)

	l.md.muxKeepAliveDisabled = mdata.GetBool(md, muxKeepAliveDisabled)
	l.md.muxKeepAliveInterval = mdata.GetDuration(md, muxKeepAliveInterval)
	l.md.muxKeepAliveTimeout = mdata.GetDuration(md, muxKeepAliveTimeout)
	l.md.muxMaxFrameSize = mdata.GetInt(md, muxMaxFrameSize)
	l.md.muxMaxReceiveBuffer = mdata.GetInt(md, muxMaxReceiveBuffer)
	l.md.muxMaxStreamBuffer = mdata.GetInt(md, muxMaxStreamBuffer)

	if mm := mdata.GetStringMapString(md, header); len(mm) > 0 {
		hd := http.Header{}
		for k, v := range mm {
			hd.Add(k, v)
		}
		l.md.header = hd
	}
	return
}
