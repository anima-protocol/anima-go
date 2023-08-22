package signature

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Signature struct {
	Algorithm    string       `json:"algorithm"`
	StdSignature StdSignature `json:"stdSignature"`
}

type StdSignature struct {
	PubKey    PubKey `json:"pub_key"`
	Signature string `json:"signature"`
}

type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func VerifyPersonalSignature(publicAddress string, data []byte, userSignature string) error {
	//convert userSignature hex string to bytes
	sigHex, err := hexutil.Decode(userSignature)
	if err != nil {
		return err
	}

	var sig Signature
	err = json.Unmarshal(sigHex, &sig)
	if err != nil {
		return err
	}

	hrp, _, err := bech32.Decode(publicAddress)
	if err != nil {
		return err
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(sig.StdSignature.PubKey.Value)
	if err != nil {
		return err
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(sig.StdSignature.Signature)
	if err != nil {
		return err
	}

	success, err := VerifyADR36Amino(hrp, publicAddress, data, pubKeyBytes, signatureBytes, sig.Algorithm)
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("signature verification failed")
	}

	return nil
}
