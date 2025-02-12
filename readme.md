## SEI MEV Protobuf definition in Go

Use 
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```

To regenerate run

```
git submodule update
buf dep update
buf generate
```
