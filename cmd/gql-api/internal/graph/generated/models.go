// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

type ERC20Tx struct {
	To     string `json:"to"`
	From   string `json:"from"`
	Amount string `json:"amount"`
	Hash   string `json:"hash"`
}

type GetERC20TxReq struct {
	Amount *string `json:"amount,omitempty"`
	To     *string `json:"to,omitempty"`
	From   *string `json:"from,omitempty"`
	Limit  *string `json:"limit,omitempty"`
	Offset *string `json:"offset,omitempty"`
}

type GetERC20TxRes struct {
	Txs []*ERC20Tx `json:"txs"`
}
