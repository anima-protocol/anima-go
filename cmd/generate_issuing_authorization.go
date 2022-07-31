package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/anima-protocol/anima-go/chains/evm"
	crypto2 "github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/utils"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"time"
)

type IssuingAuthorization struct {
	Specs     string `json:"specs"`
	Content   string `json:"content"`
	Signature string `json:"signature"`
}

const (
	PASSORT         string = "anima:specs:document/passport@1.0.0"
	ID                     = "anima:specs:document/national_id@1.0.0"
	DRIVER_LICENSE         = "anima:specs:document/driver_license@1.0.0"
	RESIDENT_PERMIT        = "anima:specs:document/resident_permit@1.0.0"
	LIVENESS               = "anima:specs:document/liveness@1.0.0"
	FACE                   = "anima:specs:face@1.0.0"
)

var documentSpecsName = []string{
	PASSORT,
	ID,
	DRIVER_LICENSE,
	RESIDENT_PERMIT,
	LIVENESS,
	FACE,
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load '.env' file, terminating.")
	}
}

func getEnv(evar string) string {
	val := os.Getenv(evar)
	if val == "" {
		log.Fatalf("[MISSING EVAR] %v\n", evar)
		os.Exit(1)
	}
	return val
}

func generateFields(specs string) map[string]interface{} {
	switch specs {
	case PASSORT:
		return map[string]interface{}{
			"passport_page":          "6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d",
			"original_passport_page": "41a4547b3caa10e7c81789a02ccf8f28f861ad5e58d29ea6ac4b004e179c05c2",
		}
	case ID, DRIVER_LICENSE, RESIDENT_PERMIT:
		return map[string]interface{}{
			"document_front":          "6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d",
			"original_document_front": "fb2e6923ee3290f90adb76851a6b515b1ac50c976285eb3b68bb2b091457cd7f",
			"document_back":           "4bf5122f344554c53bde2ebb8cd2b7e3d1600ad631c385a5d7cce23c7785459a",
			"original_document_back":  "ac521e88646e7cdf7685f5d4e5447bd982eade98429e46111064b0462d84cb95",
		}
	case LIVENESS:
		return map[string]interface{}{
			"face_vector": "17f71ed4d556a3ba04707ed8f727159739e367b45589e48fcd8ea2756a1ed4b1",
			"audit_trail": crypto2.HashStr("audit_trail"),
		}
	case FACE:
		return map[string]interface{}{
			"face_vector": "17f71ed4d556a3ba04707ed8f727159739e367b45589e48fcd8ea2756a1ed4b1",
		}
	}

	return map[string]interface{}{}
}

func generateAttributes(specs string) map[string]bool {
	switch specs {
	case PASSORT:
		return map[string]bool{
			"firstname":              true,
			"lastname":               true,
			"birth_date":             true,
			"nationality":            true,
			"document_country":       true,
			"document_number":        true,
			"document_expiration":    true,
			"passport_page":          true,
			"original_passport_page": true,
		}
	case ID, RESIDENT_PERMIT:
		return map[string]bool{
			"firstname":               true,
			"lastname":                true,
			"birth_date":              true,
			"nationality":             true,
			"document_country":        true,
			"document_number":         true,
			"document_expiration":     true,
			"document_front":          true,
			"original_document_front": true,
			"document_back":           true,
			"original_document_back":  true,
		}
	case DRIVER_LICENSE:
		return map[string]bool{
			"firstname":               true,
			"lastname":                true,
			"birth_date":              true,
			"document_country":        true,
			"document_number":         true,
			"document_expiration":     true,
			"document_front":          true,
			"original_document_front": true,
			"document_back":           true,
			"original_document_back":  true,
		}
	case LIVENESS:
		return map[string]bool{
			"face_vector": true,
			"audit_trail": true,
		}
	case FACE:
		return map[string]bool{
			"face_vector": true,
		}
	}
	return map[string]bool{}
}

func generateFieldsTypes(specs string) []apitypes.Type {
	switch specs {
	case PASSORT:
		return []apitypes.Type{
			{
				Name: "passport_page",
				Type: "string",
			},
			{
				Name: "original_passport_page",
				Type: "string",
			},
		}
	case ID, RESIDENT_PERMIT, DRIVER_LICENSE:
		return []apitypes.Type{
			{
				Name: "document_front",
				Type: "string",
			},
			{
				Name: "original_document_front",
				Type: "string",
			},
			{
				Name: "document_back",
				Type: "string",
			},
			{
				Name: "original_document_back",
				Type: "string",
			},
		}
	case LIVENESS:
		return []apitypes.Type{
			{
				Name: "face_vector",
				Type: "string",
			},
			{
				Name: "audit_trail",
				Type: "string",
			},
		}
	case FACE:
		return []apitypes.Type{
			{
				Name: "face_vector",
				Type: "string",
			},
		}
	}
	return []apitypes.Type{}
}

