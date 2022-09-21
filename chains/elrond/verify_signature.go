package evm

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
)

const ElrondHrp = "erd"

func VerifySignature(publicAddress string, data []byte, userSignature string) (bool, error) {
	fmt.Printf("check signature length\n")
	if len(userSignature) < 3 {
		return false, fmt.Errorf("invalid signature length: %d", len(userSignature)) // TODO: Replace with APIError ?
	}

	fmt.Printf("check signature start with '0x'\n")
	if userSignature[0:2] == "0x" {
		userSignature = userSignature[2:]
	}

	fmt.Printf("fetch ERD712 message\n")
	message, err := GetERD712Message(data)
	if err != nil {
		return false, err
	}

	fmt.Printf("fetch_signature = %v\n", userSignature)
	sigBytes, err := hex.DecodeString(userSignature)
	if err != nil {
		return false, err
	}

	hrp, decodedBech32Addr, err := bech32.Decode(publicAddress)
	if err != nil {
		return false, err
	}
	if hrp != ElrondHrp {
		return false, fmt.Errorf("invalid hrp from public addres: %s", hrp)
	}

	valid := ed25519.Verify(decodedBech32Addr, message, sigBytes)
	if !valid {
		return false, fmt.Errorf("unable to verify signature")
	}

	return true, nil
}
