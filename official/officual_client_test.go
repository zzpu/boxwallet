package official_test

import (
	"path"
	"testing"

	"github.com/zzpu/boxwallet/bcconfig"
	"github.com/zzpu/boxwallet/bctrans/clientseries"
	"github.com/zzpu/boxwallet/mock"
)

func TestNewBtcNode(t *testing.T) {
	path := path.Join(mock.Gopath, mock.ProjectDir, mock.ConfigDir, "official/btc.yml")
	provide, err := bcconfig.FromConfigString(path, mock.YmlExt)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	client := clientseries.NewOmniSeriesClient(provide)
	t.Log(client)
}

func TestNewLtcNode(t *testing.T) {
	path := path.Join(mock.Gopath, mock.ProjectDir, mock.ConfigDir, "official/ltc.yml")
	provide, err := bcconfig.FromConfigString(path, mock.YmlExt)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	client := clientseries.NewLtcSeriesClient(provide)
	t.Log(client)
}

func TestNewUsdtNode(t *testing.T) {
	path := path.Join(mock.Gopath, mock.ProjectDir, mock.ConfigDir, "official/usdt.yml")
	provide, err := bcconfig.FromConfigString(path, mock.YmlExt)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	client := clientseries.NewOmniSeriesClient(provide)
	t.Log(client)
}
func TestNewEthNode(t *testing.T) {
	path := path.Join(mock.Gopath, mock.ProjectDir, mock.ConfigDir, "official/eth.yml")
	provide, err := bcconfig.FromConfigString(path, mock.YmlExt)
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	client := clientseries.NewEthSeriesClient(provide)
	t.Log(client)
}
