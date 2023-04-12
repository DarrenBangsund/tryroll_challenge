package app

import (
	"fmt"
	"strconv"
	"tryroll/challenge/cmd/tryroll-service/internal/db"
	"tryroll/challenge/cmd/tryroll-service/internal/models"
	"tryroll/challenge/cmd/tryroll-service/internal/odm"

	"github.com/ethereum/go-ethereum/common/math"
)

// App defines the app layer and abstracts the domain from the caller
type App struct {
	ODM *odm.ODM
}

// NewApp creates a new app
func NewApp(odm *odm.ODM) *App {
	return &App{
		ODM: odm,
	}
}

// AddERC20TokenTx creates a new tx and inserts it into the db
func (a *App) AddERC20TokenTx(to, from, value, hash string) error {
	bigValue, ok := math.ParseBig256(value)
	if !ok {
		return fmt.Errorf("failed to parse value")
	}

	erc20Tx := models.NewERC20Tx(from, to, hash, bigValue)

	a.ODM.ERC20TxCollection.Insert(erc20Tx)

	return nil
}

// GetERC20TokenTx will fetch the txs that match the filters
func (a *App) GetERC20TokenTx(amount, to, from, limit, offset string) ([]*models.ERC20Tx, error) {
	query := db.NewQuery()

	if amount != "" {
		bigAmount, ok := math.ParseBig256(amount)
		if !ok {
			return nil, fmt.Errorf("failed to parse input")
		}

		query.WithFilter(
			db.Filter{
				Key: "amount",
				Func: func(value interface{}) (bool, error) {
					tx := value.(*models.ERC20Tx)

					return tx.Value.Cmp(bigAmount) > 0, nil
				},
			},
		)
	}

	if to != "" {
		query.WithFilter(db.Filter{
			Key:   "to",
			Value: to,
			Op:    db.FilterOpEQ,
		})
	}

	if from != "" {
		query.WithFilter(db.Filter{
			Key:   "from",
			Value: from,
			Op:    db.FilterOpEQ,
		})
	}

	if limit != "" {
		limInt, err := strconv.ParseUint(limit, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse limit: %w", err)
		}

		query.WithLimit(uint(limInt))
	}

	if offset != "" {
		offInt, err := strconv.ParseUint(offset, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse offset: %w", err)
		}

		query.WithOffset(uint(offInt))
	}

	txs, err := a.ODM.ERC20TxCollection.Find(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ERC20Txs: %w", err)
	}

	out := []*models.ERC20Tx{}

	for _, tx := range txs {
		out = append(out, tx.(*models.ERC20Tx))
	}

	return out, nil
}
