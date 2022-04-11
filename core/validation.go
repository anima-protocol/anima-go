package core

import (
	"github.com/anima-protocol/anima-go/chains/evm"
	"github.com/anima-protocol/anima-go/models"
)

var ExtractIssuingAuthorization = map[string]func([]byte, string) (*models.IssuingAuthorization, error){
	"anima:specs:issuing/authorization/eip712@1.0.0": evm.GetIssuingAuthorizationEIP712,
}
