package cryptoapi

import (
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"
)

type CryptoapiHandler interface {
	GetInternalTransactionByAddress(address string, timestampFrom, timestampTo string) (*TransactionsResponse, error)
	GetTokensHolders(response TransactionsResponse) map[string]Holder
	GetTransactionByAddress(address string) ([]string, error)
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

// https://rest.cryptoapis.io/blockchain-data/ethereum
func (c *CryptoAPI) GetInternalTransactionByAddress(address string, timestampFrom, timestampTo uint) (*TransactionsResponse, error) {
	url := fmt.Sprintf("%s/goerli/addresses/%s/internal-by-time-range?fromTimestamp=%d&toTimestamp=%d", c.api, address, timestampFrom, timestampTo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to init request")
	}

	req = c.setHeader(req)
	rawRes, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to send request")
	}
	response := new(TransactionsResponse)
	err = json.NewDecoder(rawRes.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode data")
	}
	return response, nil
}

func (c *CryptoAPI) GetTokensHolders(response TransactionsResponse) map[string]Holder {
	txs := response.Data.Items

	holders := make(map[string]Holder)

	for _, tx := range txs {
		address := tx.RecipientAddress
		holder, ok := holders[address]
		if !ok {
			holder = Holder{}
			amount, err := strconv.Atoi(tx.TokensAmount)
			if err != nil {
				continue
			}
			holder.TokensAmount = amount
			holder.Address = address

			continue
		}

		holder.Address = address
		amount, err := strconv.Atoi(tx.TokensAmount)
		if err != nil {
			continue
		}
		holder.TokensAmount += amount

	}

	return holders
}

func (c *CryptoAPI) GetTokensHoldersByTime(response TransactionsResponse, timeFrom uint, timeTo uint) map[string]Holder {
	txs := response.Data.Items
	holders := make(map[string]Holder)

	for _, tx := range txs {
		address := tx.RecipientAddress
		holder, ok := holders[address]
		if !ok {
			holder = Holder{}
			amount, err := strconv.Atoi(tx.TokensAmount)
			if err != nil {
				continue
			}
			holder.TimeTo = timeTo
			holder.TimeFrom = timeFrom
			holder.TokensAmount = amount
			holder.Address = address

			continue
		}

		holder.Address = address
		amount, err := strconv.Atoi(tx.TokensAmount)
		if err != nil {
			continue
		}
		holder.TokensAmount += amount

	}

	return holders
}
