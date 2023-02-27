package test

import (
	"didSample/abi/DidContract"
	"didSample/did"
	"didSample/eth"
	"didSample/log"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

// 网络rpc
const Rpc = ""

// did合约地址
var Address = common.HexToAddress("")

// 16进制私钥Hex
const PrivateKey = ""

// 链下生成did
func TestPreCreateDid(t *testing.T) {
	_, _, err := utils.CreateDid()
	if err != nil {
		log.Errorf(err.Error())
		return
	}
}

func TestCreateDid(t *testing.T) {
	didDocument, did, err := utils.CreateDid()
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("did document: %s\n", didDocument)
	log.Infof("did : %s\n", did)

	cli, err := ethclient.Dial(Rpc)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := DidContract.NewDidContract(Address, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	tx, err := instance.CreateDid(auth, did, didDocument)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("tx Hash: %s\n", tx.Hash().String())
}

func TestUpdateDidAuth(t *testing.T) {
	did := "did:bsn:2VwpU9b8cktekvGhPeD89gSt3EyR"
	privateKey := "09388158cf84cc1b4ada96cac2e912d3651cc428d3b179c6758ffc58e3eadc00"

	cli, err := ethclient.Dial(Rpc)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := DidContract.NewDidContract(Address, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	result, err := instance.GetDocument(nil, did)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	var didDocument utils.Document

	err = json.Unmarshal([]byte(result.DidDocument), &didDocument)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	err = didDocument.ResetDidAuthTest(privateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	didDocumentBytes, err := json.Marshal(didDocument)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	tx, err := instance.UpdateDocument(auth, did, string(didDocumentBytes))
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("tx Hash: %s\n", tx.Hash().String())
}

func TestGetDocument(t *testing.T) {
	did := "did:bsn:2VwpU9b8cktekvGhPeD89gSt3EyR"
	cli, err := ethclient.Dial(Rpc)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := DidContract.NewDidContract(Address, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	document, err := instance.GetDocument(&bind.CallOpts{}, did)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("did document: %s\n", document.DidDocument)
}
