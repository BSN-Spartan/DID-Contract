# DID-Contract
[![Smart Contract](https://badgen.net/badge/smart-contract/Solidity/orange)](https://soliditylang.org/) 
 
Digital identity is the body of information about an individual, organization or electronic device that exists online. It forms a public key that can be queried and identified through the network or related equipment. With the emergence and popularity of the Internet, it is generally believed that the evolution of digital identity has gone through three stages: centralized identity, alliance identity, and decentralized identity. Ideally, in the decentralized identity stage, users can fully control their own information.

In the W3C Recommendation, DID (Decentralized Identifiers) is defined as a new globally unique identifier, which can achieve verifiable and decentralized digital identity. This identifier can be used not only for people, but for everything, such as a car, an animal, and even a machine. [Click to learn more about DID.](https://www.w3.org/TR/2022/REC-did-core-20220719/#abstract)

## Prerequisite

Before using a smart contract, it is important to have a basic understanding of Ethereum and Solidity.

## Overview 

* The DID cannot perform real-name authentication. You can implement real-name authentication function in the business system, and then bind the user in the system to the DID.

* If the same sender calls the "Create DID" function multiple times, the smart contract will generate multiple different DIDs. If you need to ensure that each user has only one DID, you can logically determine whether this sender has been bound to a DID in the business system.

## Usage

Get the DID contract source code by command:

```
$ git clone 请按需补充代码拉取地址！！！！
```

For beginners, the contracts in this application can be deployed by the steps in [Spartan Quick Testing](https://www.spartan.bsn.foundation/main/quick-testing#step1).

In this application, there are two contracts: DidContract and LibBytesMap. Follow steps below to deploy and use them:

1. Deploy LibBytesMap contract.
2. Deploy DidContract contract.
3. In DidContract contract, call `createDid` function to upload the DID information to the chain. There are two parameters needed in this function: DID and DID Document. You can generate the DID information offline by following the W3C [Decentralized Identifiers (DIDs) v1.0](https://www.w3.org/TR/2022/REC-did-core-20220719/#abstract) document. In this application, we provide a demo to generate the DID information. The demo is developed by Golang, and here is an example of DID information:

```
"did": "did:bsn:Hpd2VT61i45Zg1HvrDZSy1KXzuu"
"document": {
		"authentication": {
			"publicKey": "4071408688821737472790885138742742191513769530255358839905040346912264139845565859273655612205467385164603960860045145974595099920266596559138013858407430",
			"type": "Secp256k1"
		},
		"created": "2023-02-27 14:40:02",
		"did": "did:bsn:Hpd2VT61i45Zg1HvrDZSy1KXzuu",
		"proof": {
			"creator": "did:bsn:Hpd2VT61i45Zg1HvrDZSy1KXzuu",
			"signatureValue": "n96yvQ26HhwKTlOd6c8Q6jlqbHUZHhsRMSjBKNJC1gNFRDn12bJgRWXMKzQg86RG5kRblXsYAD4/Wib3vPfNWAA=",
			"type": "Secp256k1"
		},
		"recovery": {
			"publicKey": "8895525627127748663885857832812217274835025000271644880797853149272014876033685686567538414189810304729773362512771498170772617685995518726229509125075455",
			"type": "Secp256k1"
		},
		"updated": "2023-02-27 14:40:02",
		"version": "1"
	}

```

4. After successfully uploading the DID information, we can fill in the DID in `getDocument` function to get the Document information.
5. In case the private key that signs the DID Document is lost or exposed to others, you can generate another private and public key pair and update the DID Document, and upload the new Document to the chain by `updateDocument` function.


### Main Functions

### createDid

"createDid" is a public function. It will generate a DID identifier and the description document of the sender. The DID and document are one-to-one relationships and stored in the form of Key-Value on the chain.

```
function createDid(string memory did, string memory didDocument)
```

### getDocument

"getDocument" is a public function. Anyone can query the DID Document on the chain by its corresponding DID. After obtaining the DID Document, you can check the creation date, public key and other information of the DID, so as to further verify the signature to ensure the integrity and authenticity of the transmitted data.

```
function getDocument(string memory did)
```


### updateDocument

"updateDocument" is a public function. If the authentication private key of the DID is lost or exposed to others, you can update the DID Document by a new private and public key pair. After updating the authentication public key and the signature in the DID Document, you can submit it to the chain through this function. The sender will be verified internally in the contract. If it is the same sender with the one creating the DID, the document will be updated; otherwise, the update will be failed.

```
function updateDocument(string memory did, string memory didDocument)
```


## License

Spartan-DID Contract is released under the [Spartan License](https://github.com/BSN-Spartan/Beginner-Level-Contracts/blob/main/Spartan%20License.md).
