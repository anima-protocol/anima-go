package elrond

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func GetERD712Message(data []byte) ([]byte, error) {
	fmt.Printf("---> ERD_712 = %v\n", string(data))
	signerData := apitypes.TypedData{}
	err := json.Unmarshal(data, &signerData)
	if err != nil {
		fmt.Printf("-> err_1 %v\n", err)
		return []byte{}, err
	}

	rawData := []byte(fmt.Sprintf("\x17Elrond Signed Message:\n%d%s", len(data), data))
	challengeHash := crypto.Keccak256Hash(rawData)
	return challengeHash[:], nil
}
