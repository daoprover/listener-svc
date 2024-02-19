package cryptoapi

type TransactionsResponse struct {
	APIVersion string                   `json:"apiVersion"`
	RequestID  string                   `json:"requestId"`
	Context    string                   `json:"context"`
	Data       TransactionsResponseData `json:"data"`
}

type TransactionsResponseData struct {
	Limit  int               `json:"limit"`
	Offset int               `json:"offset"`
	Total  int               `json:"total"`
	Items  []TransactionItem `json:"items"`
}

type TransactionItem struct {
	ContractAddress      string `json:"contractAddress"`
	MinedInBlockHeight   int    `json:"minedInBlockHeight"`
	RecipientAddress     string `json:"recipient"`
	SenderAddress        string `json:"sender"`
	TokenDecimals        int    `json:"tokenDecimals"`
	TokenID              string `json:"tokenId"`
	TokenName            string `json:"tokenName"`
	TokenSymbol          string `json:"tokenSymbol"`
	TokenType            string `json:"tokenType"`
	TokensAmount         string `json:"amount"`
	TransactionHash      string `json:"transactionHash"`
	TransactionTimestamp int    `json:"transactionTimestamp"`
	Timestamp            uint   `json:"timestamp"`
}

type Holder struct {
	Address         string
	TokensAmount    float64
	TimeFrom        uint
	TimeTo          uint
	TokenAmountTime int
}
