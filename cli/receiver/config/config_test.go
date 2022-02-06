package config

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	cfg := `host: "127.0.0.1"
port: "5020"
conn_ttl: 10
log_level: "DEBUG"

storage:
  rabbitmq:
    host: "localhost"
    port: "5672"
    user: "guest"
    password: "guest"
    exchange: "receiver"
  postgresql:
    host: "localhost"
    port: "5432"
    user: "postgres"
    password: "postgres"
    database: "receiver"
    table: "points"
    sslmode: "disable"
`

	file, err := ioutil.TempFile("/tmp", "config.toml")
	if !assert.NoError(t, err) {
		return
	}
	defer os.Remove(file.Name())

	if _, err = file.WriteString(cfg); !assert.NoError(t, err) {
		return
	}

	conf, err := New(file.Name())
	if assert.NoError(t, err) {
		assert.Equal(t, Settings{
			Host:     "127.0.0.1",
			Port:     "5020",
			ConnTTl:  10,
			LogLevel: "DEBUG",
			Store: map[string]map[string]string{
				"postgresql": {
					"host":     "localhost",
					"port":     "5432",
					"user":     "postgres",
					"password": "postgres",
					"database": "receiver",
					"table":    "points",
					"sslmode":  "disable",
				},
				"rabbitmq": {
					"exchange": "receiver",
					"host":     "localhost",
					"password": "guest",
					"port":     "5672",
					"user":     "guest",
				},
			},
		},
			conf,
		)
	}
}
