package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	rpc_server "github.com/lightlykard/gobackend-2/hw9/pkg/grpc/pricerdr/server"
	"github.com/lightlykard/gobackend-2/hw9/pkg/storage"
)

func main() {
	mode := flag.String("mode", "", "set up mode, use n to set a new dir to save data")
	flag.Parse()
	dir := "./storage"
	st, err := storage.New(dir, *mode)
	if err != nil {
		log.Fatalf("can't set up storage %s", err)
	}
	server := rpc_server.New(st)
	prt := ":9000"
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go func() {
		rpc_server.Listen(ctx, server, prt)
	}()
	<-ctx.Done()
	server.Shutdown()
}
