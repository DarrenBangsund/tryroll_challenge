package convert

import (
	"tryroll/challenge/cmd/tryroll-service/internal/models"
	"tryroll/challenge/pkg/etherscan/account"

	"github.com/ethereum/go-ethereum/common/math"
)

func EtherscanTxToDAL(input account.ERC20TokenTx) (models.ERC20Tx, error) {
	value, _ := math.ParseBig256(input.Value)

	return models.ERC20Tx{
		From:  input.From,
		To:    input.To,
		Value: value,
		Hash:  input.Hash,
	}, nil
}