func generateAttributesTypes(specs string) []apitypes.Type {
	switch specs {
	case PASSORT:
		return []apitypes.Type{
			{
				Name: "firstname",
				Type: "bool",
			},
			{
				Name: "lastname",
				Type: "bool",
			},
			{
				Name: "birth_date",
				Type: "bool",
			},
			{
				Name: "nationality",
				Type: "bool",
			},
			{
				Name: "document_country",
				Type: "bool",
			},
			{
				Name: "document_number",
				Type: "bool",
			},
			{
				Name: "document_expiration",
				Type: "bool",
			},
			{
				Name: "passport_page",
				Type: "bool",
			},
			{
				Name: "original_passport_page",
				Type: "bool",
			},
		}
	case ID, RESIDENT_PERMIT:
		return []apitypes.Type{
			{
				Name: "firstname",
				Type: "bool",
			},
			{
				Name: "lastname",
				Type: "bool",
			},
			{
				Name: "birth_date",
				Type: "bool",
			},
			{
				Name: "nationality",
				Type: "bool",
			},
			{
				Name: "document_country",
				Type: "bool",
			},
			{
				Name: "document_number",
				Type: "bool",
			},
			{
				Name: "document_expiration",
				Type: "bool",
			},
			{
				Name: "document_front",
				Type: "bool",
			},
			{
				Name: "original_document_front",
				Type: "bool",
			},
			{
				Name: "document_back",
				Type: "bool",
			},
			{
				Name: "original_document_back",
				Type: "bool",
			},
		}
	case DRIVER_LICENSE:
		return []apitypes.Type{
			{
				Name: "firstname",
				Type: "bool",
			},
			{
				Name: "lastname",
				Type: "bool",
			},
			{
				Name: "birth_date",
				Type: "bool",
			},
			{
				Name: "document_country",
				Type: "bool",
			},
			{
				Name: "document_number",
				Type: "bool",
			},
			{
				Name: "document_expiration",
				Type: "bool",
			},
			{
				Name: "document_front",
				Type: "bool",
			},
			{
				Name: "original_document_front",
				Type: "bool",
			},
			{
				Name: "document_back",
				Type: "bool",
			},
			{
				Name: "original_document_back",
				Type: "bool",
			},
		}
	case LIVENESS:
		return []apitypes.Type{
			{
				Name: "face_vector",
				Type: "bool",
			},
			{
				Name: "audit_trail",
				Type: "bool",
			},
		}
	case FACE:
		return []apitypes.Type{
			{
				Name: "face_vector",
				Type: "bool",
			},
		}
	}
	return []apitypes.Type{}
}

func generateTypes(specs string) apitypes.Types {
	baseType := apitypes.Types{
		"Main": []apitypes.Type{
			{
				Name: "issuer",
				Type: "Issuer",
			},
			{
				Name: "owner",
				Type: "Owner",
			},
			{
				Name: "specs",
				Type: "string",
			},
			{
				Name: "requested_at",
				Type: "uint64",
			},
			{
				Name: "fields",
				Type: "Fields",
			},
			{
				Name: "attributes",
				Type: "Attributes",
			},
		},
		"Fields":     generateFieldsTypes(specs),
		"Attributes": generateAttributesTypes(specs),
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
			{
				Name: "public_key_encryption",
				Type: "string",
			},
		},
		"Issuer": []apitypes.Type{
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
	}
	return baseType
}

func main() {
	loadEnv()
	PRIVATE_OWNER_SIGNING_KEY := getEnv("PRIVATE_OWNER_SIGNING_KEY")
	PUBLIC_OWNER_ADDRESS := getEnv("PUBLIC_OWNER_ADDRESS")
	PUBLIC_OWNER_ENCRYPTION_KEY := getEnv("PUBLIC_OWNER_ENCRYPTION_KEY")

	specsPtr := flag.String("specs", documentSpecsName[rand.Intn(len(documentSpecsName))], "specs type to generate")
	validPtr := flag.Bool("valid", true, "is signature valid or not")

	flag.Parse()
	if !utils.InArray(*specsPtr, documentSpecsName) {
		log.Fatal("Invalid specs value")
		os.Exit(1)
	}

	authorization := &evm.IssuingAuthorizationEIP712{
		Domain: apitypes.TypedDataDomain{
			Name:    models.PROTOCOL_NAME,
			Version: models.PROTOCOL_VERSION,
			ChainId: math.NewHexOrDecimal256(models.CHAIN_ETH_ID),
		},
		Message: models.IssuingAuthorization{
			Specs:       *specsPtr,
			RequestedAt: uint64(time.Now().Unix()),
			Fields:      generateFields(*specsPtr),
			Attributes:  generateAttributes(*specsPtr),
			Owner: models.AnimaOwner{
				ID:                  fmt.Sprintf("anima:owner:%s", PUBLIC_OWNER_ADDRESS),
				PublicAddress:       PUBLIC_OWNER_ADDRESS,
				Chain:               "ETH",
				Wallet:              "METAMASK",
				PublicKeyEncryption: PUBLIC_OWNER_ENCRYPTION_KEY,
			},
			Issuer: models.AnimaIssuer{
				ID:            "anima:issuer:synaps@1.0.0",
				PublicAddress: "0x6bf88580aF74117322bB4bA54Ac497A66B77B4B6",
				Chain:         "ETH",
			},
		},
		Types: generateTypes(*specsPtr),
	}

	challenge, err := json.Marshal(authorization)
	if err != nil {
		log.Fatal("Error while transforming authorization in json")
		os.Exit(1)
	}

	challengeHash, err := evm.GetEIP712Message(challenge)

	fmt.Printf("%v\n", err)
	privateKey, _ := crypto.HexToECDSA(PRIVATE_OWNER_SIGNING_KEY)
	signature, err := crypto.Sign(challengeHash, privateKey)
	if err != nil {
		fmt.Printf("Error while signing: %v\n", err)
		os.Exit(1)
	}
	if *validPtr == false {
		signature[2] = 23
		signature[3] = 23
		signature[4] = 23
	}
	hexSignature := "0x" + hex.EncodeToString(signature)

	base64Challenge := base64.StdEncoding.EncodeToString(challenge)

	result := IssuingAuthorization{
		Specs:     "anima:specs:issuing/authorization/eip712@1.0.0",
		Content:   base64Challenge,
		Signature: hexSignature,
	}

	jsonResult, _ := json.Marshal(result)

	fmt.Print(string(jsonResult))
}
