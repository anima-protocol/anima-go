package chains

const (
	ETH      = "ETH"
	EVM      = "EVM"
	ELROND   = "ELROND"
	COSMOS   = "COSMOS"
	STARKNET = "STARKNET"
)

var SUPPORTED = map[string]bool{
	ETH:      true,
	EVM:      true,
	ELROND:   true,
	COSMOS:   true,
	STARKNET: true,
}
