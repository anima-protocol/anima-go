package signature

import (
	"context"
	"encoding/json"
	"fmt"
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

	starknetClient := client.NewStarknetClient(sig.ChainId)

	valid := starknetClient.IsValidSignature(context.Background(), publicAddress, messageHash, sig.Signature.R, sig.Signature.S)

	if !valid {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
