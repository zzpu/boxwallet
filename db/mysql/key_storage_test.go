package mysql_test

import (
	"path"
	"testing"

	"github.com/zzpu/boxwallet/bcconfig"
	mysqlConf "github.com/zzpu/boxwallet/bcconfig/mysql"
	mysqlDB "github.com/zzpu/boxwallet/db/mysql"
	"github.com/zzpu/boxwallet/mock"
)

func TestNewKeyStorage(t *testing.T) {
	path1 := path.Join(mock.Gopath, mock.ProjectDir, mock.MySqlConfPath)
	provide, err := bcconfig.FromConfigString(path1, mock.YmlExt)
	if err != nil {
		t.Fail()
		t.Error(err)
		return
	}
	conf, err := mysqlConf.DecodeConfig(provide)
	if err != nil {
		t.Fail()
		t.Error(err)
		return
	}
	ksdb := mysqlDB.NewKeyStorage(conf.KeyStorage)
	if !ksdb.HasTable(&mysqlDB.Key{}) {
		ksdb.CreateTable(&mysqlDB.Key{})
	}
	t.Log(ksdb)
}
