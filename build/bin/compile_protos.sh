#!/bin/sh

proto_imports="./proto:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src"

protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,primitive.md  --gogofaster_out=import_path=atomix/primitive,plugins=grpc:proto proto/atomix/primitive/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,headers.md    --gogofaster_out=Matomix/primitive/primitive.proto=github.com/atomix/api/proto/atomix/primitive,import_path=atomix/headers,plugins=grpc:proto proto/atomix/headers/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,controller.md --gogo_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Matomix/partition/partition.proto=github.com/atomix/api/proto/atomix/partition,import_path=atomix/controller,plugins=grpc:proto proto/atomix/controller/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,counter.md    --gogofaster_out=Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/counter,plugins=grpc:proto proto/atomix/counter/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,election.md   --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/election,plugins=grpc:proto proto/atomix/election/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,indexmap.md   --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/indexedmap,plugins=grpc:proto proto/atomix/indexedmap/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,leader.md     --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/leader,plugins=grpc:proto proto/atomix/leader/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,list.md       --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/list,plugins=grpc:proto proto/atomix/list/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,lock.md       --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/lock,plugins=grpc:proto proto/atomix/lock/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,log.md --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/log,plugins=grpc:proto proto/atomix/log/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,map.md        --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/map,plugins=grpc:proto proto/atomix/map/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,map.md        --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,Matomix/primitive/primitive.proto=github.com/atomix/api/proto/atomix/primitive,import_path=atomix/metadata,plugins=grpc:proto proto/atomix/metadata/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,session.md    --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/session,plugins=grpc:proto proto/atomix/session/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,set.md        --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/set,plugins=grpc:proto proto/atomix/set/*.proto
protoc -I=$proto_imports --doc_out=docs  --doc_opt=markdown,value.md      --gogofaster_out=Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Matomix/headers/headers.proto=github.com/atomix/api/proto/atomix/headers,import_path=atomix/value,plugins=grpc:proto proto/atomix/value/*.proto
