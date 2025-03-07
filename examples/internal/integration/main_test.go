package integration_test

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/davidhong1/grpc-gateway/v2/examples/internal/gateway"
	"github.com/davidhong1/grpc-gateway/v2/examples/internal/server"
	gwruntime "github.com/davidhong1/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/grpclog"
)

var (
	endpoint   = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	openAPIDir = flag.String("openapi_dir", "examples/internal/proto/examplepb", "path to the directory which contains OpenAPI definitions")
)

func runGateway(ctx context.Context, addr string, opts ...gwruntime.ServeMuxOption) error {
	return gateway.Run(ctx, gateway.Options{
		Addr: addr,
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
		OpenAPIDir: *openAPIDir,
		Mux:        opts,
	})
}

func waitForGateway(ctx context.Context, port uint16) error {
	ch := time.After(10 * time.Second)

	for {
		r, err := http.Get(fmt.Sprintf("http://localhost:%d/healthz", port))
		if err == nil && r.StatusCode == http.StatusOK {
			return nil
		}

		grpclog.Infof("Waiting for localhost:%d to get ready", port)
		select {
		case <-ctx.Done():
			return err
		case <-ch:
			return err
		case <-time.After(10 * time.Millisecond):
		}
	}
}

func runServers(ctx context.Context) <-chan error {
	ch := make(chan error, 3)
	go func() {
		if err := server.Run(ctx, *network, *endpoint); err != nil {
			ch <- fmt.Errorf("cannot run grpc service: %v", err)
		}
	}()
	go func() {
		if err := runGateway(ctx, ":8088"); err != nil {
			ch <- fmt.Errorf("cannot run gateway service: %v", err)
		}
	}()
	go func() {
		if err := server.RunInProcessGateway(ctx, ":8089"); err != nil {
			ch <- fmt.Errorf("cannot run in process gateway service: %v", err)
		}
	}()
	return ch
}

func TestMain(m *testing.M) {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := runServers(ctx)

	ch := make(chan int, 1)
	go func() {
		if err := waitForGateway(ctx, 8088); err != nil {
			grpclog.Errorf("waitForGateway(ctx, 8088) failed with %v; want success", err)
		}
		ch <- m.Run()
	}()

	select {
	case err := <-errCh:
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	case status := <-ch:
		cancel()
		os.Exit(status)
	}
}
