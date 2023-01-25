package evm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(content []byte, sig string, issuerPublicAddress string) error {
	credentialContentBytes := new(bytes.Buffer)

	err := json.Compact(credentialContentBytes, content)
	if err != nil {
		return err
	}

	hashedMessage := ethCrypto.Keccak256Hash(credentialContentBytes.Bytes())

	return verifyEVMSignatureBase(issuerPublicAddress, hashedMessage.Bytes(), sig)
}

func SignData(data interface{}, signingFunc func([]byte) (string, error)) (string, error) {
	dataByte, err := json.Marshal(&data)
	if err != nil {
		return "", err
	}

	compactedDataByte := new(bytes.Buffer)
	err = json.Compact(compactedDataByte, dataByte)
	if err != nil {
		return "", err
	}

	hashedMessage := ethCrypto.Keccak256Hash(compactedDataByte.Bytes())

	signature, err := signingFunc(hashedMessage.Bytes())
	if err != nil {
		return "", err
	}

	return signature, nil
}

func verifyEVMSignatureBase(publicAddress string, data []byte, userSignature string) error {

	recoveredAddr, err := Recover(data, userSignature)
	if err != nil {
		return err
	}

	if !strings.EqualFold(recoveredAddr, publicAddress) {
		return fmt.Errorf("public address and signer address does not match")
	}

	fmt.Printf("--> All good!\n")

	return nil
}
