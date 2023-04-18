package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func main() {
	privateKey, err := crypto.HexToECDSA("ac888af1a65f6a88ab4aba785c2b617b923592254b2901574df1a62fe40744bf")
	if err != nil {
		return
	}

	if keyjson, err := keystore.EncryptKey(&keystore.Key{
		Id:         uuid.UUID{},
		Address:    common.Address{},
		PrivateKey: privateKey,
	}, "snoopy9121++", keystore.LightScryptN, keystore.LightScryptP); err != nil {
		fmt.Printf("err: %v", err)
	} else {
		fmt.Println(string(keyjson))
		fromKey, err := keystore.DecryptKey([]byte(keyjson), "snoopy9121++")
		fmt.Printf("decryptkey err: %v\n", err)
		fmt.Printf("%s", crypto.PubkeyToAddress(fromKey.PrivateKey.PublicKey))
	}

}
