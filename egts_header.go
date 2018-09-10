package main

import (
	"bytes"
	"fmt"
	"encoding/binary"
	"strconv"
)

type EgtsHeader struct {
	ProtocolVersion  byte
	SecurityKeyID    byte
	Prefix           string
	Route            string
	EncryptionAlg    string
	Compression      string
	Priority         string
	HeaderLength     byte
	HeaderEncoding   byte
	FrameDataLength  uint16
	PacketIdentifier uint16
	PacketType       byte
	PeerAddress      uint16
	RecipientAddress uint16
	TimeToLive       byte
	HeaderCheckSum   byte
}

//Decode декодирует байтовую строку в заголовк
func (eh *EgtsHeader) Decode(header []byte) error {
	var (
		err   error
		flags byte
	)
	buf := bytes.NewReader(header)
	if eh.ProtocolVersion, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить версию протокола: %v\n", err)
	}

	if eh.SecurityKeyID, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить идентификатор ключа: %v\n", err)
	}

	//разбираем флаги
	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось флаги: %v\n", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)
	eh.Prefix = flagBits[:2]         // flags << 7, flags << 6
	eh.Route = flagBits[2:3]         // flags << 5
	eh.EncryptionAlg = flagBits[3:5] // flags << 4, flags << 3
	eh.Compression = flagBits[5:6]   // flags << 2
	eh.Priority = flagBits[6:]       // flags << 1, flags << 0

	if eh.HeaderLength, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить длину заголовка: %v\n", err)
	}

	if eh.HeaderEncoding, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить метод кодирования: %v\n", err)
	}

	tmpIntBuf := make([]byte, 2)
	if _, err = buf.Read(tmpIntBuf); err != nil {
		return fmt.Errorf("Не удалось получить длину секции данных: %v\n", err)
	}
	eh.FrameDataLength = binary.LittleEndian.Uint16(tmpIntBuf)

	if _, err = buf.Read(tmpIntBuf); err != nil {
		return fmt.Errorf("Не удалось получить идентификатор пакета: %v\n", err)
	}
	eh.PacketIdentifier = binary.LittleEndian.Uint16(tmpIntBuf)

	if eh.PacketType, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить идентификатор пакета: %v\n", err)
	}

	if eh.Route == "1" {
		if _, err = buf.Read(tmpIntBuf); err != nil {
			return fmt.Errorf("Не удалось получить адрес апк отправителя: %v\n", err)
		}
		eh.PeerAddress = binary.LittleEndian.Uint16(tmpIntBuf)

		if _, err = buf.Read(tmpIntBuf); err != nil {
			return fmt.Errorf("Не удалось получить адрес апк получателя: %v\n", err)
		}
		eh.RecipientAddress = binary.LittleEndian.Uint16(tmpIntBuf)

		if eh.TimeToLive, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Не удалось получить TTL пакета: %v\n", err)
		}
	}

	if eh.HeaderCheckSum, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить crc: %v\n", err)
	}

	return err
}

//Encode кодирует заколовок в байтовую строку
func (eh *EgtsHeader) Encode() ([]byte, error) {
	var (
		header []byte
		err    error
		flags  uint64
	)
	buf := new(bytes.Buffer)

	if err = buf.WriteByte(eh.ProtocolVersion); err != nil {
		return nil, fmt.Errorf("Не удалось записать версию протокола: %v\n", err)
	}
	if err = buf.WriteByte(eh.SecurityKeyID); err != nil {
		return nil, fmt.Errorf("Не удалось записать  идентификатор ключа: %v\n", err)
	}

	//собираем флаги
	flagsBits := eh.Prefix + eh.Prefix + eh.Route + eh.EncryptionAlg + eh.Compression + eh.Priority
	if flags, err = strconv.ParseUint(flagsBits, 2, 8); err != nil {
		return nil, fmt.Errorf("Не удалось сгенерировать байт флагов: %v\n", err)
	}

	if err = buf.WriteByte(uint8(flags)); err != nil {
		return nil, fmt.Errorf("Не удалось записать флаги: %v\n", err)
	}

	if err = buf.WriteByte(eh.HeaderLength); err != nil {
		return nil, fmt.Errorf("Не удалось записать длину заголовка: %v\n", err)
	}

	if err = buf.WriteByte(eh.HeaderEncoding); err != nil {
		return nil, fmt.Errorf("Не удалось записать метод кодирования: %v\n", err)
	}

	tmpIntBuf := make([]byte, 2)
	binary.LittleEndian.PutUint16(tmpIntBuf, eh.FrameDataLength)
	if _, err = buf.Write(tmpIntBuf); err != nil {
		return nil, fmt.Errorf("Не удалось записать длину секции данных: %v\n", err)
	}

	binary.LittleEndian.PutUint16(tmpIntBuf, eh.PacketIdentifier)
	if _, err = buf.Write(tmpIntBuf); err != nil {
		return nil, fmt.Errorf("Не удалось записать идентификатор пакета: %v", err)
	}

	if err = buf.WriteByte(eh.PacketType); err != nil {
		return nil, fmt.Errorf("Не удалось записать идентификатор пакета: %v\n", err)
	}

	if eh.Route == "1" {
		binary.LittleEndian.PutUint16(tmpIntBuf, eh.PeerAddress)
		if _, err = buf.Write(tmpIntBuf); err != nil {
			return nil, fmt.Errorf("Не удалось записать адрес апк отправителя: %v\n", err)
		}

		binary.LittleEndian.PutUint16(tmpIntBuf, eh.RecipientAddress)
		if _, err = buf.Write(tmpIntBuf); err != nil {
			return nil, fmt.Errorf("Не удалось записать адрес апк получателя: %v\n", err)
		}

		if err = buf.WriteByte(eh.TimeToLive); err != nil {
			return nil, fmt.Errorf("Не удалось записать TTL пакета: %v\n", err)
		}
	}

	header = buf.Bytes()
	header = append(header, crc8(header))

	return header, err
}

