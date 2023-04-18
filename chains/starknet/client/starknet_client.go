package client

import (
	"context"
	"fmt"
	"github.com/dontpanicdao/caigo/gateway"
	"github.com/dontpanicdao/caigo/types"
	"math/big"
)

type StarknetClient struct {
	provider *gateway.GatewayProvider
}

func NewStarknetClient(chainId string) *StarknetClient {
	return &StarknetClient{
		provider: gateway.NewProvider(gateway.WithChain(chainId)),
	}
}

func (c *StarknetClient) IsValidSignature(context context.Context, address string, messageHash *big.Int, r string, s string) bool {
	callResp, err := c.provider.Call(context, types.FunctionCall{
		ContractAddress:    types.HexToHash(address),
		EntryPointSelector: "isValidSignature",
		Calldata: []string{
			messageHash.String(),
			"2",
			r,
			s,
		},
	}, "")

	if err != nil {
		fmt.Errorf("error calling isValidSignature: %v\n", err)
		return false
	} else if callResp[0] == "0x1" {
		return true
	}
	return false
}
