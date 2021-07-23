## PostgreSQL plugin

Plugin depends on [pq](http://github.com/lib/pq) package

Config section:

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

Parameters description:

- *plugin* - path to *.so library file  
- *host* - postgres server address
- *port* - postgres port
- *user* - postgres user
- *password* - user password
- *database* - db name
- *table* - table in db where will be insert data. Table must have *point* field (jsonb type)
- *sslmode* - postgres ssl mode

Simple db table for inserting data example:

```sql
create table points ( 
    point jsonb
);
```