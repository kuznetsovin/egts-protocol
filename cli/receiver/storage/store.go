package storage

import (
	"errors"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/storage/store/nats"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/storage/store/postgresql"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/storage/store/rabbitmq"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/storage/store/redis"
	"github.com/kuznetsovin/egts-protocol/cli/receiver/storage/store/tarantool_queue"
)

var InvalidStorage = errors.New("storage not found")
var UnknownStorage = errors.New("storage isn't support yet")

type Store interface {
	Connector
	Saver
}

//Saver интерфейс для подключения внешних хранилищ
type Saver interface {
	// Save сохранение в хранилище
	Save(interface{ ToBytes() ([]byte, error) }) error
}

//Connector интерфейс для подключения внешних хранилищ
type Connector interface {
	// Init установка соединения с хранилищем
	Init(map[string]string) error

	// Close закрытие соединения с хранилищем
	Close() error
}

//Repository набор выходных хранилищ
type Repository struct {
	storages []Saver
}

//AddStore добавляет хранилище для сохранения данных
func (r *Repository) AddStore(s Saver) {
	r.storages = append(r.storages, s)
}

//Save сохраняет данные во все установленные хранилища
func (r *Repository) Save(m interface{ ToBytes() ([]byte, error) }) error {
	for _, store := range r.storages {
		if err := store.Save(m); err != nil {
			return err
		}
	}
	return nil
}

//LoadStorages загружает хранилища из структуры конфига
func (r *Repository) LoadStorages(storages map[string]map[string]string) error {
	if len(storages) == 0 {
		return InvalidStorage
	}

	var db Store
	for store, params := range storages {
		switch store {
		case "rabbitmq":
			db = &rabbitmq.Connector{}
		case "postgresql":
			db = &postgresql.Connector{}
		case "nats":
			db = &nats.Connector{}
		case "tarantool_queue":
			db = &tarantool_queue.Connector{}
		case "redis":
			db = &redis.Connector{}
		default:
			return UnknownStorage
		}

		if err := db.Init(params); err != nil {
			return err
		}

		r.AddStore(db)
	}
	return nil
}

//NewRepository создает пустой репозиторий
func NewRepository() *Repository {
	return &Repository{}
}
