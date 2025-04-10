module github.com/SiloMEV/silo-mev-protobuf-go

go 1.21

require (
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/grpc v1.69.2 // indirect
	google.golang.org/protobuf v1.36.2 // indirect
	github.com/gogo/gateway v1.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)