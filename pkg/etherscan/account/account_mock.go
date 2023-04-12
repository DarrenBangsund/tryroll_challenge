package account

type EtherscanAccountMock struct {
}

func NewAccountModuleMock() *EtherscanAccountMock {
	return &EtherscanAccountMock{}
}

func (ea *EtherscanAccountMock) GetERC20TokenTx(contractAddress, address, startBlock, offset, page string) ([]ERC20TokenTx, error) {
	return []ERC20TokenTx{
		{
			From:  "mock-from",
			To:    "mock-to",
			Value: "100",
			Hash:  "mock-hash",
		},
	}, nil
}
