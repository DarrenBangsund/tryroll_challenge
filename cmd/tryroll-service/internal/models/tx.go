package models

import (
	"math/big"
)

// ERC20Tx is the ERC20Tx domain object
type ERC20Tx struct {
	From  string   `db:"from"`
	To    string   `db:"to"`
	Value *big.Int `db:"value"`
	Hash  string   `db:"hash"`
}

// NewERC20Tx configures and returns an ERC20Tx object
func NewERC20Tx(from, to, hash string, value *big.Int) *ERC20Tx {
	return &ERC20Tx{
		From:  from,
		To:    to,
		Hash:  hash,
		Value: value,
	}
}
