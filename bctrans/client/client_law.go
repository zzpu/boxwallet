package client

import (
	"github.com/zzpu/boxwallet/bccoin"
	"github.com/zzpu/boxwallet/bccore"
	"github.com/zzpu/boxwallet/bckey"
	"github.com/zzpu/boxwallet/signature"
)

type TxInfo struct {
	T       bccore.BloclChainType
	TxBytes []byte
}

type Walleter interface {
	GetBalance(address string, token string, local bool) (balance bccoin.CoinAmounter, err error)
	CreateTx(addrsFrom []string, token string, addrsTo []*bccoin.AddressAmount, feeCeo float64) (txu signature.TxUtil, err error)
	SendTx(txu signature.TxUtil) error
	GetNewAddress() (bckey.GenericKey, error)
	ImportAddress(address string, local bool) error
	ChooseClientNode() (local bool, err error)
}
