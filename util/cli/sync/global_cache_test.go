package sync

import (
	"fmt"
	"os"
	"testing"

	tml "github.com/BurntSushi/toml"
	"github.com/33cn/externaldb/escli"
	"github.com/33cn/externaldb/proto"
)

func TestInitLastSyncSeqCache(t *testing.T) {
	var cfg proto.ConfigNew
	if _, err := tml.DecodeFile("../../../config/externaldb.toml", &cfg); err != nil {
		log.Info("init config failed", "err", err)
		os.Exit(0)
	}

	EsWrite, err := escli.NewESLongConnect(cfg.ConvertEs.Host, cfg.ConvertEs.Prefix, cfg.EsVersion, cfg.ConvertEs.User, cfg.ConvertEs.Pwd)
	if err != nil {
		fmt.Println("esCheck failed", "err", err.Error())
		log.Error("ES 连接失败，请确保ES服务正常开启，ES配置正确，网络通常",
			" host ", cfg.ConvertEs.Host, " Prefix ", cfg.ConvertEs.Prefix, " EsVersion ", cfg.EsVersion,
			" User ", cfg.ConvertEs.User, " Pwd ", cfg.ConvertEs.Pwd)
		return
	}
	if err := InitLastSyncSeqCache(EsWrite, "last_seq", cfg.Sync.StartSeq); err != nil {
		t.Errorf("InitLastSyncSeqCache() error = %v", err)
	}

}
