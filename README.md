# go-passage
Passage delivery framework

## Description
Passage is designed to facilitate efficient message passing between multiple applications deployed in a cluster

## Protobuf
Compile and generate Protobuf libraries for master
```proto
protoc -I api/protobuf-spec/proto/master --go_out=plugins=grpc:api/protobuf-spec/bin/master/ api/protobuf-spec/proto/master/*.proto
```

## Dependencies
- Go-mysql [ref](https://github.com/go-sql-driver/mysql)
- protobuf

## Executing the code:
```go
go run main.go --stderrthreshold=INFO
```