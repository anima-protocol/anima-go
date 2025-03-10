package signature

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/anima-protocol/anima-go/chains/starknet/client"
	"github.com/anima-protocol/anima-go/chains/starknet/starknetTypedData"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"

	"github.com/NethermindEth/starknet.go/utils"
)

type StarknetSignature struct {
	ChainId       string    `json:"chainId"`
	Signature     Signature `json:"signature"`
	FullSignature []string  `json:"fullSignature"`
}

type Signature struct {
	R string `json:"r"`
	S string `json:"s"`
}

func VerifyPersonalSignature(publicAddress string, data []byte, userSignature string, rpcProviderUrl string) error {
	ctx := context.Background()
	// convert userSignature hex string to bytes
	sigHex, err := hexutil.Decode(userSignature)
	if err != nil {
		return err
	}

	var sig StarknetSignature
	err = json.Unmarshal(sigHex, &sig)
	if err != nil {
		return err
	}

	typedDataMessage, err := starknetTypedData.CreateStarknetAuthorizationTypedDataMessage(data)
	if err != nil {
		return err
	}

	typedData, err := starknetTypedData.CreateStarknetAuthorizationTypedDataDefinition(sig.ChainId, typedDataMessage)
	if err != nil {
		return err
	}

	messageHash, err := typedData.GetMessageHash(publicAddress)
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

	var finalSignature []string

	if sig.FullSignature == nil || len(sig.FullSignature) == 0 {
		finalSignature = []string{sig.Signature.R, sig.Signature.S}
	} else {
		finalSignature = sig.FullSignature
	}

	starknetClient := client.NewStarknetClient(rpcProviderUrl)

	valid, err := starknetClient.IsValidSignature(ctx, publicAddress, utils.FeltToBigInt(messageHash), finalSignature)
	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
