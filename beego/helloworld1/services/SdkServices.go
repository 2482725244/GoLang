package services

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	conf2 "github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"helloworld/constant"
	"log"
)

var helloworldSession HelloWorldSession

func init() {
	configs, err := conf2.ParseConfigFile("sdk/config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	client, err := client.Dial((*conf2.Config)(&configs[0]))
	if err != nil {
		log.Fatal(err)
	}
	address := constant.ContractAddress

	contractAddress := common.HexToAddress(address)
	instance, err := NewHelloWorld(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	helloworldSession = HelloWorldSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
}

func GetServiceImpl() string {
	value, err := helloworldSession.Get()
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func SetServiceImpl(value string) {
	tx, receipt, err := helloworldSession.Set(value) // call set API
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())
}
