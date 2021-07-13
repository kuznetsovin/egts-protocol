package egts

// BinaryData интерфейс для работы с бинарными секциями
type BinaryData interface {
	Decode([]byte) error
	Encode() ([]byte, error)
	Length() uint16
}
