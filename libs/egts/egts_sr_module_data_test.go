package egts

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testSrModuleData = SrModuleData{
		ModuleType:      1,
		VendorID:        0,
		FirmwareVersion: 257,
		SoftwareVersion: 259,
		Modification:    1,
		State:           1,
		SerialNumber:    "TB1011010000001022023",
		Description:     "rev170122.A",
	}
	//nolint:lll
	testSrModuleDataBytes = []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x3, 0x1, 0x1, 0x1, 0x54, 0x42, 0x31, 0x30, 0x31, 0x31, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x32, 0x32, 0x30, 0x32, 0x33, 0x0, 0x72, 0x65, 0x76, 0x31, 0x37, 0x30, 0x31, 0x32, 0x32, 0x2e, 0x41, 0x0}
)

// проверяем, что рекордсет работает правильно с данным типом подзаписи
//nolint:dupl
func TestEgtsSrModuleDataRs(t *testing.T) {
	stateDataRDBytes := append([]byte{0x02, 0x2d, 0x00}, testSrModuleDataBytes...)
	stateDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrModuleDataType,
			SubrecordLength: uint16(len(testSrModuleDataBytes)),
			SubrecordData:   &testSrModuleData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := stateDataRD.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, testBytes, stateDataRDBytes)

		if assert.NoError(t, testStruct.Decode(stateDataRDBytes)) {
			assert.Equal(t, stateDataRD, testStruct)
		}
	}
}

//nolint:funlen
func TestSrModuleData_Decode(t *testing.T) {
	type fields struct {
		ModuleType      int8
		VendorID        uint32
		FirmwareVersion uint16
		SoftwareVersion uint16
		Modification    byte
		State           byte
		SerialNumber    string
		_               byte
		Description     string
		_               byte
	}
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			fields:  fields(testSrModuleData),
			args:    args{content: testSrModuleDataBytes},
			wantErr: false,
		},
		{
			name:    "Error - ModuleType",
			fields:  fields(testSrModuleData),
			args:    args{content: []byte{}},
			wantErr: true,
		},
		{
			name:    "Error - VendorID",
			fields:  fields(testSrModuleData),
			args:    args{content: []byte{0x01}},
			wantErr: true,
		},
		{
			name:    "Error - FirmwareVersion",
			fields:  fields(testSrModuleData),
			args:    args{content: []byte{0x01, 0x0, 0x0, 0x0, 0x0}},
			wantErr: true,
		},
		{
			name:    "Error - SoftwareVersion",
			fields:  fields(testSrModuleData),
			args:    args{content: []byte{0x01, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1}},
			wantErr: true,
		},
		{
			name:    "Error - Modification",
			fields:  fields(testSrModuleData),
			args:    args{content: []byte{0x01, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x3, 0x1}},
			wantErr: true,
		},
		{
			name:    "Error - State",
			fields:  fields(testSrModuleData),
			args:    args{content: []byte{0x01, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x3, 0x1, 0x1}},
			wantErr: true,
		},
		{
			name:    "Error - SerialNumber",
			fields:  fields(testSrModuleData),
			args:    args{content: []byte{0x01, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x3, 0x1, 0x1, 0x1}},
			wantErr: true,
		},
		{
			name:   "Error - Description",
			fields: fields(testSrModuleData),
			//nolint:lll
			args:    args{content: []byte{0x01, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x3, 0x1, 0x1, 0x1, 0x54, 0x42, 0x31, 0x30, 0x31, 0x31, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x32, 0x32, 0x30, 0x32, 0x33, 0x00}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &SrModuleData{
				ModuleType:      tt.fields.ModuleType,
				VendorID:        tt.fields.VendorID,
				FirmwareVersion: tt.fields.FirmwareVersion,
				SoftwareVersion: tt.fields.SoftwareVersion,
				Modification:    tt.fields.Modification,
				State:           tt.fields.State,
				SerialNumber:    tt.fields.SerialNumber,
				Description:     tt.fields.Description,
			}
			if err := e.Decode(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("SrModuleData.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSrModuleData_Encode(t *testing.T) {
	type fields struct {
		ModuleType      int8
		VendorID        uint32
		FirmwareVersion uint16
		SoftwareVersion uint16
		Modification    byte
		State           byte
		SerialNumber    string
		_               byte
		Description     string
		_               byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "Success",
			fields:  fields(testSrModuleData),
			want:    testSrModuleDataBytes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &SrModuleData{
				ModuleType:      tt.fields.ModuleType,
				VendorID:        tt.fields.VendorID,
				FirmwareVersion: tt.fields.FirmwareVersion,
				SoftwareVersion: tt.fields.SoftwareVersion,
				Modification:    tt.fields.Modification,
				State:           tt.fields.State,
				SerialNumber:    tt.fields.SerialNumber,
				Description:     tt.fields.Description,
			}
			got, err := e.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("SrModuleData.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SrModuleData.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSrModuleData_Length(t *testing.T) {
	type fields struct {
		ModuleType      int8
		VendorID        uint32
		FirmwareVersion uint16
		SoftwareVersion uint16
		Modification    byte
		State           byte
		SerialNumber    string
		_               byte
		Description     string
		_               byte
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		{
			name:   "Success",
			fields: fields(testSrModuleData),
			want:   45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &SrModuleData{
				ModuleType:      tt.fields.ModuleType,
				VendorID:        tt.fields.VendorID,
				FirmwareVersion: tt.fields.FirmwareVersion,
				SoftwareVersion: tt.fields.SoftwareVersion,
				Modification:    tt.fields.Modification,
				State:           tt.fields.State,
				SerialNumber:    tt.fields.SerialNumber,
				Description:     tt.fields.Description,
			}
			if got := e.Length(); got != tt.want {
				t.Errorf("SrModuleData.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}
