package main

import (
	"fmt"
	"github.com/spf13/viper"
	"hyperledger.abchain.org/adapter/hyfabric/bsnclient"
	"hyperledger.abchain.org/adapter/hyfabric/client"
)

func main() {
	fabricSdkTest()
}

func fabricSdkTest() {
	cfg := bsnclient.NewRPCConfig()
	defer cfg.Quit()
	vp := viper.New()
	vp.SetConfigName("demo") // name of config file (without extension)
	vp.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	vp.AddConfigPath(".")
	if err := vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err := cfg.Load(vp.Sub("fabric"))
	if err != nil {
		panic(err)
	}

	chanInfo, err := cfg.Chain()
	if err != nil {
		panic(err)
	}

	fmt.Println(chanInfo.GetChain())
}

func bsnClientTest() {
	cfg := client.NewRPCConfig()
	defer cfg.Quit()

	vp := viper.New()
	vp.SetConfigName("demo") // name of config file (without extension)
	vp.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	vp.AddConfigPath(".")
	if err := vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err := cfg.Load(vp.Sub("fabric"))
	if err != nil {
		panic(err)
	}
	chanInfo, err := cfg.Chain()
	if err != nil {
		panic(err)
	}
	height, err := chanInfo.GetChain()
	if err != nil {
		panic(err)
	}
	fmt.Println(height)

	for i := 0; i <= int(height.Height); i++ {
		bl, err := chanInfo.GetBlock(height.Height)
		if err != nil {
			panic(err)
		}
		var stat []int

		for _, event := range bl.TxEvents {
			stat = append(stat, event.Status)
		}
		fmt.Printf("block: %v, tx number: %v, stat: %v ", height.Height, len(bl.Transactions), stat)
	}
}
