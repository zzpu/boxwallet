package token

import (
	"boxwallet/bccoin"
	"boxwallet/bccore"
)

type TokenLawer interface {
	GetTokenInfo(contract bccore.Token) (*bccoin.CoinInfo, error)
}
