package signature

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/NethermindEth/starknet.go/curve"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/anima-protocol/anima-go/chains/starknet/client"
	"github.com/anima-protocol/anima-go/chains/starknet/starknetTypedData"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
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

	messageHash, err := typedData.GetMessageHash(utils.HexToBN(publicAddress), typedDataMessage, curve.Curve)
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

	starknetClient := client.NewStarknetClient(sig.ChainId)

	valid, err := starknetClient.IsValidSignature(ctx, publicAddress, messageHash, finalSignature)
	if err != nil {
		return err
	}

	if !valid {
		buggedTypedDataMessage := starknetTypedData.CreateBuggedStarknetAuthorizationTypedDataMessage(data)

		buggedMessageHash, err := typedData.GetMessageHash(utils.HexToBN(publicAddress), buggedTypedDataMessage, curve.Curve)
		if err != nil {
			return err
		}

		validBugged, _ := starknetClient.IsValidSignature(ctx, publicAddress, buggedMessageHash, finalSignature)

		if !validBugged {
			return fmt.Errorf("invalid signature")
		}
	}

	return nil
}
