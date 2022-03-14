package ws

import (
	"net/http"
	"time"

	mdata "github.com/go-gost/gost/v3/pkg/metadata"
)

const (
	defaultPath = "/ws"
)

type metadata struct {
	host string
	path string

	handshakeTimeout  time.Duration
	readHeaderTimeout time.Duration
	readBufferSize    int
	writeBufferSize   int
	enableCompression bool

	header    http.Header
	keepAlive time.Duration
}

func (d *wsDialer) parseMetadata(md mdata.Metadata) (err error) {
	const (
		host = "host"
		path = "path"

		handshakeTimeout  = "handshakeTimeout"
		readHeaderTimeout = "readHeaderTimeout"
		readBufferSize    = "readBufferSize"
		writeBufferSize   = "writeBufferSize"
		enableCompression = "enableCompression"

		header    = "header"
		keepAlive = "keepAlive"
	)

	d.md.host = mdata.GetString(md, host)

	d.md.path = mdata.GetString(md, path)
	if d.md.path == "" {
		d.md.path = defaultPath
	}

	d.md.handshakeTimeout = mdata.GetDuration(md, handshakeTimeout)
	d.md.readHeaderTimeout = mdata.GetDuration(md, readHeaderTimeout)
	d.md.readBufferSize = mdata.GetInt(md, readBufferSize)
	d.md.writeBufferSize = mdata.GetInt(md, writeBufferSize)
	d.md.enableCompression = mdata.GetBool(md, enableCompression)

	if m := mdata.GetStringMapString(md, header); len(m) > 0 {
		h := http.Header{}
		for k, v := range m {
			h.Add(k, v)
		}
		d.md.header = h
	}
	d.md.keepAlive = mdata.GetDuration(md, keepAlive)

	return
}
