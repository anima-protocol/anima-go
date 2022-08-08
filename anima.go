package anima

import (
	"github.com/anima-protocol/anima-go/core"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
	"github.com/anima-protocol/anima-go/validators"
)

// Issue - Issue new credential to Anima Protocol
func Issue(anima *models.Protocol, issuer *protocol.AnimaIssuer, requestDocument *protocol.IssueDocumentRequest, requestLiveness *protocol.IssueDocumentRequest) error {
	if err := validators.ValidateProtocol(anima); err != nil {
		return err
	}

	liveness, err := core.SignIssuing(anima, issuer, requestLiveness, anima.SigningFunc)
	if err != nil {
		return err
	}

	issueRequest := &protocol.IssueRequest{
		Liveness: liveness,
		Document: nil,
	}

	if requestDocument != nil {
		err = core.AddLivenessToIssuingDocument(liveness, requestDocument)
		if err != nil {
			return err
		}

		document, err := core.SignIssuing(anima, issuer, requestDocument, anima.SigningFunc)
		if err != nil {
			return err
		}

		issueRequest.Document = document
	}

	return protocol.Issue(anima, issueRequest)
}

// Verify - Verify Sharing Request from Anima Protocol
func Verify(anima *models.Protocol, request *protocol.VerifyRequest) (*protocol.VerifyResponse, error) {
	if err := validators.ValidateProtocol(anima); err != nil {
		return &protocol.VerifyResponse{}, err
	}

	return protocol.Verify(anima, request)
}

// RegisterVerifier - Register Verifier on Anima Protocol
func RegisterVerifier(anima *models.Protocol, request *protocol.RegisterVerifierRequest) (*protocol.RegisterVerifierResponse, error) {
	if err := validators.ValidateProtocol(anima); err != nil {
		return &protocol.RegisterVerifierResponse{}, err
	}

	return protocol.RegisterVerifier(anima, request)
}
