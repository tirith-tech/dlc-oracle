
check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

generate-swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./docs/swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger docs/swagger.yaml

generate-protobuf: 
	protoc rpc/protobuf/oracle.proto --go_out=plugins=grpc:.