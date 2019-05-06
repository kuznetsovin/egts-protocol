package main

import (
	"github.com/kuznetsovin/libegts"
	"sync/atomic"
)

var (
	pidCounter uint32 = 0
	rnCounter  uint32 = 0
)

func getNextPid() uint16 {
	if pidCounter < 65535 {
		atomic.AddUint32(&pidCounter, 1)
	} else {
		pidCounter = 0
	}
	return uint16(atomic.LoadUint32(&pidCounter))
}

func getNextRN() uint16 {
	if rnCounter < 65535 {
		atomic.AddUint32(&rnCounter, 1)
	} else {
		rnCounter = 0
	}
	return uint16(atomic.LoadUint32(&rnCounter))
}

func createPtResponse(p *egts.Package, resultCode, serviceType uint8, srResponses egts.RecordDataSet) ([]byte, error) {
	respSection := egts.PtResponse{
		ResponsePacketID: p.PacketIdentifier,
		ProcessingResult: resultCode,
	}

	if srResponses != nil {
		respSection.SDR = &egts.ServiceDataSet{
			egts.ServiceDataRecord{
				RecordLength:             srResponses.Length(),
				RecordNumber:             getNextRN(),
				SourceServiceOnDevice:    "0",
				RecipientServiceOnDevice: "0",
				Group:                    "1",
				RecordProcessingPriority: "00",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "0",
				SourceServiceType:        serviceType,
				RecipientServiceType:     serviceType,
				RecordDataSet:            srResponses,
			},
		}
	}

	respPkg := egts.Package{
		ProtocolVersion:   1,
		SecurityKeyID:     0,
		Prefix:            "00",
		Route:             "0",
		EncryptionAlg:     "00",
		Compression:       "0",
		Priority:          "00",
		HeaderLength:      11,
		HeaderEncoding:    0,
		FrameDataLength:   respSection.Length(),
		PacketIdentifier:  getNextPid(),
		PacketType:        egts.PtResponsePacket,
		ServicesFrameData: &respSection,
	}

	return respPkg.Encode()
}

func createSrResultCode(p *egts.Package, resultCode uint8) ([]byte, error) {
	rds := egts.RecordDataSet{
		egts.RecordData{
			SubrecordType:   egts.SrResultCodeType,
			SubrecordLength: uint16(1),
			SubrecordData: &egts.SrResultCode{
				ResultCode: resultCode,
			},
		},
	}

	sfd := egts.ServiceDataSet{
		egts.ServiceDataRecord{
			RecordLength:             rds.Length(),
			RecordNumber:             getNextRN(),
			SourceServiceOnDevice:    "0",
			RecipientServiceOnDevice: "0",
			Group:                    "1",
			RecordProcessingPriority: "00",
			TimeFieldExists:          "0",
			EventIDFieldExists:       "0",
			ObjectIDFieldExists:      "0",
			SourceServiceType:        egts.AuthService,
			RecipientServiceType:     egts.AuthService,
			RecordDataSet:            rds,
		},
	}

	respPkg := egts.Package{
		ProtocolVersion:   1,
		SecurityKeyID:     0,
		Prefix:            "00",
		Route:             "0",
		EncryptionAlg:     "00",
		Compression:       "0",
		Priority:          "00",
		HeaderLength:      11,
		HeaderEncoding:    0,
		FrameDataLength:   sfd.Length(),
		PacketIdentifier:  getNextPid(),
		PacketType:        egts.PtResponsePacket,
		ServicesFrameData: &sfd,
	}

	return respPkg.Encode()
}
