package models

const (
	/* PROTOCOL */
	PROTOCOL_NAME    = "anima"
	PROTOCOL_VERSION = "1.0"

	/* NETWORK */
	MAINNET = "protocol.anima.io:443"
	TESTNET = "protocol-tesnet.anima.io:443"

	/* CHAIN */
	CHAIN_ETH    = "ETH"
	CHAIN_ETH_ID = 1
)

var AVAILABLE_CHAIN = []string{CHAIN_ETH}
