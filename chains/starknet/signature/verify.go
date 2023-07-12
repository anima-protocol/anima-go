package signature

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/anima-protocol/anima-go/chains/starknet/client"
	"github.com/anima-protocol/anima-go/chains/starknet/starknetTypedData"
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type StarknetSignature struct {
	ChainId   string    `json:"chainId"`
	Signature Signature `json:"signature"`
}

type Signature struct {
	R string `json:"r"`
	S string `json:"s"`
}

func VerifyPersonalSignature(publicAddress string, data []byte, userSignature string) error {
	ctx := context.Background()
	//convert userSignature hex string to bytes
	sigHex, err := hexutil.Decode(userSignature)
	if err != nil {
		return err
	}

	var sig StarknetSignature
	err = json.Unmarshal(sigHex, &sig)
	if err != nil {
		return err
	}

	typedData, err := starknetTypedData.CreateStarknetAuthorizationTypedDataDefinition(sig.ChainId)
	if err != nil {
		return err
	}

	typedDataMessage := starknetTypedData.CreateStarknetAuthorizationTypedDataMessage(data)

	messageHash, err := typedData.GetMessageHash(types.HexToBN(publicAddress), typedDataMessage, caigo.Curve)
	if err != nil {
		return err
	}

	if strings.HasPrefix(sig.ChainId, "0x") {
		chainIdWithoutPrefix := sig.ChainId[2:]
		bs, err := hex.DecodeString(chainIdWithoutPrefix)
		if err != nil {
			return err
		}
		sig.ChainId = string(bs)
	}

	starknetClient := client.NewStarknetClient(sig.ChainId)

	valid, err := starknetClient.IsValidSignature(ctx, publicAddress, messageHash, sig.Signature.R, sig.Signature.S)
	if err != nil {
		return err
	}

	if !valid {
		buggedTypedDataMessage := starknetTypedData.CreateBuggedStarknetAuthorizationTypedDataMessage(data)
		fmt.Printf("Bugged message: %v\n", buggedTypedDataMessage)
		buggedMessageHash, err := typedData.GetMessageHash(types.HexToBN(publicAddress), buggedTypedDataMessage, caigo.Curve)
		if err != nil {
			return err
		}
		fmt.Printf("Bugged message hash: %s\n", buggedMessageHash)

		validBugged, _ := starknetClient.IsValidSignature(ctx, publicAddress, buggedMessageHash, sig.Signature.R, sig.Signature.S)

		if !validBugged {
			return fmt.Errorf("invalid signature")
		}
	}

	return nil
}
