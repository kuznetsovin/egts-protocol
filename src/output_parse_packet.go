package main

import (
	"encoding/json"
	"time"
)

type EgtsParsePacket struct {
	Client         uint32           `json:"client"`
	PacketID       uint32           `json:"packet_id"`
	NavigationTime time.Time        `json:"navigation_time"`
	Latitude       float64          `json:"latitude"`
	Longitude      float64          `json:"longitude"`
	Speed          uint16           `json:"speed"`
	Pdop           uint16           `json:"pdop"`
	Nsat           uint8            `json:"nsat"`
	Course         uint8            `json:"course"`
	AnSensors      map[uint8]uint32 `json:"an_sensors"`
	LiquidSensors  []LiquidSensor   `json:"liquid_sensors"`
}

func (eep *EgtsParsePacket) ToBytes() ([]byte, error) {
	return json.Marshal(eep)
}

type LiquidSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	ErrorFlag    string `json:"error_flag"`
	ValueMm      uint32 `json:"value_mm"`
	ValueL       uint32 `json:"value_l"`
}
