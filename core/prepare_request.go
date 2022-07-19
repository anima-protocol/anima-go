package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/protocol"
)

func PrepareIssuingRequest(request *protocol.IssueRequest) error {
	if request.Document != nil {
		if request.Liveness == nil {
			return fmt.Errorf("try to issue document without liveness document")
		}

		docBytes, err := json.Marshal(request.Document.Document)
		if err != nil {
			return err
		}

		docContentBytes := new(bytes.Buffer)
		err = json.Compact(docContentBytes, docBytes)
		if err != nil {
			return err
		}

		request.Document.Document.Liveness = &protocol.IssLiveness{
			Specs: request.Liveness.Document.Specs,
			Id:    fmt.Sprintf("anima:document:%s", crypto.Hash(docContentBytes.Bytes())),
		}
	}
	return nil
}
