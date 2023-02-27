/*
* @Author Duanraudon
* @Description
* @FileName update.go
* @ProductName GoLand
* @Date 2023/2/2 16:29
 */

package utils

import (
	"didSample/log"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
)

func (document *Document) ResetDidAuthTest(privateKey string) error {
	newRecyKeyInfo, err := GenerateKeyInfo()
	if err != nil {
		return err
	}
	newRecyKeyInfoBytes, err := json.Marshal(newRecyKeyInfo)
	if err != nil {
		return err
	}
	log.Infof("生成新的备公私密钥对，生成结果为%s", newRecyKeyInfoBytes)
	document.Recovery = Recovery{
		PublicKey: newRecyKeyInfo.PublicKey,
		Type:      SECP256K1,
	}

	documentBytes, err := json.Marshal(document)
	if err != nil {
		return err
	}

	authenticationPri, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err
	}

	// 对document进行哈希摘要算法计算
	documentKeccak256 := crypto.Keccak256(documentBytes)
	sign, err := crypto.Sign(documentKeccak256, authenticationPri)
	if err != nil {
		return err
	}

	signString := hex.EncodeToString(sign)

	log.Infof("对新document进行哈希摘要算法计算并用主公私密钥对的私钥将其进行签名，并对签名结果进行哈希转换，结果为%s\n", signString)

	document.Proof.SignatureValue = signString

	return nil
}
