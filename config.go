package main

/*
Описание конфигурационного файла
*/

import (
	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Srv    service
	Broker broker
}

func (c *Config) Load(confPath string) error {
	if _, err := toml.DecodeFile(confPath, c); err != nil {
		log.Warn("Error decode file")
		return err
	}

	return nil
}

func (c *Config) GetListenAddress() string {
	return c.Srv.GetServerAddress()
}

type service struct {
	Host string
	Port string
}

func (s *service) GetServerAddress() string {
	return s.Host + ":" + s.Port
}

type broker struct {
	Host           string
	Port           string
	Queue          string
	User           string
	Password       string
	RequestTimeout int `toml:"request_timeout"`
}
