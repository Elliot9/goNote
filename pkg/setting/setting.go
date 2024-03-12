package setting

import (
	"github.com/go-ini/ini"
	"log"
)

type Server struct {
	HttpPort int
}

var ServerSetting = &Server{}

type Database struct {
	Host     string
	User     string
	Password string
	Name     string
}

var DatabaseSetting = &Database{}

var config *ini.File

func Setup() {
	var err error
	config, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", ServerSetting)
	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
