package main

import (
	"fmt"
	"github.com/spf13/viper"
	"hyperledger.abchain.org/adapter/hyfabric/client"
)

func main() {
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
}
