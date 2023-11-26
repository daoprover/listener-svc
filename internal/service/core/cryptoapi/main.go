package cryptoapi

import (
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type CryptoapiHandler interface {
}

type CryptoAPI struct {
	api    string
	apiKey string
}

func NewCryptoAPI(api string, apiKey string) CryptoAPI {
	return CryptoAPI{
		api:    api,
		apiKey: apiKey,
	}
}

func (c *CryptoAPI) setHeader(r *http.Request) *http.Request {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-API-Key", c.apiKey)

	return r
}

func (c *CryptoAPI) GetTransactionByAddress(address string) ([]string, error) {

	return nil, nil
}

// https://rest.cryptoapis.ioblockchain-data/ethereum
func (c *CryptoAPI) GetInternalTransactionByAddress(address string, timestampFrom, timestampTo string) (interface{}, error) {
	url := fmt.Sprintf("%s/goerli/addresses/%s/internal-by-time-range?fromTimestamp=%s&toTimestamp=%s", c.api, address, timestampFrom, timestampTo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to init request")
	}

	req = c.setHeader(req)
	rawRes, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send request")
	}
	response := TransactionsResponse{}
	err = json.NewDecoder(rawRes.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode data")
	}
	return nil, nil
}

func (c *CryptoAPI) GetTokensHolders() error {

}
