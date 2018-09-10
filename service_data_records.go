package main

import (
	"bytes"
	"encoding/binary"
)

type RecordDataSet []RecordData

func (f *RecordDataSet) Encode() ([]byte, error) {
	var result []byte
	buf := new(bytes.Buffer)

	for _, rd := range *f {
		rec, err := rd.Encode()
		if err != nil {
			return result, err
		}
		buf.Write(rec)
	}

	result = buf.Bytes()

	return result, nil
}

func (f *RecordDataSet) Length() uint16 {
	var result uint16

	if recBytes, err := f.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}

type ServiceDataSet []ServiceDataRecord

func (f *ServiceDataSet) Encode() ([]byte, error) {
	var result []byte
	buf := new(bytes.Buffer)

	for _, rd := range *f {
		rec, err := rd.Encode()
		if err != nil {
			return result, err
		}
		buf.Write(rec)
	}

	result = buf.Bytes()

	return result, nil
}

func (f *ServiceDataSet) Length() uint16 {
	var result uint16

	if recBytes, err := f.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}

type ServiceDataRecord struct {
	RecordLength uint16
	RecordNumber uint16
	SourceServiceOnDevice byte
	RecipientServiceOnDevice byte
	Group byte
	RecordProcessingPriority byte
	TimeFieldExists byte
	EventIDFieldExists byte
	ObjectIDFieldExists byte
	ObjectIdentifier uint32
	EventIdentifier uint32
	Time uint32
	SourceServiceType byte
	RecipientServiceType byte
	RecordDataSet
}

func (sdr *ServiceDataRecord) Encode() ([]byte, error) {
	var result []byte
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, sdr.RecordLength); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, sdr.RecordNumber); err != nil {
		return result, err
	}

	//составной байт
	flagByte := byte(0) | sdr.ObjectIDFieldExists         // 0 бит
	flagByte = flagByte | sdr.EventIDFieldExists<<1       // 1 бит
	flagByte = flagByte | sdr.TimeFieldExists<<2          // 2 бит
	flagByte = flagByte | sdr.RecordProcessingPriority<<4 // 3-4 бит
	flagByte = flagByte | sdr.Group<<5                    // 5 бит
	flagByte = flagByte | sdr.RecipientServiceOnDevice<<6 // 6 бит
	flagByte = flagByte | sdr.SourceServiceOnDevice<<7    // 7 бит
	buf.WriteByte(flagByte)

	if sdr.ObjectIDFieldExists == 1 {
		if err := binary.Write(buf, binary.LittleEndian, sdr.ObjectIdentifier); err != nil {
			return result, err
		}
	}

	if sdr.EventIDFieldExists == 1 {
		if err := binary.Write(buf, binary.LittleEndian, sdr.EventIDFieldExists); err != nil {
			return result, err
		}
	}

	if sdr.TimeFieldExists == 1 {
		if err := binary.Write(buf, binary.LittleEndian, sdr.Time); err != nil {
			return result, err
		}
	}

	if err := binary.Write(buf, binary.LittleEndian, sdr.SourceServiceType); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, sdr.RecipientServiceType); err != nil {
		return result, err
	}

	rd, err := sdr.RecordDataSet.Encode()
	if err != nil {
		return result, err
	}
	buf.Write(rd)

	result = buf.Bytes()
	return result, nil
}

type EGTS_PT_APPDATA struct {
	ServiceDataSet
}

func (ad *EGTS_PT_APPDATA) Encode() ([]byte, error) {
	return ad.ServiceDataSet.Encode()
}

func (ad *EGTS_PT_APPDATA) Length() uint16 {
	var result uint16

	if recBytes, err := ad.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}
	return result
}
