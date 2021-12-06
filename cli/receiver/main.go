package main

import (
	"github.com/kuznetsovin/egts-protocol/cli/receiver/config"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/server"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/storage"
	log "github.com/sirupsen/logrus"
	"os"
	"plugin"
)

func main() {
	var (
		store storage.Connector
	)

	if len(os.Args) != 2 {
		log.Fatalf("Не задан путь до конфига")
	}

	cfg, err := config.New(os.Args[1])
	if err != nil {
		log.Fatalf("Ошибка парсинга конфига: %v", err)
	}

	log.SetLevel(cfg.GetLogLevel())

	if cfg.Store != nil {
		plug, err := plugin.Open(cfg.Store["plugin"])
		if err != nil {
			log.WithField("err", err).Fatal("Не удалось загрузить плагин хранилища")
		}

		connector, err := plug.Lookup("Connector")
		if err != nil {
			log.WithField("err", err).Fatal("Не удалось загрузить коннектор")
		}

		store = connector.(storage.Connector)
	} else {
		store = storage.LogConnector{}
	}

	if err := store.Init(cfg.Store); err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	srv := server.New(cfg.GetListenAddress(), cfg.GetEmptyConnTTL(), store)

	srv.Run()
}
