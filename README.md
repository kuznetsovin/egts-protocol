# EGTS receiver

EGTS receiver server realization writen on Go. 

Library for implementation EGTS protocol that parsing binary packag based on 
[GOST R 54619 - 2011](./docs`/gost54619-2011.pdf) and 
[Order No. 285 of the Ministry of Transport of Russia dated July 31, 2012](./docs/mitransNo285.pdf). 
Describe fields you can find in these documents. 

More information you can read in [article](https://www.swe-notes.ru/post/protocol-egts/) (Russian).
 
Server save all navigation data from ```EGTS_SR_POS_DATA``` section. If packet have several records with 
```EGTS_SR_POS_DATA``` section, it saves all of them. 

Storage for data realized as plugins. Any plugin must have ```[store]``` section in configure file. 
Plugin interface will be described below.

If configure file has't section for a plugin (```[store]```), then packet will be print to stdout.

## Install

```bash
git clone https://github.com/kuznetsovin/egts-protocol
cd egts-protocol/tools && ./build-receiver.sh
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

## Usage only Golang EGTS library

Example for encoding packet:

```go
package main 

import (
    "github.com/kuznetsovin/egts-protocol/app/egts"
    "log"
)

func main() {
    pkg := egts.Package{
    		ProtocolVersion:  1,
    		SecurityKeyID:    0,
    		Prefix:           "00",
    		Route:            "0",
    		EncryptionAlg:    "00",
    		Compression:      "0",
    		Priority:         "11",
    		HeaderLength:     11,
    		HeaderEncoding:   0,
    		FrameDataLength:  3,
    		PacketIdentifier: 137,
    		PacketType:       egts.PtResponsePacket,
    		HeaderCheckSum:   74,
    		ServicesFrameData: &egts.PtResponse{
    			ResponsePacketID: 14357,
    			ProcessingResult: 0,
    		},
    	}
    
    rawPkg, err := pkg.Encode()
	if err != nil {
		log.Fatal(err)
	}
    
    log.Println("Bytes packet: ", rawPkg)
}

```

Example for decoding packet:

```go
package main 

import (
    "github.com/kuznetsovin/egts-protocol/app/egts"
    "log"
)

func main() {
    pkg := []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x03, 0x00, 0x89, 0x00, 0x00, 0x4A, 0x15, 0x38, 0x00, 0x33, 0xE8}
    result := egts.Package{}

    state, err := result.Decode(pkg)
    if err != nil {
 		log.Fatal(err)
 	}
    
    log.Println("State: ", state)
    log.Println("Package: ", result)
}
```

# Store plugins

That create a new plugin you must implementation ```Connector``` interface:

```go
type Connector interface {
	// setup store connection
	Init(map[string]string) error
	
	// save to store method
	Save(interface{ ToBytes() ([]byte, error) }) error
	
	// close connection with store
	Close() error
}
```

All plugins available in [store folder](/receiver/store).
