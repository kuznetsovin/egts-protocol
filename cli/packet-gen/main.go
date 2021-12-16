package main

import (
	"flag"
	"fmt"
	"github.com/kuznetsovin/egts-protocol/libs/egts"
	"net"
	"os"
	"time"
)

/*
EGTS packet generator.

Util create egts packet  from setting parameters.

Usage:
  -pid int
    	Packet identifier (require)
  -oid int
    	Client identifier (require)
  -time string
    	Timestamp in RFC 3339 format (require)
  -lat float
    	Latitude
  -liquid int
    	Liquid level for first sensor
  -lon float
    	Longitude
  -server string
    	Egts server address in format <ip>:<port> (default "localhost:5555")
  -timeout int
    	Ack waiting time in seconds, Default: 5

Example

```
./packet-gen --pid 1 --oid 12 --time 2021-12-16T09:12:00Z --lat 45 --lon 60.344 --server localhost:5555
```

Created by Igor Kuznetsov
*/

func main() {

	pid := 0
	oid := 0
	ts := ""
	liqLvl := 0
	lat := 0.0
	lon := 0.0
	server := ""
	ackTimeout := 0

	flag.IntVar(&pid, "pid", 0, "Packet identifier (require)")
	flag.IntVar(&oid, "oid", 0, "Client identifier (require)")
	flag.StringVar(&ts, "time", "", "Timestamp in RFC 3339 format (require)")
	flag.IntVar(&liqLvl, "liquid", 0, "Liquid level for first sensor")
	flag.Float64Var(&lat, "lat", 0, "Latitude")
	flag.Float64Var(&lon, "lon", 0, "Longitude")
	flag.IntVar(&ackTimeout, "timeout", 0, "Ack waiting time in seconds, Default: 5")
	flag.StringVar(&server, "server", "localhost:5555", "Egts server address in format <ip>:<port>")

	flag.Parse()

	if pid == 0 {
		fmt.Println("Packet identifier is require. See to help (-h)")
		os.Exit(1)
	}

	if oid == 0 {
		fmt.Println("Client identifier is require. See to help (-h)")
		os.Exit(1)
	}

	if ts == "" {
		fmt.Println("Timestamp is require. See to help (-h)")
		os.Exit(1)
	}
	timestamp, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		fmt.Println("Parsing timestamp failed: ", timestamp)
		os.Exit(1)
	}

	egtsSendPacket := egts.Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "10",
		HeaderLength:     11,
		HeaderEncoding:   0,
		PacketIdentifier: uint16(pid),
		PacketType:       1,
		ServicesFrameData: &egts.ServiceDataSet{
			egts.ServiceDataRecord{
				RecordNumber:             1,
				SourceServiceOnDevice:    "0",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "10",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "1",
				ObjectIDFieldExists:      "1",
				EventIdentifier:          3436,
				ObjectIdentifier:         uint32(oid),
				SourceServiceType:        2,
				RecipientServiceType:     2,
				RecordDataSet: egts.RecordDataSet{
					egts.RecordData{
						SubrecordType: 16,
						SubrecordData: &egts.SrPosData{
							NavigationTime:      time.Date(2021, time.February, 20, 0, 30, 40, 0, time.UTC),
							Latitude:            lat,
							Longitude:           lon,
							ALTE:                "1",
							LOHS:                "0",
							LAHS:                "0",
							MV:                  "1",
							BB:                  "1",
							CS:                  "0",
							FIX:                 "1",
							VLD:                 "1",
							DirectionHighestBit: 0,
							AltitudeSign:        0,
							Speed:               34,
							Direction:           172,
							Odometer:            191,
							DigitalInputs:       144,
							Source:              0,
							Altitude:            30,
						},
					},
					egts.RecordData{
						SubrecordType: 27,
						SubrecordData: &egts.SrLiquidLevelSensor{
							LiquidLevelSensorErrorFlag: "1",
							LiquidLevelSensorValueUnit: "00",
							RawDataFlag:                "0",
							LiquidLevelSensorNumber:    1,
							ModuleAddress:              uint16(1),
							LiquidLevelSensorData:      uint32(liqLvl),
						},
					},
				},
			},
		},
	}

	sendBytes, err := egtsSendPacket.Encode()
	if err != nil {
		fmt.Println("Encode message failed: ", err)
		os.Exit(1)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed:", err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Dial failed:", err)
		os.Exit(1)
	}

	_, err = conn.Write(sendBytes)
	if err != nil {
		fmt.Println("Write to server failed:", err)
		os.Exit(1)
	}

	ackBuf := make([]byte, 1024)

	_ = conn.SetReadDeadline(time.Now().Add(time.Duration(ackTimeout) * time.Second))
	ackLen, err := conn.Read(ackBuf)
	if err != nil {
		fmt.Println("Read from server failed:", err)
		os.Exit(1)
	}

	ackPacket := egts.Package{}
	_, err = ackPacket.Decode(ackBuf[:ackLen])
	if err != nil {
		fmt.Println("Parse ack packet failed:", err)
		os.Exit(1)
	}

	ack, ok := ackPacket.ServicesFrameData.(*egts.PtResponse)
	if !ok {
		fmt.Println("Received packet is not egts ack")
		os.Exit(1)
	}

	if ack.ResponsePacketID != egtsSendPacket.PacketIdentifier {
		fmt.Printf("Incorrect response packet id: %d (actual) != %d (expected)",
			ack.ResponsePacketID, egtsSendPacket.PacketIdentifier)
		os.Exit(1)
	}

	if ack.ProcessingResult != 0 {
		fmt.Printf("Incorrect processing result: %d (actual) != 0 (expected)", ack.ProcessingResult)
		os.Exit(1)
	}

	for _, rec := range *ack.SDR.(*egts.ServiceDataSet) {
		for _, subRec := range rec.RecordDataSet {
			if subRec.SubrecordType == egts.SrRecordResponseType {
				if response, ok := subRec.SubrecordData.(*egts.SrResponse); ok {
					if response.RecordStatus != 0 {
						fmt.Printf("Incorrect record status: %d (actual) != 0 (expected)", response.RecordStatus)
						os.Exit(1)
					}
				}
			}
		}
	}

	fmt.Println("Packet sent and processed correct")
	os.Exit(1)
}
