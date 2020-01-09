package egts

import (
	"testing"
)

func TestEgtsSrAbsAnSensData_Encode(t *testing.T) {
	a := SrAbsAnSensData{
		ASN: 0x98,
		ASV: 0x123456,
	}
	data, err := a.Encode()
	if err != nil {
		t.Errorf("Ошибка декодирования: %v", err)
	}

	if data[0] != 0x98 ||
		data[1] != 0x12 ||
		data[2] != 0x34 ||
		data[3] != 0x56 {
		t.Error("Ошибка декодирования")
	}

}
func TestEgtsSrAbsAnSensData_Decode(t *testing.T) {
	data := []byte{0x98, 0x12, 0x34, 0x56}
	a := SrAbsAnSensData{}
	if err := a.Decode(data); err != nil {
		t.Errorf("Ошибка кодирования: %v", err)
	}

	if a.ASN != 0x98 || a.ASV != 0x123456 {
		t.Error("Ошибка декодирования")
	}
}
