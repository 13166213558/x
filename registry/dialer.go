package registry

import (
	"github.com/go-gost/core/dialer"
	"github.com/go-gost/core/logger"
)

type NewDialer func(opts ...dialer.Option) dialer.Dialer

type dialerRegistry struct {
	registry
}

func (r *dialerRegistry) Register(name string, v NewDialer) error {
	if err := r.registry.Register(name, v); err != nil {
		logger.Default().Fatal(err)
	}
	return nil
}

func (r *dialerRegistry) Get(name string) NewDialer {
	if v := r.registry.Get(name); v != nil {
		return v.(NewDialer)
	}
	return nil
}
