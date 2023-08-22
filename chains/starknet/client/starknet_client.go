package client

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/anima-protocol/anima-go/chains/starknet/errors"
	"github.com/dontpanicdao/caigo/gateway"
	"github.com/dontpanicdao/caigo/types"
)

type StarknetClient struct {
	provider *gateway.GatewayProvider
}

func NewStarknetClient(chainId string) *StarknetClient {
	return &StarknetClient{
		provider: gateway.NewProvider(gateway.WithChain(chainId)),
	}
}

func (c *StarknetClient) IsValidSignature(context context.Context, address string, messageHash *big.Int, fullSignature []string) (bool, error) {
	callResp, err := c.provider.Call(context, types.FunctionCall{
		ContractAddress:    types.HexToHash(address),
		EntryPointSelector: "isValidSignature",
		Calldata: append([]string{
			messageHash.String(),
			strconv.Itoa(len(fullSignature)),
		}, fullSignature...),
	}, "")

	if err != nil {
		if strings.Contains(err.Error(), "StarknetErrorCode.UNINITIALIZED_CONTRACT") {
			return false, errors.Error_Not_Deployed
		}
		fmt.Printf("MessageHash: %s\n", messageHash.String())
		return false, nil
	} else if callResp[0] == "0x1" {
		return true, nil
	}
	return false, nil
}
