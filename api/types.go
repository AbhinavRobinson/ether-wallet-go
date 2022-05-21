package api

import (
	app "ether-wallet-go/app"
)

type CreateSimpleWalletBody struct {
	ChainConfig app.ChainConfig
	Nickname    string
}
