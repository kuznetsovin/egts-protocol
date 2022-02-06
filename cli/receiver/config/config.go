package config

/*
Описание конфигурационного файла
*/

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Host     string                       `yaml:"host"`
	Port     string                       `yaml:"port"`
	ConnTTl  int                          `yaml:"conn_ttl"`
	LogLevel string                       `yaml:"log_level"`
	Store    map[string]map[string]string `yaml:"storage"`
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
	data, err := os.ReadFile(confPath)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(data, &c)
	return c, err
}
