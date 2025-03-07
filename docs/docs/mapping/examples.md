---
layout: default
title: Examples
nav_order: 0
parent: Mapping
---

# Examples

Examples are available under `examples/internal` directory.

- [`proto/examplepb/echo_service.proto`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/echo_service.proto), [`proto/examplepb/a_bit_of_everything.proto`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/a_bit_of_everything.proto), [`proto/examplepb/unannotated_echo_service.proto`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/unannotated_echo_service.proto):
  protobuf service definitions.
- [`proto/examplepb/echo_service.pb.go`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/echo_service.pb.go), [`proto/examplepb/a_bit_of_everything.pb.go`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/a_bit_of_everything.pb.go), [`proto/examplepb/unannotated_echo_service.pb.go`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/unannotated_echo_service.pb.go):
  generated Go service stubs and types.
- [`proto/examplepb/echo_service.pb.gw.go`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/echo_service.pb.gw.go), [`proto/examplepb/a_bit_of_everything.pb.gw.go`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/a_bit_of_everything.pb.gw.go), [`proto/examplepb/unannotated_echo_service.pb.gw.go`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/unannotated_echo_service.pb.gw.go):
  generated gRPC-Gateway clients.
- [`proto/examplepb/unannotated_echo_service.yaml`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/proto/examplepb/unannotated_echo_service.yaml):
  gRPC API Configuration for `unannotated_echo_service.proto`.
- [`server/main.go`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/server/main.go):
  service implementation.
- [`main.go`](https://github.com/davidhong1/grpc-gateway/tree/main/examples/internal/gateway/main.go):
  entrypoint of the generated reverse proxy.

To use the same port for custom HTTP handlers (e.g. serving `swagger.json`),
gRPC-Gateway, and a gRPC server, see [this code example by CoreOS](https://github.com/philips/grpc-gateway-example/blob/main/cmd/serve.go) (and it's accompanying
[blog post](https://coreos.com/blog/grpc-protobufs-swagger.html)).
