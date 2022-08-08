package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlTest struct {
	conn  *sql.DB
	table string
}

func (pt *mysqlTest) pointCount() (int, error) {
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

func initTestMysql(conf map[string]string) (mysqlTest, error) {
	var err error = nil

	result := mysqlTest{
		conn: nil,
	}

	if result.conn, err = sql.Open("mysql", conf["uri"]); err != nil {
		return result, err
	}

	if err = result.conn.Ping(); err != nil {
		return result, err
	}

	result.table = conf["table"]

	creat_table_query := fmt.Sprintf("create table %s ( point json );", result.table)

	_, err = result.conn.Exec(creat_table_query)

	return result, err
}
