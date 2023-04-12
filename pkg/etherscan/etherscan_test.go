package etherscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Account(t *testing.T) {
	client := NewEtherscan(&EtherscanConfig{
		BaseURL: "",
		APIKey:  "",
		Mock:    true,
	})

	res, err := client.Account.GetERC20TokenTx("", "", "", "", "")
	assert.NoError(t, err)
	assert.Equal(t, len(res), 1)
}
