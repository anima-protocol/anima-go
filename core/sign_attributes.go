package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/anima-protocol/anima-go/chains/evm"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/protocol"
)

func SignIssuing(anima *models.Protocol, issuer *protocol.AnimaIssuer, request *protocol.IssueDocumentRequest, signingFunc func([]byte) (string, error)) (*protocol.IssueDocumentRequest, error) {
	issuingAuthorization, err := GetIssuingAuthorization(request.Document)
	if err != nil {
		return nil, err
	}

	if issuingAuthorization.PublicKeyEncryption != "" {
		request.Document.EncryptionKey = issuingAuthorization.PublicKeyEncryption
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
		proofSignature, err := evm.SignBytes(proofContentBytes.Bytes(), signingFunc)
		if err != nil {
			return nil, err
		}

		request.Proof.Signature = "0x" + proofSignature
	}

	proofId := fmt.Sprintf("anima:proof:%s", crypto.HashSHA256(proofContentBytes.Bytes()))

	owner := &protocol.AnimaOwner{
		Id:            issuingAuthorization.Owner.ID,
		PublicAddress: issuingAuthorization.Owner.PublicAddress,
		Chain:         issuingAuthorization.Owner.Chain,
	}

	request.Document.Owner = owner

	issuedAt := time.Now().Unix()
	// Sign Attributes
	for name := range request.Attributes {
		attrType := request.Document.Attributes[name].Content.Type

		request.Attributes[name].Content = &protocol.IssDocumentAttributeContent{
			Type:          attrType,
			Name:          request.Document.Attributes[name].Content.Name,
			Format:        request.Document.Attributes[name].Content.Format,
			Authorization: request.Document.Authorization,
			Owner:         owner,
		}
		contentHashes := []string{}
		contentHash := ""

		if strings.HasSuffix(attrType, "[]") {
			request.Attributes[name].Content.Values = request.Document.Attributes[name].Content.Values

			if strings.HasPrefix(attrType, "file") {
				contentHashes = request.Document.Attributes[name].Content.Values
			} else {
				for _, value := range request.Document.Attributes[name].Content.Values {
					contentHashes = append(contentHashes, crypto.HashSHA256Str(value))
				}
			}
		} else {
			request.Attributes[name].Content.Value = request.Document.Attributes[name].Content.Value
			if request.Attributes[name].Content.Type == "file" {
				contentHash = request.Document.Attributes[name].Content.Value
			} else {
				contentHash = crypto.HashSHA256Str(request.Document.Attributes[name].Content.Value)
			}
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
				Specs: "anima:schema:attribute@1.0.0",
				Id:    fmt.Sprintf("anima:attribute:%s", crypto.HashSHA256(attrContentBytes.Bytes())),
				Hash:  contentHash,
				Name:  name,
			},
			Proof: &protocol.IssAttributeCredentialContentProof{
				Specs: request.Proof.Specs,
				Id:    proofId,
			},
			Authorization: &protocol.IssAttributeCredentialContentAuthorization{
				Content:   request.Document.Authorization.Content,
				Signature: request.Document.Authorization.Signature,
			},
		}
		if strings.HasSuffix(attrType, "[]") {
			request.Attributes[name].Credential.Content.Attribute.Hashes = contentHashes
		} else {
			request.Attributes[name].Credential.Content.Attribute.Hash = contentHash
		}

		request.Document.Attributes[name].Credential = &protocol.IssDocumentAttributeCredential{
			Specs: "anima:schema:credential@1.0.0",
			Id:    fmt.Sprintf("anima:credential:%s", crypto.HashSHA256(attrContentBytes.Bytes())),
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
			Id:    fmt.Sprintf("anima:document:%s", crypto.HashSHA256(documentContentBytes.Bytes())),
		}

		switch anima.Chain {
		case models.CHAIN_ETH:
			signature, err := evm.SignInterfaceData(request.Attributes[name].Credential.Content, signingFunc)
			if err != nil {
				return nil, err
			}

			request.Attributes[name].Credential.Signature = "0x" + signature
		}
	}

	return request, nil
}
