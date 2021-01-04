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

func DeleteByUuid(uuid string) {
	for idx, tc := range Checks {
		if tc.Uuid == uuid {
			Checks = append(Checks[:idx], Checks[idx+1:]...)
			return
		}
	}
}

func UpdateAvailability(idx int, available int) {
	Mutex.Lock()
	Checks[idx].Available = available
	Checks[idx].LastCheck = time.Now()
	Mutex.Unlock()
}
