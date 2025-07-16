package cmd

import (
	"blog-tech/common"
	sctx "blog-tech/plugin"
	"flag"
)

type config struct {
	grpcPort              int
	grpcServerAddress     string
	grpcUserAddress       string
	grpcCategoryAddress   string
	grpcTagAddress        string
	grpcArticleTagAddress string
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ID() string {
	return common.KeyCompConf
}

func (c *config) InitFlags() {
	flag.IntVar(
		&c.grpcPort,
		"grpc-port",
		3100,
		"gRPC Port. Default: 3100",
	)

	flag.StringVar(
		&c.grpcServerAddress,
		"grpc-server-address",
		"localhost:3101",
		"gRPC server address. Default: localhost:3101",
	)

	flag.StringVar(
		&c.grpcUserAddress,
		"grpc-user-address",
		"localhost:3201",
		"gRPC user address. Default: localhost:3201",
	)

	flag.StringVar(
		&c.grpcCategoryAddress,
		"grpc-category-address",
		"localhost:3301",
		"gRPC category address. Default: localhost:3301",
	)

	flag.StringVar(
		&c.grpcTagAddress,
		"grpc-tag-address",
		"localhost:3401",
		"gRPC tag address. Default: localhost:3401",
	)

	flag.StringVar(
		&c.grpcArticleTagAddress,
		"grpc-article-tag-address",
		"localhost:3501",
		"gRPC article tag address. Default: localhost:3501",
	)
}

func (c *config) Activate(_ sctx.ServiceContext) error {
	return nil
}

func (c *config) Stop() error {
	return nil
}

func (c *config) GetGRPCPort() int {
	return c.grpcPort
}

func (c *config) GetGRPCServerAddress() string {
	return c.grpcServerAddress
}

func (c *config) GetGRPCUserAddress() string {
	return c.grpcUserAddress
}

func (c *config) GetGRPCCategoryAddress() string {
	return c.grpcCategoryAddress
}

func (c *config) GetGRPCTagAddress() string {
	return c.grpcTagAddress
}

func (c *config) GetGRPCArticleTagAddress() string {
	return c.grpcArticleTagAddress
}
