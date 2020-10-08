package setting

import (
	"path/filepath"

	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
)

const (
	DefaultConfDir = "conf"
)

type App struct {
	LogLevel log.Level
}

var AppSetting = &App{}

type Server struct {
	RunMode  string
	HttpPort int
}

var ServerSetting = &Server{}

type Etcd struct {
	Endpoints []string
}

var EtcdSetting = &Etcd{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load(filepath.Join(DefaultConfDir, "app.ini"))
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("etcd", EtcdSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
