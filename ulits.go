package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
)

var (
	pidCounter uint32 = 0
	rnCounter  uint32 = 0
)

func printDecodePackage(bytesPkg []byte) string {
	pkg := EgtsPackage{}

	_, err := pkg.Decode(bytesPkg)
	if err != nil {
		return fmt.Sprintf("Не удалось расшифровать пакет:\n %v\n", err)
	}

	jsonPkg, err := json.MarshalIndent(pkg, "", "    ")
	if err != nil {
		return fmt.Sprintf("Не сформировать отладочный json:\n %v\n", err)
	}

	return string(jsonPkg)
}

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
