package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/config"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var cfgSrcEnv *string

func init() {
	cfgSrcEnv = flag.String("env", "./.env", "ENV config source")
	flag.Parse()
}
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.NewConfig(*cfgSrcEnv)
	if err != nil {
		panic(err)
	}

	sp := NewServiceProvider(cfg)

	server := grpc.NewServer()
	server.RegisterService(&desc.ShopService_ServiceDesc, sp.GetProductServer())

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		fmt.Println("context cancelled")
		server.GracefulStop()
	}()

	server.Serve(lis)
}
