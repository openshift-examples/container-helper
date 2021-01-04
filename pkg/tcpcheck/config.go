package tcpcheck

import (
	"sync"
	"time"
)

var Mutex = &sync.Mutex{}

type Data struct {
	Uuid      string    `json:"uuid" validate:"nonzero"`
	Name      string    `json:"uuid" validate:"nonzero"`
	Host      string    `json:"host" validate:"nonzero"`
	Port      int       `json:"port" validate:"nonzero"`
	Available int       `json:"port"`
	LastCheck time.Time `json:"last_check"`
}
type Config struct {
	Tcpchecks []Data
}
