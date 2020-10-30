package egts

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

//SrTermIdentity структура подзаписи типа EGTS_SR_TERM_IDENTITY, которая используется АС при запросе
//авторизации на телематическую платформу и содержит учетные данные АС.
type SrTermIdentity struct {
	TerminalIdentifier       uint32 `json:"TID"`
	MNE                      string `json:"MNE"`
	BSE                      string `json:"BSE"`
	NIDE                     string `json:"NIDE"`
	SSRA                     string `json:"SSRA"`
	LNGCE                    string `json:"LNGCE"`
	IMSIE                    string `json:"IMSIE"`
	IMEIE                    string `json:"IMEIE"`
	HDIDE                    string `json:"HDIDE"`
	HomeDispatcherIdentifier uint16 `json:"HDID"`
	IMEI                     string `json:"IMEI"`
	IMSI                     string `json:"IMSI"`
	LanguageCode             string `json:"LNGC"`
	NetworkIdentifier        []byte `json:"NID"`
	BufferSize               uint16 `json:"BS"`
	MobileNumber             string `json:"MSISDN"`
}

//Decode разбирает байты в структуру подзаписи
func (e *SrTermIdentity) Decode(content []byte) error {
	var (
		err   error
		flags byte
	)
	buf := bytes.NewReader(content)

	tmpBuf := make([]byte, 4)
	if _, err = buf.Read(tmpBuf); err != nil {
		return fmt.Errorf("Не удалось получить идентификатор терминал при авторизации")
	}
	e.TerminalIdentifier = binary.LittleEndian.Uint32(tmpBuf)

	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось считать байт флагов term identify: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)
	e.MNE = flagBits[:1]
	e.BSE = flagBits[1:2]
	e.NIDE = flagBits[2:3]
	e.SSRA = flagBits[3:4]
	e.LNGCE = flagBits[4:5]
	e.IMSIE = flagBits[5:6]
	e.IMEIE = flagBits[6:7]
	e.HDIDE = flagBits[7:]

	if e.HDIDE == "1" {
		tmpBuf = make([]byte, 2)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить идентификатор «домашней» телематической платформы при авторизации")
		}
		e.HomeDispatcherIdentifier = binary.LittleEndian.Uint16(tmpBuf)

	}

	if e.IMEIE == "1" {
		tmpBuf = make([]byte, 15)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить IMEI при авторизации")
		}
		e.IMEI = string(tmpBuf)
	}

	if e.IMSIE == "1" {
		tmpBuf = make([]byte, 16)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить IMSI при авторизации")
		}
		e.IMSI = string(tmpBuf)
	}

	if e.LNGCE == "1" {
		tmpBuf = make([]byte, 3)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить код языка при авторизации")
		}
		e.LanguageCode = string(tmpBuf)
	}

	if e.NIDE == "1" {
		e.NetworkIdentifier = make([]byte, 3)
		if _, err = buf.Read(e.NetworkIdentifier); err != nil {
			return fmt.Errorf("Не удалось получить код идентификатор сети оператора при авторизации")
		}
	}

	if e.BSE == "1" {
		tmpBuf = make([]byte, 2)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить максимальный размер буфера при авторизации")
		}
		e.BufferSize = binary.LittleEndian.Uint16(tmpBuf)
	}

	if e.MNE == "1" {
		tmpBuf = make([]byte, 15)
		if _, err = buf.Read(tmpBuf); err != nil {
			return fmt.Errorf("Не удалось получить телефонный номер мобильного абонента")
		}
		e.MobileNumber = string(tmpBuf)
	}

	return err
}

//Encode преобразовывает подзапись в набор байт
func (e *SrTermIdentity) Encode() ([]byte, error) {
	var (
		result []byte
		flags  uint64
		err    error
	)
	buf := new(bytes.Buffer)

	if err = binary.Write(buf, binary.LittleEndian, e.TerminalIdentifier); err != nil {
		return result, fmt.Errorf("Не удалось записать идентификатор терминал при авторизации")
	}

	flags, err = strconv.ParseUint(e.MNE+e.BSE+e.NIDE+e.SSRA+e.LNGCE+e.IMSIE+e.IMEIE+e.HDIDE, 2, 8)
	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать байт флагов term identify: %v", err)
	}

	if e.HDIDE == "1" {
		if err = binary.Write(buf, binary.LittleEndian, e.HomeDispatcherIdentifier); err != nil {
			return result, fmt.Errorf("Не удалось записать идентификатор «домашней» телематической платформы при авторизации")
		}
	}

	if e.IMEIE == "1" {
		if _, err = buf.Write([]byte(e.IMEI)); err != nil {
			return result, fmt.Errorf("Не удалось записать IMEI при авторизации")
		}
	}

	if e.IMSIE == "1" {
		if _, err = buf.Write([]byte(e.IMSI)); err != nil {
			return result, fmt.Errorf("Не удалось записать IMSI при авторизации")
		}
	}

	if e.LNGCE == "1" {
		if _, err = buf.Write([]byte(e.LanguageCode)); err != nil {
			return result, fmt.Errorf("Не удалось записать IMSI при авторизации")
		}
	}

	if e.NIDE == "1" {
		if _, err = buf.Write(e.NetworkIdentifier); err != nil {
			return result, fmt.Errorf("Не удалось записать код идентификатор сети оператора при авторизации")
		}
	}

	if e.BSE == "1" {
		if err = binary.Write(buf, binary.LittleEndian, e.BufferSize); err != nil {
			return result, fmt.Errorf("Не удалось записать максимальный размер буфера при авторизации")
		}
	}

	if e.MNE == "1" {
		if _, err = buf.Write([]byte(e.MobileNumber)); err != nil {
			return result, fmt.Errorf("Не удалось записать телефонный номер мобильного абонента")
		}
	}

	result = buf.Bytes()
	return result, err
}

//Length получает длинну закодированной подзаписи
func (e *SrTermIdentity) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
