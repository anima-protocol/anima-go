package elrond

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func GetERDPersonalSignMessage(data []byte) ([]byte, error) {
	fmt.Printf("---> ERD = %v\n", string(data))
	rawData := []byte(fmt.Sprintf("\x17Elrond Signed Message:\n%d%s", len(data), data))
	challengeHash := crypto.Keccak256Hash(rawData)
	return challengeHash[:], nil
}
