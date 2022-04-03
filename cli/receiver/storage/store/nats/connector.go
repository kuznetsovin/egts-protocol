package nats

/*
Плагин для работы с NATS.
Плагин отправляет пакет в топик NATS messaging system.

Раздел настроек, которые должны отвечають в конфиге для подключения хранилища:

servers = "nats://localhost:1222, nats://localhost:1223, nats://localhost:1224"
topic = "receiver"
*/

import (
	"fmt"
	natsLib "github.com/nats-io/nats.go"
)

type Connector struct {
	connection *natsLib.Conn
	config     map[string]string
}

func (c *Connector) Init(cfg map[string]string) error {
	var (
		err error
	)
	if cfg == nil {
		return fmt.Errorf("Не корректная ссылка на конфигурацию")
	}
	c.config = cfg

	var options = make([]natsLib.Option, 3)

	options = append(options, natsLib.Name(fmt.Sprintf("EGTS handler, topic: %s", c.config["topic"])))

	if user, uOk := c.config["user"]; uOk {
		if password, pOk := c.config["password"]; pOk {
			options = append(options, natsLib.UserInfo(user, password))
		}
	}

	if c.connection, err = natsLib.Connect(c.config["servers"], options...); err != nil {
		return fmt.Errorf("Ошибка подключения к nats шине: %v", err)
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

	if err = c.connection.Publish(c.config["topic"], innerPkg); err != nil {
		return fmt.Errorf("Не удалось отправить сообщение в топик: %v", err)
	}
	return nil
}

func (c *Connector) Close() error {
	c.connection.Close()
	return nil
}
