package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/samarthasthan/lockchain/api"
	"github.com/samarthasthan/lockchain/internal/controller"
	"github.com/samarthasthan/lockchain/internal/ethereum"
	"github.com/samarthasthan/lockchain/pkg/env"
)

var (
	CONTRACT_ADDRESS string
	URL              string
)

func init() {
	CONTRACT_ADDRESS = env.GetEnv("CONTRACT_ADDRESS", "0xaBcE422607945154E7806560d83C6B5393eDb4D0")
	URL = env.GetEnv("URL", "http://localhost:8545")
}

func main() {
	ec, err := ethclient.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	contractAddr := common.HexToAddress(CONTRACT_ADDRESS)
	contract, err := ethereum.NewEthereum(contractAddr, ec)
	if err != nil {
		panic(err)
	}

	controller := controller.NewController(ec, contract)

	// Start the server
	handlers := api.NewHandler(controller)
	handlers.Handle()
	handlers.Logger.Fatal(handlers.Start(":1248"))
}
