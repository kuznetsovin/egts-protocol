package storage

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type LogConnector struct{}

func (c LogConnector) Init(cfg map[string]string) error {
	return nil
}

func (c LogConnector) Save(msg interface{ ToBytes() ([]byte, error) }) error {
	jsonPkg, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		return err
	}

	log.WithField("packet", string(jsonPkg)).Info("Export packet")
	return nil
}

func (c LogConnector) Close() error {
	return nil
}
