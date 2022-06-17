package registry

import "context"

type ServInfo struct {
	Addr     string
	Name     string
	Group    string
	Metadata map[string]string
}

type Registry interface {
	Registry(ctx context.Context, servInfo *ServInfo) error
	Deregister(ctx context.Context, servInfo *ServInfo) error
}
