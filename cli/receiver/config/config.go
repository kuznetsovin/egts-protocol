package config

/*
Описание конфигурационного файла
*/

import (
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/BurntSushi/toml"
)

type Settings struct {
	Host     string
	Port     string
	ConnTTl  int `toml:"conn_ttl"`
	Store    map[string]string
	LogLevel string
}

func (s *Settings) GetEmptyConnTTL() time.Duration {
	return time.Duration(s.ConnTTl) * time.Second
}
func (s *Settings) GetListenAddress() string {
	return s.Host + ":" + s.Port
}

func (s *Settings) GetLogLevel() log.Level {
	var lvl log.Level

	switch s.LogLevel {
	case "DEBUG":
		lvl = log.DebugLevel
		break
	case "INFO":
		lvl = log.InfoLevel
		break
	case "WARN":
		lvl = log.WarnLevel
		break
	case "ERROR":
		lvl = log.ErrorLevel
		break
	default:
		lvl = log.InfoLevel
	}
	return lvl
}

func New(confPath string) (Settings, error) {
	c := Settings{}
	_, err := toml.DecodeFile(confPath, &c)
	return c, err
}
