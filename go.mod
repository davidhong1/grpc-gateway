module github.com/davidhong1/grpc-gateway/v2

go 1.19

replace github.com/davidhong1/grpc-gateway/v2 => ./

require (
	github.com/antihax/optional v1.0.0
	github.com/google/go-cmp v0.6.0
	github.com/rogpeppe/fastuuid v1.2.0
	golang.org/x/oauth2 v0.16.0
	golang.org/x/text v0.14.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.32.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20240116215550-a9fa1716bcac // indirect
)
