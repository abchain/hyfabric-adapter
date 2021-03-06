module hyperledger.abchain.org/adapter/hyfabric

go 1.12

require (
	github.com/fsouza/go-dockerclient v1.6.5 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/hyperledger/fabric-amcl v0.0.0-20200424173818-327c9e2cf77a // indirect
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20200511190512-bcfeb58dd83a
	github.com/hyperledger/fabric-protos-go v0.0.0-20191121202242-f5500d5e3e85
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta2
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.2.1
	github.com/sykesm/zap-logfmt v0.0.3 // indirect
	go.uber.org/zap v1.15.0 // indirect
	hyperledger.abchain.org v0.2.0
)

replace (
	github.com/hyperledger/fabric-sdk-go => dev.stringon.com/abchain/fabric-sdk-go.git v1.0.0-beta2-priv
	hyperledger.abchain.org => dev.stringon.com/abchain/hyperledger.git v0.2.2
)
