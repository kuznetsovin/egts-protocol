package main

import (
	"github.com/kuznetsovin/egts-protocol/cli/receiver/config"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/server"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/storage"
	log "github.com/sirupsen/logrus"
	"os"
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

	//if config.Store != nil {
	//	plug, err := plugin.Open(config.Store["plugin"])
	//	if err != nil {
	//		logger.Fatalf("Не удалость загрузить плагин хранилища: %v", err)
	//	}
	//
	//	connector, err := plug.Lookup("Connector")
	//	if err != nil {
	//		logger.Fatalf("Не удалось загрузить коннектор: %v", err)
	//	}
	//
	//	store = connector.(Connector)
	//} else {
	//	store = defaultConnector{}
	//}

	store = storage.LogConnector{}

	//if err := store.Init(config.Store); err != nil {
	//	logger.Fatal(err)
	//}
	//defer store.Close()

	srv := server.New(cfg.GetListenAddress(), cfg.GetEmptyConnTTL(), store)

	srv.Run()
}
