package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/anima-protocol/anima-go/crypto"
	"github.com/anima-protocol/anima-go/protocol"
)

func PrepareIssuingRequest(liveness *protocol.IssueDocumentRequest, doc *protocol.IssueDocumentRequest) error {
	if doc != nil {
		if liveness == nil {
			return fmt.Errorf("try to issue document without liveness document")
		}

		livenessBytes, err := json.Marshal(liveness.Document)
		if err != nil {
			return err
		}

		livenessContentBytes := new(bytes.Buffer)
		err = json.Compact(livenessContentBytes, livenessBytes)
		if err != nil {
			return err
		}

		doc.Document.Liveness = &protocol.IssLiveness{
			Specs: liveness.Document.Specs,
			Id:    fmt.Sprintf("anima:document:%s", crypto.Hash(livenessContentBytes.Bytes())),
		}
	}
	return nil
}
