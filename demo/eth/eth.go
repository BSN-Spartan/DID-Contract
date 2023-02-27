package eth

import (
	"context"
	"didSample/log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GenAuth(cli *ethclient.Client, privateKey string) (auth *bind.TransactOpts, err error) {
	// 查询链ID
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	_privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	fromAddress := crypto.PubkeyToAddress(_privateKey.PublicKey)
	nonce, err := cli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	auth, err = bind.NewKeyedTransactorWithChainID(_privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))

	return
}
