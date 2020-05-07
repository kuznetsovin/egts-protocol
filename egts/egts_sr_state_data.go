package egts

import (
	"bytes"
	"fmt"
	"strconv"
)

//SrStateData структура подзаписи типа EGTS_SR_STATE_DATA, которая используется для передачи на
//аппаратно-программный комплекс информации о состоянии абонентского терминала  (текущий режим работы,
//напряжение основного и резервного источников питания и т.д.)
type SrStateData struct {
	State                  uint8  `json:"ST"`
	MainPowerSourceVoltage uint8  `json:"MPSV"`
	BackUpBatteryVoltage   uint8  `json:"BBV"`
	InternalBatteryVoltage uint8  `json:"IBV"`
	NMS                    string `json:"NMS"`
	IBU                    string `json:"IBU"`
	BBU                    string `json:"BBU"`
}

//Decode разбирает байты в структуру подзаписи
func (e *SrStateData) Decode(content []byte) error {
	var (
		err   error
		flags byte
	)

	buf := bytes.NewReader(content)
	if e.State, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить текущий режим работы: %v", err)
	}

	if e.MainPowerSourceVoltage, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить значение напряжения основного источника питания: %v", err)
	}

	if e.BackUpBatteryVoltage, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить значение напряжения резервной батареи: %v", err)
	}

	if e.InternalBatteryVoltage, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить значение напряжения внутренней батареи: %v", err)
	}

	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить байт флагов state_data: %v", err)
	}
	flagBits := fmt.Sprintf("%08b", flags)
	e.NMS = flagBits[5:6]
	e.IBU = flagBits[6:7]
	e.BBU = flagBits[7:]

	return err
}

//Encode преобразовывает подзапись в набор байт
func (e *SrStateData) Encode() ([]byte, error) {
	var (
		err    error
		flags  uint64
		result []byte
	)
	buf := new(bytes.Buffer)

	if err = buf.WriteByte(e.State); err != nil {
		return result, fmt.Errorf("Не удалось записать текущий режим работы: %v", err)
	}

	if err = buf.WriteByte(e.MainPowerSourceVoltage); err != nil {
		return result, fmt.Errorf("Не удалось записать значение напряжения основного источника питания: %v", err)
	}

	if err = buf.WriteByte(e.BackUpBatteryVoltage); err != nil {
		return result, fmt.Errorf("Не удалось записать значение напряжения резервной батареи: %v", err)
	}

	if err = buf.WriteByte(e.InternalBatteryVoltage); err != nil {
		return result, fmt.Errorf("Не удалось записать значение напряжения внутренней батареи: %v", err)
	}

	if flags, err = strconv.ParseUint("00000"+e.NMS+e.IBU+e.BBU, 2, 8); err != nil {
		return result, fmt.Errorf("Не удалось сгенерировать байт флагов state_data: %v", err)
	}

	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать байт флагов state_data: %v", err)
	}

	result = buf.Bytes()

	return result, err
}

//Length получает длинну закодированной подзаписи
func (e *SrStateData) Length() uint16 {
	var result uint16

	if recBytes, err := e.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
