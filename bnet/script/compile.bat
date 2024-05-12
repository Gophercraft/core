protoc --go_opt=module=github.com/Gophercraft/core/bnet -I bnet\proto\login bnet\proto\login\*.proto --go_out=bnet
protoc --go_opt=module=github.com/Gophercraft/core/bnet --go-bnet_opt=module=github.com/Gophercraft/core/bnet -I bnet\proto bnet\proto\bgs\low\pb\client\*.proto bnet\proto\bgs\low\pb\client\global_extensions\*.proto bnet\proto\bgs\low\pb\client\api\client\v1\*.proto bnet\proto\bgs\low\pb\client\api\client\v2\*.proto --go_out bnet\ --go-bnet_out bnet\
protoc --go_opt=module=github.com/Gophercraft/core/bnet -I bnet\proto\realmlist bnet\proto\realmlist\*.proto --go_out=bnet
