# protoc

protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  

## 增加双端口

需要将googleapis下载到本地使用

```shell
tree proto 
proto
├── common.pb.go
├── common.proto
├── google
│   └── api
│       ├── annotations.proto
│       └── http.proto
├── readme.md
├── tag.pb.go
├── tag.pb.gw.go
├── tag.proto
└── tag_grpc.pb.go
```

````shell
 protoc -I ./proto/ \
 -I $GOPATH/src \
  --go_out ./proto --go_opt paths=source_relative \
  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
./proto/*.proto
``` 