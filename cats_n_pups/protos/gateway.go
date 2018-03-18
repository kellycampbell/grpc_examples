package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	catpic "github.com/kellycampbell/grpc_examples/cats_n_pups/protos/catpic"
	puppypic "github.com/kellycampbell/grpc_examples/cats_n_pups/protos/puppypic"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	endpoint = flag.String(
		"endpoint",
		"localhost:50051",
		"gRPC service endpoint (default: localhost:50051)",
	)
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := catpic.RegisterCatPicsServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		glog.Error("couldn't register service handler endpoint: %v", err)
		return err
	}
	err = puppypic.RegisterPuppyPicsServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		glog.Error("couldn't register service handler endpoint: %v", err)
		return err
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		glog.Error("couldn't listen and serve: %v", err)
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	glog.Info("Service starting up on port 8080...")
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
