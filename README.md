# EGTS receiver

EGTS receiver server realization writen on Go. For implementation EGTS protocol usage [go-egts](https://github.com/egts/go-egts)
library.
 
Server save all navigation data from ```EGTS_SR_POS_DATA``` section. If packet have several records with 
```EGTS_SR_POS_DATA``` section, it saves all of them. 

Storage for data realized as [plugins](https://github.com/egts/egts-receiver-plugins). Any plugin must have ```[store]```
section in configure file.

If configure file has't section for a plugin (```[store]```), then packet will be print to stdout.

## Install

```bash
git clone https://github.com/egts/egts-receiver
cd egts-receiver/build && ./build-receiver.sh
```

## Run

```bash
./receiver config.toml
```

```config.toml``` - configure file

## Config format

```toml
[srv]
host = "127.0.0.1"
port = "6000"
con_live_sec = 10

[log]
level = "DEBUG"
```

Parameters description:

- *host* - bind address  
- *port* - bind port 
- *con_live_sec* - if server not received data longer time in the parameter, then the connection is closed. 
- *log* - logging level