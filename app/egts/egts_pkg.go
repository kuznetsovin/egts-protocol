package egts

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
)

const DEFAULT_HEADER_LEN = 11

// Package стуркура для описания пакета ЕГТС
type Package struct {
	ProtocolVersion           byte       `json:"PRV"`
	SecurityKeyID             byte       `json:"SKID"`
	Prefix                    string     `json:"PRF"`
	Route                     string     `json:"RTE"`
	EncryptionAlg             string     `json:"ENA"`
	Compression               string     `json:"CMP"`
	Priority                  string     `json:"PR"`
	HeaderLength              byte       `json:"HL"`
	HeaderEncoding            byte       `json:"HE"`
	FrameDataLength           uint16     `json:"FDL"`
	PacketIdentifier          uint16     `json:"PID"`
	PacketType                byte       `json:"PT"`
	PeerAddress               uint16     `json:"PRA"`
	RecipientAddress          uint16     `json:"RCA"`
	TimeToLive                byte       `json:"TTL"`
	HeaderCheckSum            byte       `json:"HCS"`
	ServicesFrameData         BinaryData `json:"SFRD"`
	ServicesFrameDataCheckSum uint16     `json:"SFRCS"`
}

// Decode разбирает набор байт в структуру пакета
func (p *Package) Decode(content []byte) (uint8, error) {
	var (
		err   error
		flags byte
	)
	buf := bytes.NewReader(content)
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

	if p.HeaderCheckSum != crc8(content[:p.HeaderLength-1]) {
		return egtsPcHeaderCrcError, fmt.Errorf("Не верная сумма заголовка пакета")
	}

	dataFrameBytes := make([]byte, p.FrameDataLength)
	if _, err = buf.Read(dataFrameBytes); err != nil {
		return egtsPcIncDataform, fmt.Errorf("Не считать тело пакета: %v", err)
	}
	switch p.PacketType {
	case PtAppdataPacket:
		p.ServicesFrameData = &ServiceDataSet{}
		break
	case PtResponsePacket:
		p.ServicesFrameData = &PtResponse{}
		break
	default:
		return egtsPcUnsType, fmt.Errorf("Неизвестный тип пакета: %d", p.PacketType)
	}

	if err = p.ServicesFrameData.Decode(dataFrameBytes); err != nil {
		return egtsPcDecryptError, err
	}

	crcBytes := make([]byte, 2)
	if _, err = buf.Read(crcBytes); err != nil {
		return egtsPcDecryptError, fmt.Errorf("Не удалось считать crc16 пакета: %v", err)
	}
	p.ServicesFrameDataCheckSum = binary.LittleEndian.Uint16(crcBytes)

	if p.ServicesFrameDataCheckSum != crc16(content[p.HeaderLength:uint16(p.HeaderLength)+p.FrameDataLength]) {
		return egtsPcHeaderCrcError, fmt.Errorf("Не верная сумма тела пакета")
	}
	return egtsPcOk, err
}

// Encode кодирует струткуру в байтовую строку
func (p *Package) Encode() ([]byte, error) {
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

	if p.HeaderLength == 0 {
		p.HeaderLength = DEFAULT_HEADER_LEN
		if p.Route == "1" {
			p.HeaderLength += 5
		}
	}

	if err = buf.WriteByte(p.HeaderLength); err != nil {
		return result, fmt.Errorf("Не удалось записать длину заголовка: %v", err)
	}

	if err = buf.WriteByte(p.HeaderEncoding); err != nil {
		return result, fmt.Errorf("Не удалось записать метод кодирования: %v", err)
	}

	var sfrd []byte
	if p.ServicesFrameData != nil {
		sfrd, err = p.ServicesFrameData.Encode()
		if err != nil {
			return result, err
		}
	}
	p.FrameDataLength = uint16(len(sfrd))
	if err = binary.Write(buf, binary.LittleEndian, p.FrameDataLength); err != nil {
		return result, fmt.Errorf("Не удалось записать длину секции данных: %v", err)
	}

	if err = binary.Write(buf, binary.LittleEndian, p.PacketIdentifier); err != nil {
		return result, fmt.Errorf("Не удалось записать идентификатор пакета: %v", err)
	}

	if err = buf.WriteByte(p.PacketType); err != nil {
		return result, fmt.Errorf("Не удалось записать идентификатор пакета: %v", err)
	}

	if p.Route == "1" {
		if err = binary.Write(buf, binary.LittleEndian, p.PeerAddress); err != nil {
			return result, fmt.Errorf("Не удалось записать адрес апк отправителя: %v", err)
		}

		if err = binary.Write(buf, binary.LittleEndian, p.RecipientAddress); err != nil {
			return result, fmt.Errorf("Не удалось записать адрес апк получателя: %v", err)
		}

		if err = buf.WriteByte(p.TimeToLive); err != nil {
			return result, fmt.Errorf("Не удалось записать TTL пакета: %v", err)
		}
	}

	buf.WriteByte(crc8(buf.Bytes()))

	if p.FrameDataLength > 0 {
		buf.Write(sfrd)
		if err := binary.Write(buf, binary.LittleEndian, crc16(sfrd)); err != nil {
			return result, fmt.Errorf("Не удалось записать crc16 пакета: %v", err)
		}
	}

	result = buf.Bytes()
	return result, err
}

//ToBytes переводит пакет в json
func (p *Package) ToBytes() ([]byte, error) {
	return json.Marshal(p)
}
