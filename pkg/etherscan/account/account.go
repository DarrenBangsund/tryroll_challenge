package account

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type EtherscanAccount struct {
	client *resty.Client
}

func NewAccountModule(r resty.Client) *EtherscanAccount {
	c := r.SetQueryParam("module", "account")

	return &EtherscanAccount{
		client: c,
	}
}

type ERC20TokenTx struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	Hash  string `json:"hash"`
	// BlockNumber       string `json:"blockNumber"`
	// TimeStamp         string `json:"timeStamp"`
	// Nonce             string `json:"nonce"`
	// BlockHash         string `json:"blockHash"`
	// ContractAddress   string `json:"contractAddress"`
	// TokenName         string `json:"tokenName"`
	// TokenSymbol       string `json:"tokenSymbol"`
	// TokenDecimal      string `json:"tokenDecimal"`
	// TransactionIndex  string `json:"transactionIndex"`
	// Gas               string `json:"gas"`
	// GasPrice          string `json:"gasPrice"`
	// GasUsed           string `json:"gasUsed"`
	// CumulativeGasUsed string `json:"cumulativeGasUsed"`
	// Input             string `json:"input"`
	// Confirmations     string `json:"confirmations"`
}

type GetERC20TokenTxResponse struct {
	Results []ERC20TokenTx `json:"result"`
}

func (ea *EtherscanAccount) GetERC20TokenTx(contractAddress, address, startBlock, offset, page string) ([]ERC20TokenTx, error) {
	r, err := ea.client.R().
		SetQueryParam("action", "tokentx").
		SetQueryParam("contractaddress", contractAddress).
		SetQueryParam("address", address).
		SetQueryParam("startblock", startBlock).
		SetQueryParam("offset", offset).
		SetQueryParam("page", page).
		Get("")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}

	res := &GetERC20TokenTxResponse{}

	if err := json.Unmarshal(r.Body(), res); err != nil {
		return nil, fmt.Errorf("failed to extract body from response: %w", err)
	}

	return res.Results, nil
}
