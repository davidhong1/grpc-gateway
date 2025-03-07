---
layout: default
title: FAQ
nav_order: 7
---

# FAQ

## How can I write the annotations which gRPC-Gateway requires?

The gRPC-Gateway follows the spec of [`google.api.HttpRule`](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto), so first check out the documentation if it is feasible in the spec.

For situations where annotating the proto file is not an option please reference the documentation on [gRPC API Configuration](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/)

See also [a past discussion](https://groups.google.com/d/msg/grpc-io/Xqx80hG0D44/VNCDHjeE6pUJ) in the grpc-io mailing list.

## I want to support a certain style of HTTP request but the code generated by gRPC-Gateway does not. How can I support this style?

See the question above at first.

The gRPC-Gateway is intended to cover 80% of use cases without forcing you to write comprehensive but complicated annotations. So the gateway itself does not always cover all the use cases you have by design. In other words, the gateway automates typical boring boilerplate mapping between gRPC and HTTP/1 communication, but it does not do arbitrarily complex custom mappings for you.

On the other hand, you can still add whatever you want as a middleware which wraps
[`runtime.ServeMux`](https://pkg.go.dev/github.com/davidhong1/grpc-gateway/runtime?tab=doc#ServeMux). Since `runtime.ServeMux` is just a standard [`http.Handler`](http://golang.org/pkg/http#Handler), you can easily write a custom wrapper of `runtime.ServeMux`, leveraged with existing third-party libraries in Go (e.g. [gateway main.go program](https://github.com/davidhong1/grpc-gateway/blob/main/examples/internal/gateway/main.go).

## My gRPC server is written in (Scala or C++ or Ruby or Haskell etc). Is there a (Scala or C++ or Ruby or Haskell etc) version of gRPC-Gateway?

As of now, No. But it should not be a big issue because the reverse-proxy which gRPC-Gateway generates usually works as an independent process and communicates with your gRPC server over TCP or a Unix domain sockets (Unix systems only).

## Why are the models in the OpenAPI specification prefixed with the last part of the proto package name?

The reason to generate the prefixes is that we don't have a guaranteed unique namespace. If two packages produce different `Foo` messages then we will have trouble.

## Why not strip the prefix?

When a message is added which happens to conflict with another message (e.g. by importing a message with the same name from a different package) it will break code that is very far away from the code that changed. This is in an effort to adhere to the [principle of least astonishment](https://en.wikipedia.org/wiki/Principle_of_least_astonishment).

## What is the difference between the gRPC-Gateway and grpc-httpjson-transcoding?

The gRPC-Gateway is a generator that generates a Go implementation of a JSON/HTTP-gRPC reverse proxy based on annotations in your proto file, while the [grpc-httpjson-transcoding](https://github.com/grpc-ecosystem/grpc-httpjson-transcoding) library doesn't require the generation step, it uses protobuf descriptors as config. It can be used as a component of an existing proxy. Google Cloud Endpoints and the gRPC-JSON transcoder filter in Envoy are using this.

<!-- TODO(v3): remove this note when default behavior matches Envoy/Cloud Endpoints -->
**Behavior differences:**
- By default, gRPC-Gateway does not escape path parameters in the same way. [This can be configured.](../mapping/customizing_your_gateway.md#Controlling-path-parameter-unescaping)

## What is the difference between the gRPC-Gateway and gRPC-web?

### Usage

In the gRPC-Gateway, we generate a reverse-proxy from the proto file annotations. In the front-end, we call directly through REST APIs. We can generate an OpenAPI v2 specification that may further be used to generate the frontend client from using `protoc-gen-openapiv2`.

In gRPC-web, the client code is generated directly from the proto files and can be used in the frontend.

### Performance

The gRPC-Gateway parses JSON to the protobuf binary format before sending it to the gRPC server. It then has to parse the reply back from the protobuf binary format to JSON again The parsing overhead has a negative impact on performance.

In gRPC-web, the message is sent in the protobuf binary format already, so there is no additional parsing cost on the proxy side.

### Maintenance

With the gRPC-Gateway, if your proto file changes, we have to regenerate the gateway reverse proxy code. If you are using the HTTP/JSON interface you probably have to change the front-end too, which means making changes in two places.

In gRPC-web, regenerating the files from the proto file will automatically update the front-end client.
