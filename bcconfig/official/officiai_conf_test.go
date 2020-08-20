package official

import (
	"testing"

	"github.com/zzpu/boxwallet/bcconfig"
)

func TestDecodeConfig(t *testing.T) {
	provide, err := bcconfig.FromConfigString("../daemon_cnf.yml", "yml")
	if err != nil {
		t.Fail()
		t.Error(err)
		return
	}
	conf, err := DecodeConfig(provide)
	if err != nil {
		t.Fail()
		t.Error(err)
		return
	}
	t.Log(conf)
}
