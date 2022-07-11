package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/anima-protocol/anima-go/chains/evm"
	animaCrypto "github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type IssuingAuthorization2 struct {
	Specs     string `json:"specs"`
	Content   string `json:"content"`
	Signature string `json:"signature"`
}

func loadEnv2() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load '.env' file, terminating.")
	}
}

func getEnv2(evar string) string {
	val := os.Getenv(evar)
	if val == "" {
		log.Fatalf("[MISSING EVAR] %v\n", evar)
		os.Exit(1)
	}
	return val
}

func main() {
	loadEnv2()
	PRIVATE_OWNER_SIGNING_KEY := getEnv2("PRIVATE_OWNER_SIGNING_KEY")
	PRIVATE_VERIFIER_SIGNING_KEY := getEnv2("PRIVATE_VERIFIER_SIGNING_KEY")

	sharing_auth := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    models.PROTOCOL_NAME,
			Version: models.PROTOCOL_VERSION,
			ChainId: math.NewHexOrDecimal256(1),
		},
		PrimaryType: "Main",
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "version",
					Type: "string",
				},
			},
			"Main": []apitypes.Type{
				{
					Name: "specs",
					Type: "string",
				},
				{
					Name: "shared_at",
					Type: "uint64",
				},
				{
					Name: "attributes",
					Type: "Attributes",
				},
				{
					Name: "verifier",
					Type: "Verifier",
				},
				{
					Name: "owner",
					Type: "Owner",
				},
			},
			"Owner": []apitypes.Type{
				{
					Name: "id",
					Type: "string",
				},
				{
					Name: "public_address",
					Type: "address",
				},
				{
					Name: "chain",
					Type: "string",
				},
				{
					Name: "wallet",
					Type: "string",
				},
			},
			"Verifier": []apitypes.Type{
				{
					Name: "id",
					Type: "string",
				},
				{
					Name: "public_address",
					Type: "address",
				},
				{
					Name: "chain",
					Type: "string",
				},
			},
			"Attributes": []apitypes.Type{
				{
					Name: "nationality",
					Type: "string",
				},
			},
		},
		Message: apitypes.TypedDataMessage{
			"specs":     "anima:specs:sharing/authorization@1.0.0",
			"shared_at": uint64(time.Now().Unix()),
			"attributes": apitypes.TypedDataMessage{
				"nationality": "anima:credential:87d0e5373127045c02cd817569d66042727295cafb38b525c410266f4f653e19",
			},
			"owner": apitypes.TypedDataMessage{
				"id":             "anima:owner:0x017f912f75c4140699606Ddb8418Ec944AAbCEBA",
				"public_address": "0x017f912f75c4140699606Ddb8418Ec944AAbCEBA",
				"chain":          "ETH",
				"wallet":         "METAMASK",
			},
			"verifier": apitypes.TypedDataMessage{
				"id":             "anima:verifier:syn_slash_bank@1.0.0",
				"public_address": "0x168FE97dCAd13e39838FB0e543d8A221904cE5BA",
				"chain":          "ETH",
			},
		},
	}

	challenge, err := json.Marshal(sharing_auth)
	if err != nil {
		log.Fatal("Error while transforming authorization in json")
		os.Exit(1)
	}

	challengeHash, _ := evm.GetEIP712Message(challenge)

	privateKey, _ := crypto.HexToECDSA(PRIVATE_OWNER_SIGNING_KEY)
	signature, err := crypto.Sign(challengeHash, privateKey)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	hexSignature := "0x" + hex.EncodeToString(signature)

	base64Challenge := base64.StdEncoding.EncodeToString(challenge)

	result := IssuingAuthorization2{
		Specs:     "anima:specs:sharing/authorization/eip712@1.0.0",
		Content:   base64Challenge,
		Signature: hexSignature,
	}

	finalAuthorization := make(map[string]interface{})
	finalAuthorization["authorization"] = result

	jsonResult, _ := json.Marshal(finalAuthorization)
	message := make(map[string]interface{})

	message["content"] = animaCrypto.Hash(jsonResult)

	sigRequest := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    models.PROTOCOL_NAME,
			Version: models.PROTOCOL_VERSION,
			ChainId: math.NewHexOrDecimal256(1),
		},
		PrimaryType: "Main",
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "version",
					Type: "string",
				},
			},
			"Main": []apitypes.Type{
				{
					Name: "content",
					Type: "string",
				},
			},
		},
		Message: message,
	}

	c, err := json.Marshal(sigRequest)
	if err != nil {
		return
	}

	verifierPrivateKey, _ := crypto.HexToECDSA(PRIVATE_VERIFIER_SIGNING_KEY)
	digest, err := evm.GetEIP712Message(c)
	if err != nil {
		return
	}

	verifierSignature, err := crypto.Sign(digest, verifierPrivateKey)
	if err != nil {
		return
	}

	fmt.Println(string(jsonResult))
	fmt.Println("0x" + hex.EncodeToString(verifierSignature))
}
