package main

/*
Описание конфигурационного файла
*/

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Srv      service
	RabbitMQ broker `toml:"rabbitmq"`
	Log      logSection
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

type service struct {
	Host string
	Port string
}

func (s *service) getServerAddress() string {
	return s.Host + ":" + s.Port
}

type broker struct {
	Host           string
	Port           string
	Exchange       string
	User           string
	Password       string
	RequestTimeout int `toml:"request_timeout"`
}

func (b *broker) GetConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", b.User, b.Password, b.Host, b.Port)
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
