package tcpcheck

import (
	"time"

	guuid "github.com/google/uuid"
)

var Checks []Data

func Add(check Data) {
	check.Uuid = guuid.New().String()
	check.Available = -1
	Checks = append(Checks, check)
}

func Get() []Data {
	return Checks
}

func GetByUuid(uuid string) Data {
	for _, tc := range Checks {
		if tc.Uuid == uuid {
			return tc
		}
	}
	return Data{}
}

func Delete(uuid string) {
	for idx, tc := range Checks {
		if tc.Uuid == uuid {
			Checks = append(Checks[:idx], Checks[idx+1:]...)
			return
		}
	}
}

func UpdateAvailability(uuid string, available int) {
	tc := GetByUuid(uuid)
	tc.mu.Lock()
	Delete(uuid)
	tc.Available = available
	tc.LastCheck = time.Now()
	Checks = append(Checks, tc)
	tc.mu.Unlock()
}
