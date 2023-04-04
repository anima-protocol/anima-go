package models

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type StdFee struct {
	Amount   []Coin  `json:"amount"`
	Gas      string  `json:"gas"`
	Payer    *string `json:"payer,omitempty"`
	Granter  *string `json:"granter,omitempty"`
	FeePayer *string `json:"feePayer,omitempty"`
}

type MsgSignData struct {
	Data   string `json:"data"`
	Signer string `json:"signer"`
}

type Msg struct {
	Type  string      `json:"type"`
	Value MsgSignData `json:"value"`
}

type StdSignDoc struct {
	AccountNumber string  `json:"account_number"`
	ChainId       string  `json:"chain_id"`
	Fee           StdFee  `json:"fee"`
	Memo          string  `json:"memo"`
	Msgs          []Msg   `json:"msgs"`
	Sequence      string  `json:"sequence"`
	TimeoutHeight *string `json:"timeout_height,omitempty"`
}
