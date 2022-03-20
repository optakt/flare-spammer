package spammer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Spammer interface {
	// Create numTx number of transactions by creating numTx number of new addresses to which the
	// root address will be sending transactions
	CreateRandomTransactions(numTx int) error
}

type NumberSpammer struct {
}

func NewSpammer() *NumberSpammer {
	return &NumberSpammer{}
}

func (s NumberSpammer) CreateRandomTransactions(numTx int) error {
	client, err := ethclient.Dial("http://localhost:9650/ext/bc/C/rpc")
	if err != nil {
		fmt.Println("Error in dialing")
		log.Fatal(err)
	}
	//6b0dd034A2FD67b932F10E3dBA1d2bbD39348695
	//c5e8f61d1ab959b397eecc0a37a6517b8e67a0e7cf1f4bce5591f3ed80199122
	//0xc783df8a850f42e7F7e57013759C285caa701eB6 -> public key
	privateKey, err := crypto.HexToECDSA("c5e8f61d1ab959b397eecc0a37a6517b8e67a0e7cf1f4bce5591f3ed80199122")
	if err != nil {
		fmt.Println("Error in privateKey")
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error in publicKey")
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("public key is: ", fromAddress) // this is used in genesis

	for i := 0; i < numTx; i++ {
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			fmt.Println("Error in nonce")
			log.Fatal(err)
		}

		value := big.NewInt(1000)  // in wei (1 eth)
		gasLimit := uint64(210000) // in units
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println("Error in gasPrice")
			log.Fatal(err)
		}
		gasPrice = big.NewInt(250000000000)
		toAddress := generateNewAddressesAndSendTx()
		var data []byte
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

		chainID, err := client.ChainID(context.Background()) // TODO figure out why this is being wrong! i.e. giving 1 instead of 20210406 -> Done! Since we were using another function (NetworkID) and not ChainID() function
		if err != nil {
			fmt.Println("Error in chainID")
			log.Fatal(err)
		}
		fmt.Println("chainID is: ", chainID.String()) // should be 20210406
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			fmt.Println("Error in signedTx")
			log.Fatal(err)
		}

		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		if err != nil {
			fmt.Println("Error in Balance")
			log.Fatal(err)
		}
		fmt.Println("balance is: ", balance.String())

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Println("Error in SendTransaction")
			log.Fatal(err) // todo break out instead of fatal
		}

		fmt.Printf("%dth tx sent: %s", i, signedTx.Hash().Hex())

	}
	return err
}

func generateNewAddressesAndSendTx() common.Address {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("SAVE BUT DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKeyBytes))

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key:", hexutil.Encode(publicKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex() // this is the address to which we can send money
	fmt.Println("Address:", address)
	addressToBeSent := crypto.PubkeyToAddress(*publicKeyECDSA)
	return addressToBeSent
}
