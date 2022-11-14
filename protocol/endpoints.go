package protocol

import (
	context "context"
	"errors"

	"github.com/anima-protocol/anima-go/chains/evm"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/utils"
	"google.golang.org/grpc/metadata"
)

func Issue(anima *models.Protocol, req *IssueDocumentRequest) (*IssueDocumentResponse, error) {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return nil, err
	}

	if !utils.InArray(anima.Chain, []string{"ETH"}) {
		return nil, errors.New("unsupported chain")
	}

	signature, err := evm.SignProtocolRequest(anima, req, anima.SigningFunc)
	if err != nil {
		return nil, err
	}

	header := metadata.New(map[string]string{"signature": signature, "chain": anima.Chain})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	response, err := client.Issue(ctx, req)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func Verify(anima *models.Protocol, req *VerifyRequest) (*VerifyResponse, error) {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return &VerifyResponse{}, err
	}

	if !utils.InArray(anima.Chain, []string{"ETH"}) {
		return &VerifyResponse{}, errors.New("unsupported chain")
	}

	signature, err := evm.SignProtocolRequest(anima, req, anima.SigningFunc)
	if err != nil {
		return &VerifyResponse{}, err
	}

	header := metadata.New(map[string]string{"signature": signature, "chain": anima.Chain})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	res, err := client.Verify(ctx, req)
	if err != nil {
		return &VerifyResponse{}, err
	}

	return res, nil
}

func RegisterVerifier(anima *models.Protocol, req *RegisterVerifierRequest) (*RegisterVerifierResponse, error) {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return &RegisterVerifierResponse{}, err
	}

	if !utils.InArray(anima.Chain, []string{"ETH"}) {
		return &RegisterVerifierResponse{}, errors.New("unsupported chain")
	}

	signature, err := evm.SignProtocolRequest(anima, req, anima.SigningFunc)
	if err != nil {
		return &RegisterVerifierResponse{}, err
	}

	header := metadata.New(map[string]string{"signature": signature, "chain": anima.Chain})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	res, err := client.RegisterVerifier(ctx, req)
	if err != nil {
		return &RegisterVerifierResponse{}, err
	}

	return res, nil
}

func RequestAction(anima *models.Protocol, req *RequestActionRequest) (*RequestActionResponse, error) {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return &RequestActionResponse{}, err
	}

	res, err := client.RequestAction(context.Background(), req)
	if err != nil {
		return &RequestActionResponse{}, err
	}

	return res, nil
}

func GrantTrustee(anima *models.Protocol, req *GrantTrusteeRequest) (*Empty, error) {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return &Empty{}, err
	}

	res, err := client.GrantTrustee(context.Background(), req)
	if err != nil {
		return &Empty{}, err
	}

	return res, nil
}

func RevokeTrustee(anima *models.Protocol, req *RevokeTrusteeRequest) (*Empty, error) {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return &Empty{}, err
	}

	res, err := client.RevokeTrustee(context.Background(), req)
	if err != nil {
		return &Empty{}, err
	}

	return res, nil
}

func ListTrustees(anima *models.Protocol, req *ListTrusteesRequest) (*ListTrusteesResponse, error) {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return &ListTrusteesResponse{}, err
	}

	res, err := client.ListTrustees(context.Background(), req)
	if err != nil {
		return &ListTrusteesResponse{}, err
	}

	return res, nil
}

func DeleteAnima(anima *models.Protocol, req *DeleteAnimaRequest) (*Empty, error) {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return &Empty{}, err
	}

	res, err := client.DeleteAnima(context.Background(), req)
	if err != nil {
		return &Empty{}, err
	}

	return res, nil
}
