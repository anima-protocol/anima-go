package models

type IssAttribute struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Format string `json:"format"`
	Value  string `json:"value"`
}

type IssOwner struct {
	ID            string `json:"id"`
	PublicAddress string `json:"public_address"`
	Chain         string `json:"chain"`
}

type IssAuthorization struct {
	Specs     string `json:"specs"`
	Content   string `json:"content"`
	Signature string `json:"signature"`
}

type IssDocument struct {
	Specs                string                  `json:"specs"`
	IssuedAt             string                  `json:"issued_at"`
	ExpiresAt            string                  `json:"expires_at"`
	Attributes           map[string]IssAttribute `json:"attributes"`
	Owner                IssOwner                `json:"owner"`
	IssuingAuthorization IssAuthorization        `json:"issuing_authorization"`
}

type IssAttributeCredential struct {
	Value      interface{}   `json:"value"`
	Credential IssCredential `json:"credential"`
}

type IssCredentialSource struct {
	ID    string `json:"id"`
	Specs string `json:"specs"`
}

type IssCredential struct {
	Content   IssCredentialContent `json:"content"`
	Signature string               `json:"signature"`
}

type IssCredentialContent struct {
	IssuedAt  string              `json:"issued_at"`
	ExpiresAt string              `json:"expires_at"`
	Source    IssCredentialSource `json:"source"`
	Owner     IssOwner            `json:"owner"`
	Name      string              `json:"name"`
	Type      string              `json:"type"`
	Format    string              `json:"format"`
	Hash      string              `json:"hash"`
}

type IssProof struct {
	Specs     string `json:"specs"`
	Content   string `json:"content"`
	Signature string `json:"signature"`
}

type IssueRequest struct {
	Document   IssDocument                       `json:"document"`
	Attributes map[string]IssAttributeCredential `json:"attributes"`
	Proof      IssProof                          `json:"proof"`
}
