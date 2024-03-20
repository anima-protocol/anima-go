package client

import (
	"context"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"math/big"
	"strings"

	"github.com/anima-protocol/anima-go/chains/starknet/errors"
)

type StarknetClient struct {
	provider *rpc.Provider
}

func NewStarknetClient(providerUrl string) *StarknetClient {
	provider, err := rpc.NewProvider(providerUrl)
	if err != nil {
		panic(err)
	}
	return &StarknetClient{
		provider: provider,
	}
}

func (c *StarknetClient) IsValidSignature(context context.Context, address string, messageHash *big.Int, fullSignature []string) (bool, error) {
	contractAddress, err := utils.HexToFelt(address)
	if err != nil {
		return false, err
	}

	var signatureCallData []*felt.Felt

	for _, sig := range fullSignature {
		signatureCallData = append(signatureCallData, utils.BigIntToFelt(utils.StrToBig(sig)))
	}

	tx := rpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt("isValidSignature"),
		Calldata: append([]*felt.Felt{
			utils.BigIntToFelt(messageHash),
			utils.Uint64ToFelt(uint64(len(fullSignature))),
		}, signatureCallData...),
	}

	callResp, err := c.provider.Call(context, tx, rpc.BlockID{Tag: "latest"})

	if err != nil {
		if strings.Contains(err.Error(), "StarknetErrorCode.UNINITIALIZED_CONTRACT") {
			return false, errors.Error_Not_Deployed
		}
		return false, nil
	} else if callResp[0].String() == "0x1" {
		return true, nil
	}
	return false, nil
}
