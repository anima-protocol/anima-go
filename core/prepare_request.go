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

		livenessBytes, err := json.Marshal(request.Liveness.Document)
		if err != nil {
			return err
		}

		livenessContentBytes := new(bytes.Buffer)
		err = json.Compact(livenessContentBytes, livenessBytes)
		if err != nil {
			return err
		}

		request.Document.Document.Liveness = &protocol.IssLiveness{
			Specs: request.Liveness.Document.Specs,
			Id:    fmt.Sprintf("anima:document:%s", crypto.Hash(livenessContentBytes.Bytes())),
		}
	}
	return nil
}
