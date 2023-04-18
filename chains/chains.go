package chains

const (
	ETH        = "ETH"
	EVM        = "EVM"
	ELROND     = "ELROND"
	MULTIVERSX = "MULTIVERSX"
	COSMOS     = "COSMOS"
	STARKNET   = "STARKNET"
)

var SUPPORTED = map[string]bool{
	ETH:        true,
	EVM:        true,
	ELROND:     true,
	MULTIVERSX: true,
	COSMOS:     true,
	STARKNET:   true,
}
