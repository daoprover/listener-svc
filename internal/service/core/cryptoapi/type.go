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
	RecipientAddress     string `json:"recipientAddress"`
	SenderAddress        string `json:"senderAddress"`
	TokenDecimals        int    `json:"tokenDecimals"`
	TokenID              string `json:"tokenId"`
	TokenName            string `json:"tokenName"`
	TokenSymbol          string `json:"tokenSymbol"`
	TokenType            string `json:"tokenType"`
	TokensAmount         string `json:"tokensAmount"`
	TransactionHash      string `json:"transactionHash"`
	TransactionTimestamp int    `json:"transactionTimestamp"`
}
