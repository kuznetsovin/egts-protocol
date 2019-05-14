package main

/*
Описание конфигурационного файла
*/

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Srv   service
	Store map[string]string
	Log   logSection
}

func (c *Config) Load(confPath string) error {
	if _, err := toml.DecodeFile(confPath, c); err != nil {
		return fmt.Errorf("Ошибка разбора файла настроек: %v", err)
	}

	return nil
}

func (c *Config) GetListenAddress() string {
	return c.Srv.getServerAddress()
}

func (c *Config) GetLogLevel() log.Lvl {
	return c.Log.getLevel()
}

type store struct {
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	Exchange     string `toml:"exchange"`
	DeliveryMode string `toml:"delivery_mode"`
	Queue        string `toml:"queue"`
}

type service struct {
	Host       string
	Port       string
	ConLiveSec int  `toml:"con_live_sec"`
	DebugMode  bool `toml:"debug_mode"`
}

func (s *service) getEmptyConnTTL() time.Duration {
	return time.Duration(s.ConLiveSec) * time.Second
}
func (s *service) getServerAddress() string {
	return s.Host + ":" + s.Port
}

type logSection struct {
	Level string
}

func (l *logSection) getLevel() log.Lvl {
	var lvl log.Lvl

	switch l.Level {
	case "DEBUG":
		lvl = log.DEBUG
		break
	case "INFO":
		lvl = log.INFO
		break
	case "WARN":
		lvl = log.WARN
		break
	case "ERROR":
		lvl = log.ERROR
		break
	default:
		lvl = log.INFO
	}
	return lvl
}
