package transaction

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UserInfo struct {
	PrivateKey  *ecdsa.PrivateKey
	FromAddress common.Address
}

func GetClient() *ethclient.Client {
	rpcAddress := "https://rpc.ankr.com/arbitrum"
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func ContributeToPandaSale(amountInEth string, presaleAddress string, pK string) {

	amountInWei := FloatToBigInt(StringToFloat64(amountInEth))
	funcSig := "0xc1cbbca7" // function signature for contribute()
	// rawData = 0xc1cbbca700000000000000000000000000000000000000000000000001ade0dbe1d28000
	amountInWeiHex := BigIntToHex(amountInWei)
	//pad the amount to 64 bytes
	for len(amountInWeiHex) < 64 {
		amountInWeiHex = "0" + amountInWeiHex
	}
	rawData := funcSig + amountInWeiHex

	nonce := GetNonce(GetUserInfo(pK).FromAddress.Hex())
	gasprice, err := GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create a new transaction

	tx := types.NewTransaction(nonce, common.HexToAddress(presaleAddress), amountInWei, 600000, gasprice, common.FromHex(rawData))
	chainID, err := GetClient().NetworkID(context.Background())

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), GetUserInfo(pK).PrivateKey)
	if err != nil {
		log.Fatal("error signing transaction: ", err)
	}

	err = GetClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("error sending transaction: ", err)
	}
	// go awaitToConfirmTx(tx)
	fmt.Println("tx sent:", tx.Hash().Hex())

}

func EthToWei(amountInEth string) *big.Int {
	amountInWei := new(big.Int)
	amountInWei.SetString(amountInEth, 10)
	amountInWei.Mul(amountInWei, big.NewInt(1000000000000000000))
	return amountInWei
}

func GetUserInfo(pK string) UserInfo {

	PrivateKey, err := crypto.HexToECDSA(pK)
	if err != nil {
		log.Fatal("error casting private key to ECDSA")
	}

	publicKeyECDSA, ok := PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	FromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return UserInfo{
		PrivateKey:  PrivateKey,
		FromAddress: FromAddress,
	}
}

func GetNonce(address string) uint64 {
	nonce, err := GetClient().PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		log.Fatal(err)
	}
	return nonce
}

func StringToFloat64(value string) float64 {
	amountInEth, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err)
	}
	return amountInEth

}

func FloatToBigInt(amountInEth float64) *big.Int {
	amountInWei := new(big.Int)
	amountInWei.SetString(fmt.Sprintf("%.0f", amountInEth*1000000000000000000), 10)
	return amountInWei
}

func BigIntToHex(amountInWei *big.Int) string {
	return fmt.Sprintf("%x", amountInWei)
}
