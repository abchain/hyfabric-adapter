package main

import (
	"fmt"
	fchannel "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	fledger "github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	caMsp "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"hyperledger.abchain.org/adapter/hyfabric/client/channel"
	"hyperledger.abchain.org/adapter/hyfabric/client/ledger"
	"os"
)

type fabricClient struct {
	caClient *caMsp.Client
	sdk      *fabsdk.FabricSDK
	resMgmt  *resmgmt.Client
}

func Initialize(setup *fabricSetup) (*fabricClient, error) {
	fa := &fabricClient{}
	// Initialize the SDK with the configuration file
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return nil, fmt.Errorf("failed to create SDK: %v", err)
	}
	fa.sdk = sdk

	caClient, err := caMsp.New(sdk.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to create new CA client: %v", err)
	}
	fa.caClient = caClient
	fmt.Println("Initialization Successful")

	// The resource management client is responsible for managing channels (create/update channel)
	resourceManagerClientContext := fa.sdk.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName))
	if err != nil {
		return nil, err
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return nil, err
	}
	fa.resMgmt = resMgmtClient

	return fa, nil
}

// fabricSetup implementation
type fabricSetup struct {
	// Network parameters
	OrdererID string

	// Channel parameters
	ChannelID string
	// create channel
	// ChannelConfig string
	OrgAdmin   string
	OrgName    string
	ConfigFile string

	// ChainCode parameters
	ChaincodeID     string
	ChaincodeGoPath string
	ChaincodePath   string

	// CA parameters
	CaID string

	// User info
	user string
	pass string
}

func enrollUsr(client *fabricClient, setup *fabricSetup) error {
	_, err := client.caClient.GetSigningIdentity(setup.user)
	if err == caMsp.ErrUserNotFound {
		_, err = client.caClient.Register(&caMsp.RegistrationRequest{
			Name:   setup.user,
			Type:   "peer",
			Secret: setup.pass,
		})

		if err != nil {
			return err
		}

		err = client.caClient.Enroll(setup.user, caMsp.WithSecret(setup.pass))
		if err != nil {
			return err
		}
	}
	return err
}

func InstallAndInstantiateCC(client *fabricClient, setup *fabricSetup) error {
	// Create the chaincode package that will be sent to the peers
	ccPkg, err := packager.NewCCPackage(setup.ChaincodePath, setup.ChaincodeGoPath)
	if err != nil {
		return err
	}
	fmt.Println("ccPkg created")

	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: setup.ChaincodeID, Path: setup.ChaincodePath, Version: "0", Package: ccPkg}
	_, err = client.resMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return err
	}

	// Set up chaincode policy
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"org1.example.com"})

	resp, err := client.resMgmt.InstantiateCC(setup.ChannelID, resmgmt.InstantiateCCRequest{Name: setup.ChaincodeID, Path: setup.ChaincodeGoPath, Version: "0", Args: [][]byte{[]byte("init")}, Policy: ccPolicy})
	if err != nil || resp.TransactionID == "" {
		return err
	}
	fmt.Println("Chaincode Installation & Instantiation Successful")
	return nil
}

func main() {
	// Definition of the Fabric SDK properties
	fSetup := &fabricSetup{
		OrdererID:       "orderer.example.com",
		ChannelID:       "mychannel",
		OrgAdmin:        "User1",
		OrgName:         "org1",
		ConfigFile:      "demo/config.yaml",
		CaID:            "ca-org1",
		user:            "demo",
		pass:            "demo",
		ChaincodePath:   "hyperledger.abchain.org/adapter/hyfabric/demo/chaincode",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodeID:     "mycc",
	}
	faClient, err := Initialize(fSetup)
	if err != nil {
		panic(err)
	}

	err = enrollUsr(faClient, fSetup)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enroll User Success!")

	channelCtx := faClient.sdk.ChannelContext(fSetup.ChannelID, fabsdk.WithUser(fSetup.OrgAdmin), fabsdk.WithOrg(fSetup.OrgName))

	cli, err := fledger.New(channelCtx)
	if err != nil {
		panic(err)
	}
	ledgerCli := ledger.NewLedgerClient(cli)
	chain, err := ledgerCli.GetChain()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query Chain Info Success, Chain Height: %d\n", chain.Height)

	for i := 0; i < int(chain.Height); i++ {
		block, err := ledgerCli.GetBlock(int64(i))
		if err != nil {
			panic(err)
		}
		fmt.Printf("Query Block Success, Block Heiht: %d, Block Hash: %s, Block: %v\n", block.Height, block.Hash, block)
	}

	//err = InstallAndInstantiateCC(faClient, fSetup)
	//if err != nil {
	//	panic(err)
	//}

	fChanCli, err := fchannel.New(channelCtx)
	if err != nil {
		panic(err)
	}
	chanCli := channel.NewChannelClient(fSetup.ChaincodeID, fChanCli)
	res, err := chanCli.Query("query", [][]byte{[]byte("a")})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query Chain Code Success, Response: %v\n", string(res))

	// transfer 10 from a to b:wq
	txid, err := chanCli.Invoke("invoke", [][]byte{[]byte("a"), []byte("b"), []byte("10")})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invoke Chain Code Success, txid: %v\n", txid)

	res, err = chanCli.Query("query", [][]byte{[]byte("a")})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query Chain Code Success, Response: %v\n", string(res))
}
