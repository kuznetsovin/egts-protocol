package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	if _, err = file.WriteString(cfg); err != nil {
		t.Fatal(err)
	}

	conf := settings{}
	if err = conf.Load(file.Name()); err != nil {
		t.Fatal(err)
	}

	testCfg := settings{
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
	}
	if diff := cmp.Diff(testCfg, conf); diff != "" {
		t.Errorf("Записи не совпадают: (-нужно +сейчас)\n%s", diff)
	}
}
