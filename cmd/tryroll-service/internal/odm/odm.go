package odm

import "tryroll/challenge/cmd/tryroll-service/internal/db"

// CollectionName enumerates the collections this ODM supports
type CollectionName string

const (
	// CollectionNameERC20Tx is the name for the ERC20Tx collection
	CollectionNameERC20Tx CollectionName = "ERC20Tx"
)

// ODM wraps the collections
type ODM struct {
	ERC20TxCollection *db.Collection
}

// NewODM configures and returns the collection implementations
func NewODM() *ODM {
	erc20Collection := db.NewCollection(string(CollectionNameERC20Tx))

	return &ODM{
		ERC20TxCollection: erc20Collection,
	}
}
