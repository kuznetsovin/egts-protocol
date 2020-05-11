# Go implements EGTS protocol

This library parsing binary package according [GOST R 54619 - 2011](./docs`/gost54619-2011.pdf) and 
[Order No. 285 of the Ministry of Transport of Russia dated July 31, 2012](./docs/mitransNo285.pdf). Describe 
fields you can find in these documents. 

More information you can read in [article](https://www.swe-notes.ru/post/protocol-egts/) (Russian).

## Usage

Example for encoding packet:

```go
package main 

import (
    "log"
    "github.com/egts/go-egts/egts"
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
    "log"
    "github.com/egts/go-egts/egts"
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

Full example usage library you can see in [egts-receiver](https://github.com/egts/egts-receiver)