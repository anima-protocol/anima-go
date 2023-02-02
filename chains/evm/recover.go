package evm

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func Recover(data []byte, userSignature string) (string, error) {

	if len(userSignature) < 3 {
		return "", fmt.Errorf("invalid signature length: %d", len(userSignature)) // TODO: Replace with APIError ?
	}

	if userSignature[0:2] == "0x" {
		userSignature = userSignature[2:]
	}

	signature, err := hex.DecodeString(userSignature)
	if err != nil {
		return "", err
	}

	if len(signature) != 65 {
		return "", fmt.Errorf("invalid signature length: %d", len(signature))
	}

	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}

	if signature[64] != 0 && signature[64] != 1 {
		return "", fmt.Errorf("invalid recovery id: %d", signature[64])
	}

	pubKeyRaw, err := crypto.Ecrecover(data, signature)
	if err != nil {
		return "", fmt.Errorf("invalid signature: %s", err.Error())
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return "", err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	fmt.Printf("address = %v\n", recoveredAddr.String())
	return strings.ToLower(recoveredAddr.String()), nil
}
