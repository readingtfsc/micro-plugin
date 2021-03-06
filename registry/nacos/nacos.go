package nacos

import (
	"context"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/readingtfsc/micro-plugin/registry"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type register struct {
	cli naming_client.INamingClient
}

type Config struct {
	Namespace string
	Addr      string
}

func NewRegistry(ctx context.Context, config *Config) (registry.Registry, error) {
	reg, err := builder(ctx, config)
	if err != nil {
		return nil, err
	}
	return reg, nil
}

type addr struct {
	ip     string
	port   uint64
	scheme string
}

func address(addrs string) ([]addr, error) {

	result := make([]addr, 0)
	array := strings.Split(addrs, ",")
	for _, v := range array {
		url, err := url.Parse(v)
		if err != nil {
			return nil, err
		}
		port, err := strconv.Atoi(url.Port())
		if err != nil {
			return nil, err
		}
		result = append(result, addr{
			ip:     url.Host,
			port:   uint64(port),
			scheme: url.Scheme,
		})
	}
	return result, nil
}

func builder(ctx context.Context, config *Config) (*register, error) {

	if config == nil {
		return nil, nil
	}

	cc := constant.NewClientConfig(
		constant.WithNamespaceId(config.Namespace),
	)

	addrs, err := address(config.Addr)
	if err != nil {
		return nil, err
	}

	scs := make([]constant.ServerConfig, 0)
	for _, a := range addrs {
		sc := constant.NewServerConfig(
			a.ip, a.port,
			constant.WithScheme(a.scheme),
		)
		scs = append(scs, *sc)
	}

	cp := vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: scs,
	}

	cli, err := clients.NewNamingClient(cp)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	reg := &register{cli: cli}
	return reg, nil

}
