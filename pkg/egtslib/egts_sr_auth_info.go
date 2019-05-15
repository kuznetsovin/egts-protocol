package egts

import (
	"bytes"
	"fmt"
	"strings"
)

//SrAuthInfo структура подзаписи типа EGTS_SR_AUTH_INFO, которая предназначена для передачи на
//телематическую платформу аутентификационных данных АС с использованием ранее переданных
//со стороны платформы параметров для осуществления шифрования данных.
type SrAuthInfo struct {
	UserName       string `json:"UNM"`
	UserPassword   string `json:"UPSW"`
	ServerSequence string `json:"SS"`
}

// Decode разбирает байты в структуру подзаписи
func (e *SrAuthInfo) Decode(content []byte) error {
	var (
		err    error
		tmpStr string
	)
	//разделитель строковых полей из ГОСТ 54619 - 2011 секции EGTS_SR_AUTH_INFO
	sep := byte(0x00)

	buf := bytes.NewBuffer(content)
	tmpStr, err = buf.ReadString(sep)
	if err != nil {
		return fmt.Errorf("Не удалось считать имя пользователя sr_auth_info: %v", err)
	}
	e.UserName = strings.TrimSuffix(tmpStr, string(sep))

	tmpStr, err = buf.ReadString(sep)
	if err != nil {
		return fmt.Errorf("Не удалось считать пароль sr_auth_info: %v", err)
	}
	e.UserPassword = strings.TrimSuffix(tmpStr, string(sep))

	if buf.Len() > 0 {
		tmpStr, err = buf.ReadString(sep)
		if err != nil {
			return fmt.Errorf("Не удалось считать SS из sr_auth_info: %v", err)
		}
		e.ServerSequence = strings.TrimSuffix(tmpStr, string(sep))
	}

	return err
}

// Encode преобразовывает подзапись в набор байт
func (e *SrAuthInfo) Encode() ([]byte, error) {
	var (
		err    error
		result []byte
	)
	//разделитель строковых полей из ГОСТ 54619 - 2011 секции EGTS_SR_AUTH_INFO
	sep := byte(0x00)

	result = append(result, []byte(e.UserName)...)
	result = append(result, sep)

	result = append(result, []byte(e.UserPassword)...)
	result = append(result, sep)

	// необязательное поле, наличие зависит от используемого алгоритма шифрования
	// специальная серверная последовательность байт, передаваемая в подзаписи EGTS_SR_AUTH_PARAMS
	if e.ServerSequence != "" {
		result = append(result, []byte(e.ServerSequence)...)
		result = append(result, sep)
	}

	return result, err
}

//Length получает длинну закодированной подзаписи
func (e *SrAuthInfo) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
