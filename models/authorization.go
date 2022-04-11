package models

type IssuingAuthorization struct {
	Specs       string            `json:"specs"`
	RequestedAt uint64            `json:"requested_at"`
	Fields      map[string]string `json:"fields"`
	Attributes  map[string]bool   `json:"attributes"`
	Owner       AnimaOwner        `json:"owner"`
	Issuer      AnimaIssuer       `json:"issuer"`
}
