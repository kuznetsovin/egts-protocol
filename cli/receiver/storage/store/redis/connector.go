package redis

/*
Плагин для работы с Redis.
Плагин отправляет пакет в redis очередь.

Раздел настроек, которые должны отвечають в конфиге для подключения хранилища:

server = "localhost:6379"
queue = "egts"
password = ""
db = 0
*/

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type Connector struct {
	conn   *redis.Client
	queue  string
	config map[string]string
}

func (c *Connector) Init(cfg map[string]string) error {
	var (
		err error
	)
	if cfg == nil {
		return fmt.Errorf("не корректная ссылка на конфигурацию")
	}
	c.config = cfg

	addr, ok := c.config["server"]
	if !ok {
		return fmt.Errorf("не задан адрес redis сервера")
	}

	configDb, _ := c.config["db"]
	if !ok {
		configDb = "0"
	}

	db, err := strconv.Atoi(configDb)
	if err != nil {
		return fmt.Errorf("не корретное имя redis бд: %v", err)
	}

	c.conn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.config["password"],
		DB:       db,
	})

	c.queue, ok = c.config["queue"]
	if !ok {
		return fmt.Errorf("не корретное имя redis очереди")
	}
	return err
}

func (c *Connector) Save(msg interface{ ToBytes() ([]byte, error) }) error {
	if msg == nil {
		return fmt.Errorf("Не корректная ссылка на пакет")
	}

	innerPkg, err := msg.ToBytes()
	if err != nil {
		return fmt.Errorf("Ошибка сериализации  пакета: %v", err)
	}

	if err := c.conn.Publish(context.Background(), c.queue, innerPkg).Err(); err != nil {
		panic(err)
	}

	return nil
}

func (c *Connector) Close() error {
	return c.conn.Close()
}
