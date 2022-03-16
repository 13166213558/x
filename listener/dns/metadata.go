package dns

import (
	"time"

	mdata "github.com/go-gost/core/metadata"
)

const (
	defaultBacklog = 128
)

type metadata struct {
	mode           string
	readBufferSize int
	readTimeout    time.Duration
	writeTimeout   time.Duration
	backlog        int
}

func (l *dnsListener) parseMetadata(md mdata.Metadata) (err error) {
	const (
		backlog        = "backlog"
		mode           = "mode"
		readBufferSize = "readBufferSize"
		readTimeout    = "readTimeout"
		writeTimeout   = "writeTimeout"
	)

	l.md.mode = mdata.GetString(md, mode)
	l.md.readBufferSize = mdata.GetInt(md, readBufferSize)
	l.md.readTimeout = mdata.GetDuration(md, readTimeout)
	l.md.writeTimeout = mdata.GetDuration(md, writeTimeout)

	l.md.backlog = mdata.GetInt(md, backlog)
	if l.md.backlog <= 0 {
		l.md.backlog = defaultBacklog
	}

	return
}
