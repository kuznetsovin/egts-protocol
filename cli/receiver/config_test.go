package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	cfg := `[srv]
host = "127.0.0.1"
port = "5020"
con_live_sec = 10

[store]
plugin = "rabbitmq.so"
host = "localhost"
port = "5672"
user = "guest"
password = "guest"
exchange = "receiver"

[log]
level = "DEBUG"`

	file, err := ioutil.TempFile("/tmp", "config.toml")
	if !assert.NoError(t, err) {
		return
	}
	defer os.Remove(file.Name())

	if _, err = file.WriteString(cfg); assert.NoError(t, err) {
		return
	}

	conf := settings{}
	if assert.NoError(t, conf.Load(file.Name())) {
		assert.Equal(t, settings{
			Srv: service{
				Host:       "127.0.0.1",
				Port:       "5020",
				ConLiveSec: 10,
			},
			Store: map[string]string{
				"exchange": "receiver",
				"host":     "localhost",
				"password": "guest",
				"plugin":   "rabbitmq.so",
				"port":     "5672",
				"user":     "guest",
			},
			Log: logSection{Level: "DEBUG"},
		},
			conf,
		)
	}
}
