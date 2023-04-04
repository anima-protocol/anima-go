package signature

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/anima-protocol/anima-go/chains/cosmos/secp256k1"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/btcsuite/btcutil/bech32"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func toBech32(address []byte, prefix string) (string, error) {
	words, err := bech32.ConvertBits(address, 8, 5, false)
	if err != nil {
		return "", err
	}

	return bech32.Encode(prefix, words)
}

func MakeADR36AminoSignDoc(signer string, data []byte) models.StdSignDoc {
	// data b64 encode
	encodedData := base64.StdEncoding.EncodeToString(data)

	return models.StdSignDoc{
		ChainId:       "",
		AccountNumber: "0",
		Sequence:      "0",
		Fee:           models.StdFee{Gas: "0", Amount: []models.Coin{}},
		Msgs: []models.Msg{
			{
				Type:  "sign/MsgSignData",
				Value: models.MsgSignData{Signer: signer, Data: encodedData},
			},
		},
		Memo: "",
	}
}

func VerifyADR36AminoSignDoc(
	bech32PrefixAccAddr string,
	signDoc models.StdSignDoc,
	pubKey []byte,
	signature []byte,
	algo string,
) (bool, error) {
	cryptoPubkey := secp256k1.NewPubKeySecp256k1(pubKey)

	var addressToConvert []byte
	if algo == "ethsecp256k1" {
		addressToConvert = cryptoPubkey.GetEthAddress()
	} else {
		addressToConvert = cryptoPubkey.GetAddress()
	}

	expectedSigner, err := toBech32(addressToConvert, bech32PrefixAccAddr)
	if err != nil {
		return false, err
	}

	signer := signDoc.Msgs[0].Value.Signer

	if expectedSigner != signer {
		return false, fmt.Errorf("unmatched signer: %s != %s", expectedSigner, signer)
	}

	msg, err := json.Marshal(signDoc)
	if err != nil {
		return false, err
	}

	var digest []byte
	if algo == "ethsecp256k1" {
		digest = ethCrypto.Keccak256(msg)
	} else {
		digest = crypto.HashSHA256Bytes(msg)
	}

	return cryptoPubkey.VerifyDigest32(digest, signature)
}

func VerifyADR36Amino(
	bech32PrefixAccAddr string,
	signer string,
	data []byte,
	pubKey []byte,
	signature []byte,
	algo string,
) (bool, error) {
	signDoc := MakeADR36AminoSignDoc(signer, data)

	return VerifyADR36AminoSignDoc(
		bech32PrefixAccAddr,
		signDoc,
		pubKey,
		signature,
		algo,
	)
}
