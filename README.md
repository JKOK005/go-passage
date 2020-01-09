# go-passage
Passage delivery framework

## Description
Passage is designed to facilitate efficient message passing between multiple applications deployed in a cluster

## PRotobuf
Compile and generate Protobuf libraries for master
```proto
protoc -I api/protobuf-spec/proto/master --go_out=plugins=grpc:api/protobuf-spec/bin/master/ api/protobuf-spec/proto/master/*.proto
```