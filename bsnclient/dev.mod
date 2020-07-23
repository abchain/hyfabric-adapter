module hyperledger.abchain.org/adapter/hyfabric/bsnclient

go 1.14

require (
	github.com/cloudflare/cfssl v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta2
	github.com/pkg/errors v0.8.1
	github.com/spf13/viper v1.2.1
	github.com/tjfoc/gmsm v1.3.2
	github.com/wonderivan/logger v1.0.0
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	google.golang.org/grpc v1.27.1
	hyperledger.abchain.org v0.2.0
	hyperledger.abchain.org/adapter/hyfabric v0.2.2
)

replace github.com/hyperledger/fabric-sdk-go => dev.stringon.com/abchain/fabric-sdk-go.git v1.0.0-beta2-priv
