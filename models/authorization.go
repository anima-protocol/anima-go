package models

type IssuingAuthorization struct {
	Specs               string                 `json:"specs"`
	PublicKeyEncryption string                 `json:"public_key_encryption,omitempty"`
	Nonce               string                 `json:"nonce,omitempty"`
	RequestedAt         int64                  `json:"requested_at"`
	Fields              map[string]interface{} `json:"fields"`
	Attributes          map[string]bool        `json:"attributes"`
	Owner               AnimaOwner             `json:"owner"`
	Issuer              AnimaIssuer            `json:"issuer"`
}
