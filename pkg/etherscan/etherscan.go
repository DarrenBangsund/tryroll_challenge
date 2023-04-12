package etherscan

import (
	"tryroll/challenge/pkg/etherscan/account"

	"github.com/go-resty/resty/v2"
)

const ETHERSCANURL = "https://api.etherscan.io/api"

type IAccountModule interface {
	GetERC20TokenTx(contractAddress, address, startBlock, offset, page string) ([]account.ERC20TokenTx, error)
}

type Etherscan struct {
	Account IAccountModule
}

// https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=0x9355372396e3F6daF13359B7b607a3374cc638e0&page=1&sort=asc&apikey={ENTER_KEY_HERE}

type EtherscanConfig struct {
	BaseURL string
	APIKey  string
	Mock    bool
}

func NewEtherscan(config *EtherscanConfig) *Etherscan {
	c := resty.New().
		SetBaseURL(config.BaseURL).
		SetQueryParam("apikey", config.APIKey)

	if config.Mock {
		return &Etherscan{
			Account: account.NewAccountModuleMock(),
		}
	} else {
		return &Etherscan{
			Account: account.NewAccountModule(*c),
		}
	}
}
