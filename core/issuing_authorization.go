package core

import (
	"encoding/base64"

	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func GetIssuingAuthorization(document *protocol.IssDocument) (*models.IssuingAuthorization, error) {
	specs := document.Authorization.Specs
	encodedContent := document.Authorization.Content
	signature := document.Authorization.Signature

	content, err := base64.StdEncoding.DecodeString(encodedContent)
	if err != nil {
		return nil, err
	}

	issuingAuthorization, rErr := ExtractIssuingAuthorization[specs](content, signature)
	if rErr != nil {
		return nil, rErr
	}

	return issuingAuthorization, nil
}
