package elrond

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/bech32"
)

const ElrondHrp = "erd"

func VerifyPersonalSignature(publicAddress string, data []byte, userSignature string) error {
	if len(userSignature) < 3 {
		return fmt.Errorf("invalid signature length: %d", len(userSignature)) // TODO: Replace with APIError ?
	}

	if userSignature[0:2] == "0x" {
		userSignature = userSignature[2:]
	}

	message, err := GetERDPersonalSignMessage(data)
	if err != nil {
		return err
	}

	sigBytes, err := hex.DecodeString(userSignature)
	if err != nil {
		return err
	}

	hrp, decodedBech32Addr, err := bech32.Decode(publicAddress)
	if err != nil {
		return err
	}

	if hrp != ElrondHrp {
		return fmt.Errorf("invalid hrp from public addres: %s", hrp)
	}

	converted, err := bech32.ConvertBits(decodedBech32Addr, 5, 8, false)
	if err != nil {
		return err
	}
	valid := ed25519.Verify(converted, message, sigBytes)
	if !valid {
		return fmt.Errorf("unable to verify signature")
	}

	return nil
}
