# Kratos Project Template

## 官方参考 project
https://github.com/go-kratos/beer-shop
https://github.com/Windfarer/kratos-realworld

## 说明
- 基于官方 v2.1.5 修改 internal 结构
- 只在 api 中使用 protobuf 管理接口定义
- protobuf Request/Response 听取 windfarer 建议, 尽量每个接口都单独定义, 减少耦合
- 业务中重新定义的结构, 完全有业务需要定义, 虽然和 api proto 接口定义的一些字段重复,但是为了减少耦合, 减少接口之间的相互影响
- 结构类似于 beer-shop 中 [mono-layout](https://github.com/i6u/kratos-mono-layout) 结构.
- 参考 [mcube](https://github.com/infraboard/mcube) , 在 app 目录中管理业务模块

## proto generate

参考样例[greeter.proto](api/helloworld/v1/greeter.proto) 中 service, proto 使用 go_out, go-http_out, go-grpc_out, openapi_out 4 种插件
分别生成request/response struct, http, grpc, openapi doc 代码.
```
// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply)  {
        option (google.api.http) = {
            get: "/helloworld/{name}"
        };
    }
}
```

## [Install Kratos](https://github.com/go-kratos/kratos)
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new helloworld -r https://github.com/bagechashu/kratos-layout.git

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...

kratos run

go build -o ./bin/ ./...
./bin/server -conf ./config
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/config>:/data/conf <your-docker-image-name>
```

