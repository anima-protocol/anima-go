package secp256k1

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	crypto "github.com/anima-protocol/anima-go/crypto"
	"github.com/btcsuite/btcd/btcec"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"
	"math/big"
)

type PubKeySecp256k1 struct {
	pubKey []byte
}

func NewPubKeySecp256k1(pubKey []byte) *PubKeySecp256k1 {
	return &PubKeySecp256k1{pubKey: pubKey}
}

func (pk *PubKeySecp256k1) GetAddress() []byte {
	hasherSHA256 := sha256.New()
	_, _ = hasherSHA256.Write(pk.pubKey) // does not error
	sha := hasherSHA256.Sum(nil)

	hasherRIPEMD160 := ripemd160.New()
	_, _ = hasherRIPEMD160.Write(sha) // does not error

	return hasherRIPEMD160.Sum(nil)
}

func (pk *PubKeySecp256k1) GetEthAddress() []byte {
	// Should be uncompressed. .
	pubK, err := btcec.ParsePubKey(pk.pubKey, btcec.S256())
	if err != nil {
		return nil
	}

	address := ethCrypto.PubkeyToAddress(*pubK.ToECDSA())
	return address.Bytes()
}

func (pk *PubKeySecp256k1) Verify(msg, signature []byte) (bool, error) {
	digest := crypto.HashSHA256Bytes(msg)
	return pk.VerifyDigest32(digest[:], signature)
}

func (pk *PubKeySecp256k1) VerifyDigest32(digest, signature []byte) (bool, error) {
	if len(digest) != 32 {
		panic(fmt.Sprintf("Invalid length of digest to verify: %d", len(digest)))
	}

	if len(signature) != 64 {
		panic(fmt.Sprintf("Invalid length of signature: %d", len(signature)))
	}

	pubK, err := btcec.ParsePubKey(pk.pubKey, btcec.S256())
	if err != nil {
		return false, fmt.Errorf("Failed to parse public key: %v", err)
	}

	r, s := signature[:32], signature[32:]

	return ecdsa.Verify(pubK.ToECDSA(), digest, new(big.Int).SetBytes(r), new(big.Int).SetBytes(s)), nil
}
