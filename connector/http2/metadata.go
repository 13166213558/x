package http2

import (
	"net/http"
	"time"

	mdata "github.com/go-gost/core/metadata"
)

type metadata struct {
	connectTimeout time.Duration
	header         http.Header
}

func (c *http2Connector) parseMetadata(md mdata.Metadata) (err error) {
	const (
		connectTimeout = "timeout"
		header         = "header"
	)

	c.md.connectTimeout = mdata.GetDuration(md, connectTimeout)
	if mm := mdata.GetStringMapString(md, header); len(mm) > 0 {
		hd := http.Header{}
		for k, v := range mm {
			hd.Add(k, v)
		}
		c.md.header = hd
	}
	return
}
