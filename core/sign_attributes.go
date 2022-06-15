package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/anima-protocol/anima-go/chains/evm"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func SignIssuing(anima *models.Protocol, issuer *protocol.AnimaIssuer, request *protocol.IssueRequest, signingFunc func([]byte) (string, error)) (*protocol.IssueRequest, error) {
	issuingAuthorization, err := GetIssuingAuthorization(request)
	if err != nil {
		return nil, err
	}

	// Sign Proof
	proofContent, err := base64.StdEncoding.DecodeString(request.Proof.Content)
	if err != nil {
		return nil, err
	}

	proofContentBytes := new(bytes.Buffer)
	err = json.Compact(proofContentBytes, proofContent)
	if err != nil {
		return nil, err
	}

	switch anima.Chain {
	case models.CHAIN_ETH:
		proofSignature, err := evm.SignProof(anima, proofContent, signingFunc)
		if err != nil {
			return nil, err
		}

		request.Proof.Signature = "0x" + proofSignature
	}

	proofId := fmt.Sprintf("anima:proof:%s", crypto.Hash(proofContentBytes.Bytes()))

	owner := &protocol.AnimaOwner{
		Id:            issuingAuthorization.Owner.ID,
		PublicAddress: issuingAuthorization.Owner.PublicAddress,
		Chain:         issuingAuthorization.Owner.Chain,
		Wallet:        issuingAuthorization.Owner.Wallet,
	}

	request.Document.Owner = owner

	issuedAt := time.Now().Unix()
	// Sign Attributes
	for name := range request.Attributes {
		request.Attributes[name].Content = &protocol.IssDocumentAttributeContent{
			Value:         request.Document.Attributes[name].Content.Value,
			Type:          request.Document.Attributes[name].Content.Type,
			Name:          request.Document.Attributes[name].Content.Name,
			Format:        request.Document.Attributes[name].Content.Format,
			Authorization: request.Document.Authorization,
			Owner:         owner,
		}

		contentHash := crypto.HashStr(request.Document.Attributes[name].Content.Value)
		if request.Attributes[name].Content.Type == "file" {
			contentHash = request.Document.Attributes[name].Content.Value
		}

		attrBytes, err := json.Marshal(request.Attributes[name].Content)
		if err != nil {
			return nil, err
		}

		attrContentBytes := new(bytes.Buffer)
		err = json.Compact(attrContentBytes, attrBytes)
		if err != nil {
			return nil, err
		}

		request.Attributes[name].Credential.Content = &protocol.IssAttributeCredentialContent{
			IssuedAt:  issuedAt,
			ExpiresAt: request.Document.ExpiresAt,
			Owner:     owner,
			Issuer:    issuer,
			Attribute: &protocol.IssAttributeCredentialContentAttribute{
				Specs: "anima:specs:attribute@1.0.0",
				Id:    fmt.Sprintf("anima:attribute:%s", crypto.Hash(attrContentBytes.Bytes())),
				Hash:  contentHash,
				Name:  name,
			},
			Proof: &protocol.IssAttributeCredentialContentProof{
				Specs: request.Proof.Specs,
				Id:    proofId,
			},
			Authorization: &protocol.IssAttributeCredentialContentAuthorization{
				Specs:     request.Document.Authorization.Specs,
				Content:   request.Document.Authorization.Content,
				Signature: request.Document.Authorization.Signature,
			},
		}

		request.Document.Attributes[name].Credential = &protocol.IssDocumentAttributeCredential{
			Specs: "anima:specs:credential@1.0.0",
			Id:    fmt.Sprintf("anima:credential:%s", crypto.Hash(attrContentBytes.Bytes())),
		}
	}

	for name := range request.Attributes {
		documentBytes, err := json.Marshal(request.Document)
		if err != nil {
			return nil, err
		}

		documentContentBytes := new(bytes.Buffer)
		err = json.Compact(documentContentBytes, documentBytes)
		if err != nil {
			return nil, err
		}

		request.Attributes[name].Credential.Content.Document = &protocol.IssAttributeCredentialContentDocument{
			Specs: request.Document.Specs,
			Id:    fmt.Sprintf("anima:document:%s", crypto.Hash(documentContentBytes.Bytes())),
		}

		switch anima.Chain {
		case models.CHAIN_ETH:
			signature, err := evm.SignCredential(anima, request.Attributes[name].Credential.Content, signingFunc)
			if err != nil {
				return nil, err
			}

			request.Attributes[name].Credential.Signature = "0x" + signature
		}
	}

	return request, nil
}
