package tcpcheck

import (
	"net"
	"strconv"
	"sync"
	"time"
)

func worker(to_check Data, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we're done.
	defer wg.Done()

	// fmt.Printf("Worker %s starting\n", to_check.Name)

	address := net.JoinHostPort(to_check.Host, strconv.Itoa(to_check.Port))
	// fmt.Printf("Connect to address: %-30s ", address)

	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		UpdateAvailability(to_check.Uuid, 0)
	} else {
		defer conn.Close()
		UpdateAvailability(to_check.Uuid, 1)
	}
	// fmt.Printf("Worker %s done\n", to_check.Name)
}

func CheckAll() {
	// fmt.Printf("CheckAll\n")
	// fmt.Printf("%v\n", Get())

	var wg sync.WaitGroup
	for _, tc := range Checks {
		wg.Add(1)
		go worker(tc, &wg)
	}
	wg.Wait()
}
