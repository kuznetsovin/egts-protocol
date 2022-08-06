package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type postgresqlTest struct {
	conn *sql.DB
	table string
}

func (pt *postgresqlTest) pointCount() (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", pt.table)
	result := 0

	rows, err := pt.conn.Query(query)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return result, err
		}
	}

	return result, err
}

func initTestPostgresql(conf map[string]string) (postgresqlTest, error) {
	var err error = nil

	result := postgresqlTest{
		conn: nil,
	}

	connStr := fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=%s",
		conf["database"], conf["host"], conf["port"], conf["user"], conf["password"], conf["sslmode"])

	if result.conn, err = sql.Open("postgres", connStr); err != nil {
		return result, err
	}

	if err = result.conn.Ping(); err != nil {
		return result, err
	}

	result.table = conf["table"]

	creat_table_query := fmt.Sprintf("create table %s ( point jsonb );", result.table)

	_, err = result.conn.Exec(creat_table_query)

	return result, err
}
