/*
* @Author Duanraudon
* @Description
* @FileName base.go
* @ProductName GoLand
* @Date 2023/2/2 16:30
 */

package utils

const (
	SECP256K1      = "Secp256k1"
	Content        = "https://w3id.org/did/v1"
	DidPrefix      = "did"
	DidProjectName = "bsn"
	DidSeparator   = ":"
	DefaultVersion = "1"
)

type DidDocument struct {
	Address     string   `json:"address"`
	AuthKeyInfo KeyInfo  `json:"authKeyInfo"`
	Did         string   `json:"did"`
	DidSign     string   `json:"didSign"`
	Document    Document `json:"document"`
	RecyKeyInfo KeyInfo  `json:"recyKeyInfo"`
}

type Document struct {
	Authentication Authentication `json:"authentication"`
	Created        string         `json:"created"`
	Did            string         `json:"did"`
	Proof          Proof          `json:"proof"`
	Recovery       Recovery       `json:"recovery"`
	Updated        string         `json:"updated"`
	Version        string         `json:"version"`
}

type BaseDidDocument struct {
	Context        string         `json:"context,omitempty"`
	Authentication Authentication `json:"authentication"`
	Recovery       Recovery       `json:"recovery"`
}

type Authentication struct {
	PublicKey string `json:"publicKey"`
	Types     string `json:"types"`
}
type Proof struct {
	Creator        string `json:"creator"`
	SignatureValue string `json:"signatureValue"`
	Type           string `json:"type"`
}
type Recovery struct {
	PublicKey string `json:"publicKey"`
	Type      string `json:"type"`
}

type KeyInfo struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Types      string `json:"types"`
}

type ResetDidAuth struct {
	Did            string  `json:"did,omitempty"`
	PrimaryKeyPair KeyInfo `json:"primaryKeyPair"`
	RecoveryKey    KeyInfo `json:"recoveryKey"`
}
