package config

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	cfg := `host = "127.0.0.1"
port = "5020"
con_ttl = 10
log_level = "DEBUG"

[store]
plugin = "rabbitmq.so"
host = "localhost"
port = "5672"
user = "guest"
password = "guest"
exchange = "receiver"
`

	file, err := ioutil.TempFile("/tmp", "config.toml")
	if !assert.NoError(t, err) {
		return
	}
	defer os.Remove(file.Name())

	if _, err = file.WriteString(cfg); assert.NoError(t, err) {
		return
	}

	conf, err := New(file.Name())
	if assert.NoError(t, err) {
		assert.Equal(t, Settings{
			Host:    "127.0.0.1",
			Port:    "5020",
			ConnTTl: 10,
			Store: map[string]string{
				"exchange": "receiver",
				"host":     "localhost",
				"password": "guest",
				"plugin":   "rabbitmq.so",
				"port":     "5672",
				"user":     "guest",
			},
			LogLevel: "DEBUG",
		},
			conf,
		)
	}
}
