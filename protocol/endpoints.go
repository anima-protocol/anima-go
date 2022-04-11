package protocol

import (
	context "context"
	"errors"

	"github.com/anima-protocol/anima-go/chains/evm"
	"github.com/anima-protocol/anima-go/models"
	"github.com/anima-protocol/anima-go/utils"
	"google.golang.org/grpc/metadata"
)

func Issue(anima *models.Protocol, req *IssueRequest) error {
	config := &Config{Secure: anima.Secure}
	err := Init(config, anima)
	if err != nil {
		return err
	}

	if !utils.InArray(anima.Chain, []string{"ETH"}) {
		return errors.New("unsupported chain")
	}

	signature, err := evm.SignProtocolRequest(anima, req, anima.SigningFunc)
	if err != nil {
		return err
	}

	header := metadata.New(map[string]string{"signature": signature, "chain": anima.Chain})
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	if _, err = client.Issue(ctx, req); err != nil {
		return err
	}
	return nil
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
