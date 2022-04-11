package models

type Protocol struct {
	Network     string                       `json:"network"`
	Chain       string                       `json:"chain"`
	SigningFunc func([]byte) (string, error) `json:"signing_func"`
	Secure      bool                         `json:"secure"`
}

type AnimaOwner struct {
	ID                  string `json:"id"`
	PublicAddress       string `json:"public_address"`
	Chain               string `json:"chain"`
	Wallet              string `json:"wallet,omitempty"`
	PublicKeyEncryption string `json:"public_key_encryption,omitempty"`
}

type AnimaIssuer struct {
	ID            string `json:"id"`
	PublicAddress string `json:"public_address"`
	Chain         string `json:"chain"`
}

type AnimaVerifier struct {
	ID            string `json:"id"`
	PublicAddress string `json:"public_address"`
	Chain         string `json:"chain"`
}

type AnimaProtocol struct {
	ID            string `json:"id"`
	PublicAddress string `json:"public_address"`
	Chain         string `json:"chain"`
}
