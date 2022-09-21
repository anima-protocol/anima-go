package core

import (
	elrond "github.com/anima-protocol/anima-go/chains/elrond"
	evm "github.com/anima-protocol/anima-go/chains/evm"
	"github.com/anima-protocol/anima-go/models"
)

var ExtractIssuingAuthorization = map[string]func([]byte, string) (*models.IssuingAuthorization, error){
	"anima:specs:issuing/authorization/eip712@1.0.0": evm.GetIssuingAuthorizationEIP712,
	"anima:specs:issuing/authorization/erd712@1.0.0": elrond.GetIssuingAuthorizationERD712,
}
