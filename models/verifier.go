package models

type RegisterVerifierRequest struct {
	Id            string `json:"id"`
	PublicAddress string `json:"public_address"`
	Chain         string `json:"chain"`
}
