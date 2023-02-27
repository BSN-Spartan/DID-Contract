/*
* @Author Duanraudon
* @Description
* @FileName create.go
* @ProductName GoLand
* @Date 2023/2/1 15:14
 */

package utils

import (
	"crypto/sha256"
	"didSample/log"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
	"strings"
	"time"
)

func GenerateDocument(did string, baseDidDocument *BaseDidDocument) Document {
	var didDocument = Document{
		Authentication: baseDidDocument.Authentication,
		Created:        time.Now().Format("2006-01-02 15:04:05"),
		Did:            did,
		Proof:          Proof{},
		Recovery:       baseDidDocument.Recovery,
		Updated:        time.Now().Format("2006-01-02 15:04:05"),
		Version:        DefaultVersion,
	}
	return didDocument
}

func GenerateKeyInfo() (KeyInfo, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return KeyInfo{}, err
	}
	privateKey := crypto.FromECDSA(key)
	publicKeyKey := crypto.FromECDSAPub(&key.PublicKey)

	privateKeyEncode := hex.EncodeToString(privateKey)
	publicKeyKeyEncode := hex.EncodeToString(publicKeyKey)

	var keyInfo = KeyInfo{
		PrivateKey: privateKeyEncode,
		PublicKey:  publicKeyKeyEncode,
		Types:      SECP256K1,
	}

	return keyInfo, nil
}

func GenerateDidIdentifierByBaseDidDocument(baseDidDocument *BaseDidDocument) (string, error) {
	marshal, err := json.Marshal(baseDidDocument)
	if err != nil {
		return "", err
	}

	hash := sha256.New()
	hash.Write(marshal)
	hashSum := hash.Sum(nil)
	md160 := ripemd160.New()
	md160.Write(hashSum)
	md160Sum := md160.Sum(nil)

	did := base58.Encode(md160Sum)

	return did, nil
}

func GenerateDidByDidIdentifier(didIdentifier string) string {
	var didFormatSplicing strings.Builder

	didFormatSplicing.WriteString(DidPrefix)
	didFormatSplicing.WriteString(DidSeparator)
	didFormatSplicing.WriteString(DidProjectName)
	didFormatSplicing.WriteString(DidSeparator)
	didFormatSplicing.WriteString(didIdentifier)
	s := didFormatSplicing.String()
	return s
}

func CreateDid() (string, string, error) {
	authenticationKeyPair, err := GenerateKeyInfo()
	if err != nil {
		return "", "", err
	}
	authenticationKeyPairBytes, err := json.Marshal(authenticationKeyPair)
	if err != nil {
		return "", "", err
	}

	log.Infof("生成主公私钥密钥对：%s\n", authenticationKeyPairBytes)

	recyKeyPair, err := GenerateKeyInfo()
	if err != nil {
		return "", "", err
	}

	recyKeyPairBytes, err := json.Marshal(recyKeyPair)
	if err != nil {
		return "", "", err
	}
	log.Infof("生成备公私钥密钥对：%s\n", recyKeyPairBytes)

	var authentication = Authentication{
		PublicKey: authenticationKeyPair.PublicKey,
		Types:     authenticationKeyPair.Types,
	}
	var recovery = Recovery{
		PublicKey: recyKeyPair.PublicKey,
		Type:      recyKeyPair.Types,
	}

	var baseDidDocument = &BaseDidDocument{
		Content,
		authentication,
		recovery,
	}

	baseDidDocumentBytes, err := json.Marshal(baseDidDocument)
	if err != nil {
		return "", "", err
	}
	log.Infof("拼接baseDocument， 结果为：%s\n", baseDidDocumentBytes)

	encodeDid, err := GenerateDidIdentifierByBaseDidDocument(baseDidDocument)
	if err != nil {
		return "", "", err
	}
	log.Infof("对baseDocument依次进行sha256，ripemd160，base58处理并生成did编码：%s\n", encodeDid)

	did := GenerateDidByDidIdentifier(encodeDid)

	log.Infof("拼接did标识符，拼接结果为： %s\n", did)

	document := GenerateDocument(did, baseDidDocument)

	documentBytes, err := json.Marshal(document)
	if err != nil {
		return "", "", err
	}
	log.Infof("初次拼接document，结果为%s\n", documentBytes)

	authenticationPri, err := crypto.HexToECDSA(authenticationKeyPair.PrivateKey)
	if err != nil {
		return "", "", err
	}

	//address := crypto.PubkeyToAddress(authenticationPri.PublicKey).String()

	documentKeccak256 := crypto.Keccak256(documentBytes)

	sign, err := crypto.Sign(documentKeccak256, authenticationPri)
	if err != nil {
		return "", "", err
	}

	signString := hex.EncodeToString(sign)
	log.Infof("对document进行哈希摘要算法计算并用主公私密钥对的私钥将其进行签名，并对签名结果进行哈希转换，结果为%s\n", signString)

	document.Proof.Creator = did
	document.Proof.SignatureValue = signString
	document.Proof.Type = SECP256K1

	log.Infof("将签名结果存入document里面\n")

	didDocumentS, err := json.Marshal(document)
	if err != nil {
		return "", "", err
	}
	log.Infof("生成document， 结果为：%s\n", string(didDocumentS))

	return string(didDocumentS), did, err
}
