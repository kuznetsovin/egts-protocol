package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

// BinaryData интерфейс для работы с бинарными секциями
type BinaryData interface {
	Decode([]byte) error
	Encode() ([]byte, error)
	Length() uint16
}

// EgtsPackage стуркура для описания пакета ЕГТС
type EgtsPackage struct {
	ProtocolVersion           byte
	SecurityKeyID             byte
	Prefix                    string
	Route                     string
	EncryptionAlg             string
	Compression               string
	Priority                  string
	HeaderLength              byte
	HeaderEncoding            byte
	FrameDataLength           uint16
	PacketIdentifier          uint16
	PacketType                byte
	PeerAddress               uint16
	RecipientAddress          uint16
	TimeToLive                byte
	HeaderCheckSum            byte
	ServicesFrameData         BinaryData
	ServicesFrameDataCheckSum uint16
}

// Encode кодирует струткуру в байтовую строку
func (p *EgtsPackage) Decode(pkg []byte) (int, error) {
	var (
		err   error
		flags byte
	)
	buf := bytes.NewReader(pkg)
	if p.ProtocolVersion, err = buf.ReadByte(); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить версию протокола: %v", err)
	}

	if p.SecurityKeyID, err = buf.ReadByte(); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить идентификатор ключа: %v", err)
	}

	//разбираем флаги
	if flags, err = buf.ReadByte(); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось флаги: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)
	p.Prefix = flagBits[:2]         // flags << 7, flags << 6
	p.Route = flagBits[2:3]         // flags << 5
	p.EncryptionAlg = flagBits[3:5] // flags << 4, flags << 3
	p.Compression = flagBits[5:6]   // flags << 2
	p.Priority = flagBits[6:]       // flags << 1, flags << 0

	if p.HeaderLength, err = buf.ReadByte(); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить длину заголовка: %v", err)
	}

	if p.HeaderEncoding, err = buf.ReadByte(); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить метод кодирования: %v", err)
	}

	tmpIntBuf := make([]byte, 2)
	if _, err = buf.Read(tmpIntBuf); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить длину секции данных: %v", err)
	}
	p.FrameDataLength = binary.LittleEndian.Uint16(tmpIntBuf)

	if _, err = buf.Read(tmpIntBuf); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить идентификатор пакета: %v", err)
	}
	p.PacketIdentifier = binary.LittleEndian.Uint16(tmpIntBuf)

	if p.PacketType, err = buf.ReadByte(); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить тип пакета: %v", err)
	}

	if p.Route == "1" {
		if _, err = buf.Read(tmpIntBuf); err != nil {
			return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить адрес апк отправителя: %v", err)
		}
		p.PeerAddress = binary.LittleEndian.Uint16(tmpIntBuf)

		if _, err = buf.Read(tmpIntBuf); err != nil {
			return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить адрес апк получателя: %v", err)
		}
		p.RecipientAddress = binary.LittleEndian.Uint16(tmpIntBuf)

		if p.TimeToLive, err = buf.ReadByte(); err != nil {
			return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить TTL пакета: %v", err)
		}
	}

	if p.HeaderCheckSum, err = buf.ReadByte(); err != nil {
		return egtsPcIncHeaderform, fmt.Errorf("Не удалось получить crc заголовка: %v", err)
	}

	dataFrameBytes := make([]byte, p.FrameDataLength)
	if _, err = buf.Read(dataFrameBytes); err != nil {
		return egtsPcIncDataform, fmt.Errorf("Не считать тело пакета: %v", err)
	}
	switch p.PacketType {
	case 1:
		sfd := ServiceDataSet{}
		if err = sfd.Decode(dataFrameBytes); err != nil {
			return egtsPcDecryptError, err
		}
		p.ServicesFrameData = &sfd
		break
	default:
		return egtsPcUnsType, fmt.Errorf("Неизвестный тип пакета: %d", p.PacketType)
	}

	crcBytes := make([]byte, 2)
	if _, err = buf.Read(crcBytes); err != nil {
		return egtsPcDecryptError, fmt.Errorf("Не удалось считать crc16 пакета: %v", err)
	}
	p.ServicesFrameDataCheckSum = binary.LittleEndian.Uint16(crcBytes)

	//TODO: добавить обработку ошибок (resultCode) по протоколу
	return egtsPcOk, err
}

// Encode кодирует струткуру в байтовую строку
func (p *EgtsPackage) Encode() ([]byte, error) {
	var (
		result []byte
		err    error
		flags  uint64
	)
	buf := new(bytes.Buffer)

	if err = buf.WriteByte(p.ProtocolVersion); err != nil {
		return result, fmt.Errorf("Не удалось записать версию протокола: %v", err)
	}
	if err = buf.WriteByte(p.SecurityKeyID); err != nil {
		return result, fmt.Errorf("Не удалось записать  идентификатор ключа: %v", err)
	}

	//собираем флаги
	flagsBits := p.Prefix + p.Route + p.EncryptionAlg + p.Compression + p.Priority
	if flags, err = strconv.ParseUint(flagsBits, 2, 8); err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт флагов: %v", err)
	}

	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать флаги: %v", err)
	}

	if err = buf.WriteByte(p.HeaderLength); err != nil {
		return result, fmt.Errorf("Не удалось записать длину заголовка: %v", err)
	}

	if err = buf.WriteByte(p.HeaderEncoding); err != nil {
		return result, fmt.Errorf("Не удалось записать метод кодирования: %v", err)
	}

	tmpIntBuf := make([]byte, 2)
	binary.LittleEndian.PutUint16(tmpIntBuf, p.FrameDataLength)
	if _, err = buf.Write(tmpIntBuf); err != nil {
		return result, fmt.Errorf("Не удалось записать длину секции данных: %v", err)
	}

	binary.LittleEndian.PutUint16(tmpIntBuf, p.PacketIdentifier)
	if _, err = buf.Write(tmpIntBuf); err != nil {
		return result, fmt.Errorf("Не удалось записать идентификатор пакета: %v", err)
	}

	if err = buf.WriteByte(p.PacketType); err != nil {
		return result, fmt.Errorf("Не удалось записать идентификатор пакета: %v", err)
	}

	if p.Route == "1" {
		binary.LittleEndian.PutUint16(tmpIntBuf, p.PeerAddress)
		if _, err = buf.Write(tmpIntBuf); err != nil {
			return result, fmt.Errorf("Не удалось записать адрес апк отправителя: %v", err)
		}

		binary.LittleEndian.PutUint16(tmpIntBuf, p.RecipientAddress)
		if _, err = buf.Write(tmpIntBuf); err != nil {
			return result, fmt.Errorf("Не удалось записать адрес апк получателя: %v", err)
		}

		if err = buf.WriteByte(p.TimeToLive); err != nil {
			return result, fmt.Errorf("Не удалось записать TTL пакета: %v", err)
		}
	}

	buf.WriteByte(crc8(buf.Bytes()))

	sfrd, err := p.ServicesFrameData.Encode()
	if err != nil {
		return result, err
	}

	if len(sfrd) > 0 {
		buf.Write(sfrd)

		if err := binary.Write(buf, binary.LittleEndian, crc16(sfrd)); err != nil {
			return result, fmt.Errorf("Не удалось записать crc16 пакета: %v", err)
		}
	}

	result = buf.Bytes()
	return result, err
}
