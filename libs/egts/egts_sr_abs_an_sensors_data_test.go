package egts

import (
	"bytes"
	"testing"
)

func TestEgtsSrAbsAnSensData_Encode(t *testing.T) {
	a := SrAbsAnSensData{
		SensorNumber: 0x98,
		Value:        0x123456,
	}
	data, err := a.Encode()
	if err != nil {
		t.Errorf("Ошибка кодирования: %v", err)
	}

	if data[0] != 0x98 ||
		data[1] != 0x56 ||
		data[2] != 0x34 ||
		data[3] != 0x12 {
		t.Error("Ошибка кодирования")
	}

	b := SrAbsAnSensData{}

	if err := b.Decode(data); err != nil {
		t.Error("Ошибка декодирования")
	}

	if a.Value != b.Value || a.SensorNumber != b.SensorNumber {
		t.Error("Ошибка кодирования")
	}

}
func TestEgtsSrAbsAnSensData_Decode(t *testing.T) {
	data := []byte{0x98, 0x56, 0x34, 0x12}
	a := SrAbsAnSensData{}
	if err := a.Decode(data); err != nil {
		t.Errorf("Ошибка декодирования: %v", err)
	}

	if a.SensorNumber != 0x98 || a.Value != 0x123456 {
		t.Error("Ошибка декодирования")
	}

	data2, err := a.Encode()
	if err != nil {
		t.Error("Ошибка кодирования")
	}

	if bytes.Compare(data, data2) != 0 {
		t.Error("Ошибка декодирования")
	}
}
