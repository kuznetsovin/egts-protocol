## Плагин для работы с PostgreSQL

Для работы плагина нужна библиотека для работы с [pq](github.com/lib/pq)

Секция настроек выглядит следующим образом:

```
[store]
plugin = "pg.so"
host = "localhost"
port = "5672"
user = "guest"
password = "guest"
database = "receiver"
table = "points"
sslmode = "disable"

```

Описание парамеров:

- *plugin* - путь до библиотеки
- *host* - адрес сервера
- *port* - порт
- *user* - пользователь
- *password* - пароль
- *database* - имя БД
- *table* - имя таблицы для вставки. У таблицы должно быть поле *point* типа jsonb
- *sslmode* - режим защищеного соединеия

