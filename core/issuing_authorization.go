package core

import (
	"encoding/base64"

	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func GetIssuingAuthorization(request *protocol.IssueRequest) (*models.IssuingAuthorization, error) {
	specs := request.Document.Authorization.Specs
	encodedContent := request.Document.Authorization.Content
	signature := request.Document.Authorization.Signature

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
