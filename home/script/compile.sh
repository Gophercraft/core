protoc --go_opt=module=github.com/Gophercraft/core/home/protocol --go-grpc_opt=module=github.com/Gophercraft/core/home/protocol -I schema schema/*.proto --go_out=. --go-grpc_out=.