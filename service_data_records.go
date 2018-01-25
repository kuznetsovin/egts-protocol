package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RecordDataSet []RecordData

func (f *RecordDataSet) ToBytes() ([]byte, error) {
	var result []byte
	buf := new(bytes.Buffer)

	for _, rd := range *f {
		rec, err := rd.ToBytes()
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

	if recBytes, err := f.ToBytes(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}

type ServiceDataSet []ServiceDataRecord

func (f *ServiceDataSet) ToBytes() ([]byte, error) {
	var result []byte
	buf := new(bytes.Buffer)

	for _, rd := range *f {
		rec, err := rd.ToBytes()
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

	if recBytes, err := f.ToBytes(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}

type ServiceDataRecord struct {
	// Параметр определяет размер данных из поля RD (Record Data)
	RecordLength uint16

	// Номер записи. Значения в данном поле изменяются по правилам циклического
	// счётчика в диапазоне от 0 до 65535, т.е. при достижении значения 65535,
	// следующее значение должно быть 0. Значение из данного поля используется
	// для подтверждения записи
	RecordNumber uint16

	// Битовый флаг, определяющий расположение Сервиса-отправителя
	// 1 = Сервис-отправитель расположен на стороне АТ
	// 0 = Сервис- отправитель расположен на ТП
	SourceServiceOnDevice byte

	// Битовый флаг, определяющий расположение Сервиса-получателя
	// 1 = Сервис-получатель расположен на стороне АТ
	// 0 = Сервис-получатель расположен на ТП
	RecipientServiceOnDevice byte

	// Битовый флаг, определяющий принадлежность передаваемых данных определённой группе, идентификатор которой указан в поле OID
	// 1 = данные предназначены для группы
	// 0 = принадлежность группе отсутствует
	Group byte

	// Битовое поле, определяющее приоритет обработки данной записи Сервисом
	// 00 – наивысший 01 – высокий
	// 10 – средний
	// 11 – низкий
	RecordProcessingPriority byte

	// Битовое поле, определяющее наличие в данном пакете поля TM
	// 0 = поле TM отсутствует
	// 1 = поле TM присутствует
	TimeFieldExists byte

	// Битовое поле, определяющее наличие в данном пакете поля EVID
	// 1 = поле EVID присутствует
	// 0 = поле EVID отсутствует
	EventIDFieldExists byte

	// Битовое поле, определяющее наличие в данном пакете поля OID
	// 1 = поле OID присутствует
	// 0 = поле OID отсутствует
	ObjectIDFieldExists byte

	// Идентификатор объекта, сгенерировавшего данную запись,
	// или для которого данная запись предназначена (уникальный идентификатор АТ),
	// либо идентификатор группы (при GRP=1). При передаче от АТ в одном
	// Пакете Транспортного Уровня нескольких записей подряд
	// для разных сервисов, но от одного и того же объекта,
	// поле OID может присутствовать только в первой записи,
	// а в последующих записях может быть опущено.
	ObjectIdentifier uint32

	// Уникальный идентификатор события. Поле EVID задаёт некий глобальный
	// идентификатор события и применяется, когда необходимо логически связать
	// с одним единственным событием набор нескольких информационных сущностей,
	// причём сами сущности могут быть разнесены как по разным информационным
	// пакетам, так и по времени. При этом прикладное ПО имеет возможность
	// объединить все эти сущности воедино в момент представления пользователю
	// информации о событии. Например, если с нажатием тревожной кнопки связывается
	// серия фотоснимков, поле EVID должно указываться в каждой сервисной записи,
	// связанной с этим событием на протяжении передачи всех сущностей, связанных
	// с данным событием, как бы долго не длилась передача всего пула информации.
	EventIdentifier uint32

	// Время формирования записи на стороне Отправителя (секунды с 00:00:00 01.01.2010 UTC).
	// Если в одном Пакете Транспортного Уровня передаются несколько записей,
	// относящихся к одному объекту и моменту времени, то поле метки времени
	// TM может передаваться только в составе первой записи.
	Time uint32

	// Идентификатор тип Сервиса-отправителя, сгенерировавшего данную запись.
	// Например, Сервис, обрабатывающий навигационные данные на стороне АТ,
	// Сервис команд на стороне ТП и т.д.
	SourceServiceType byte

	// Идентификатор тип Сервиса-получателя данной записи.
	// Например, Сервис, обрабатывающий навигационные данные на стороне ТП,
	// Сервис обработки команд на стороне АТ и т.д.
	RecipientServiceType byte

	// Поле, содержащее информацию, присущую определённому типу
	// Сервиса (одну или несколько подзаписей Сервиса типа,
	// указанного в поле SST или RST, в зависимости от вида
	// передаваемой информации).
	RecordDataSet
}

func (sdr *ServiceDataRecord) ToBytes() ([]byte, error) {
	var result []byte
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, sdr.RecordLength); err != nil {
		return result, err
	}

	if err := binary.Write(buf, binary.LittleEndian, sdr.RecordNumber); err != nil {
		return result, err
	}

	//составной байт
	flagsByte := fmt.Sprintf(
		"%b%b%b%02b%b%b%b",
		sdr.SourceServiceOnDevice,
		sdr.RecipientServiceOnDevice,
		sdr.Group,
		sdr.RecordProcessingPriority,
		sdr.TimeFieldExists,
		sdr.EventIDFieldExists,
		sdr.ObjectIDFieldExists,
	)
	flagByte, err := bitsToByte(flagsByte)
	if err != nil {
		return result, err
	}
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

	rd, err := sdr.RecordDataSet.ToBytes()
	if err != nil {
		return result, err
	}
	buf.Write(rd)

	result = buf.Bytes()
	return result, nil
}

type EGTS_PT_APPDATA struct {
	// Структуры, содержащие информацию Протокола Уровня Поддержки Услуг.
	// Таких структур в составе поля ServicesFrameData пакета Транспортного Уровня может быть одна или несколько,
	// идущих одна за другой. Описание внутреннего состава структур представлено в
	// документе “Терминал ЭРА ГЛОНАСС, Протокол Обмена Данными, Уровень Поддержки Услуг” и
	// перечне спецификаций на отдельные сервисы.
	ServiceDataSet
}

func (ad *EGTS_PT_APPDATA) ToBytes() ([]byte, error) {
	return ad.ServiceDataSet.ToBytes()
}

func (ad *EGTS_PT_APPDATA) Length() uint16 {
	var result uint16

	if recBytes, err := ad.ToBytes(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}
	return result
}
