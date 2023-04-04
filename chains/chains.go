package chains

const (
	ETH    = "ETH"
	EVM    = "EVM"
	ELROND = "ELROND"
	COSMOS = "COSMOS"
)

var SUPPORTED = map[string]bool{
	ETH:    true,
	EVM:    true,
	ELROND: true,
	COSMOS: true,
}
