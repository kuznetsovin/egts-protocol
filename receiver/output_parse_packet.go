package main

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

type egtsParsePacket struct {
	Client              uint32         `json:"client"`
	PacketID            uint32         `json:"packet_id"`
	NavigationTimestamp int64          `json:"navigation_unix_time"`
	ReceivedTimestamp   int64          `json:"received_unix_time"`
	Latitude            float64        `json:"latitude"`
	Longitude           float64        `json:"longitude"`
	Speed               uint16         `json:"speed"`
	Pdop                uint16         `json:"pdop"`
	Hdop                uint16         `json:"hdop"`
	Vdop                uint16         `json:"vdop"`
	Nsat                uint8          `json:"nsat"`
	Ns                  uint16         `json:"ns"`
	Course              uint8          `json:"course"`
	GUID                uuid.UUID      `json:"guid"`
	AnSensors           []anSensor     `json:"an_sensors"`
	LiquidSensors       []liquidSensor `json:"liquid_sensors"`
}

func (eep *egtsParsePacket) ToBytes() ([]byte, error) {
	return json.Marshal(eep)
}

type liquidSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	ErrorFlag    string `json:"error_flag"`
	ValueMm      uint32 `json:"value_mm"`
	ValueL       uint32 `json:"value_l"`
}

type anSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	Value        uint32 `json:"value"`
}
