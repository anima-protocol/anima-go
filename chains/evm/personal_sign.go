package evm

import (
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func GetEVMPersonalSignMessage(data []byte) ([]byte, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	hash := ethCrypto.Keccak256Hash([]byte(fullMessage))

	return hash[:], nil
}

func VerifyPersonalSignature(publicAddress string, data []byte, userSignature string) error {
	message, err := GetEVMPersonalSignMessage(data)
	if err != nil {
		return err
	}

	return verifyEVMSignatureBase(publicAddress, message, userSignature)
}