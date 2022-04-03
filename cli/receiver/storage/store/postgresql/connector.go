package postgresql

/*
Плагин для работы с PostgreSQL.
Плагин сохраняет пакет в jsonb поле point у заданной в настройках таблице.

Раздел настроек, которые должны отвечають в конфиге для подключения хранилища:

host = "localhost"
port = "5432"
user = "postgres"
password = "postgres"
database = "receiver"
table = "points"
sslmode = "disable"
*/

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Connector struct {
	connection *sql.DB
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
	connStr := fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=%s",
		c.config["database"], c.config["host"], c.config["port"], c.config["user"], c.config["password"], c.config["sslmode"])
	if c.connection, err = sql.Open("postgres", connStr); err != nil {
		return fmt.Errorf("Ошибка подключения к postgresql: %v", err)
	}

	if err = c.connection.Ping(); err != nil {
		return fmt.Errorf("postgresql недоступен: %v", err)
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

	insertQuery := fmt.Sprintf("INSERT INTO %s (point) VALUES ($1)", c.config["table"])
	if _, err = c.connection.Exec(insertQuery, innerPkg); err != nil {
		return fmt.Errorf("Не удалось вставить запись: %v", err)
	}
	return nil
}

func (c *Connector) Close() error {
	return c.connection.Close()
}
