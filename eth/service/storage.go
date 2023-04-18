package service

import (
	storage "eth/abi"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type StorageContract struct {
	client *ethclient.Client
	token  *storage.StorageToken
	auth   *bind.TransactOpts
}

func NewStorageContract(client *ethclient.Client, contractAddrStr, privateKeyStr string, chainId int64) (*StorageContract, error) {
	//创建合约对象
	storageToken, err := storage.NewStorageToken(common.HexToAddress(contractAddrStr), client)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(chainId))
	if err != nil {
		return nil, err
	}

	return &StorageContract{
		client: client,
		token:  storageToken,
		auth:   auth,
	}, nil
}

func (s *StorageContract) DoRetrieve() (*big.Int, error) {
	res, err := s.token.Retrieve(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StorageContract) DoStore(num *big.Int) (*types.Transaction, error) {
	store, err := s.token.Store(&bind.TransactOpts{
		From:      s.auth.From,
		Nonce:     nil,
		Signer:    s.auth.Signer,
		Value:     nil,
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   nil,
		NoSend:    false,
	}, num)
	if err != nil {
		fmt.Println("store err", err)
		return store, err
	}

	return store, err
}
