package token

import (
	"github.com/zzpu/boxwallet/bccoin"
	"github.com/zzpu/boxwallet/bccore"
)

type TokenLawer interface {
	GetTokenInfo(contract bccore.Token) (*bccoin.CoinInfo, error)
}
