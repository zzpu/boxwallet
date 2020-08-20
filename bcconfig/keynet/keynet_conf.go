package keynet

import (
	"github.com/mitchellh/mapstructure"
	"github.com/zzpu/boxwallet/bcconfig"
	"github.com/zzpu/boxwallet/bccore"
)

const (
	keyStorageConfigKey = "keynet"
)

type Config struct {
	Net bccore.Net
}

func DecodeConfig(cfg bcconfig.Provider) (c Config, err error) {
	m := cfg.GetStringMap(keyStorageConfigKey)
	err = mapstructure.WeakDecode(m, &c)
	return
}
