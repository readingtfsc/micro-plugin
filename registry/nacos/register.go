package nacos

import (
	"context"
	"micro-plugin/registry"
)

func (r *register) Registry(ctx context.Context, servInfo *registry.ServInfo) error {
	//TODO
	return nil
}

func (r *register) Deregister(ctx context.Context) error {
	//TODO
	return nil
}
