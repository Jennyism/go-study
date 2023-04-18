package main

import (
	"context"
	"eth/service"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var auth *bind.TransactOpts

const (
	clientUrl       = "wss://eth-sepolia.g.alchemy.com/v2/qVwQQiyc2lxTMsY2rwcqAGAlDvcu0png"
	contractAddrStr = "0xF05689224039398aEa44B1E07Ff880Cf7bEAE297"
	privateKeyStr   = "ac888af1a65f6a88ab4aba785c2b617b923592254b2901574df1a62fe40744bf"
	chainId         = 11155111
)

func main() {
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	storageContract, err := service.NewStorageContract(client, contractAddrStr, privateKeyStr, chainId)

	store, err := storageContract.DoStore(big.NewInt(16))
	if err != nil {
		fmt.Printf("store err: %v\n", err)
		return
	}
	fmt.Printf("store success: %v\n", store)

	//retrieve, err := storageContract.DoRetrieve()
	//if err != nil {
	//	fmt.Printf("retrieve err: %v\n", err)
	//	return
	//}
	//fmt.Printf("retrieve result: %s", retrieve)
	listenLog(client, contractAddrStr)
}

func listenLog(client *ethclient.Client, contractAddrStr string) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddrStr)},
	}

	logs := make(chan types.Log)

	ctx := context.Background()
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalf("sub log err: %v", err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		case done := <-ctx.Done():
			fmt.Println("ctx done", done)
		}
	}
}
