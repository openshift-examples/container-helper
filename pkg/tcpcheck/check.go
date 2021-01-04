package tcpcheck

import (
	"net"
	"strconv"
	"sync"
	"time"
)

func worker(to_check Data, idx int, wg *sync.WaitGroup) {
	// On return, notify the WaitGroup that we're done.
	defer wg.Done()

	// fmt.Printf("Worker %s starting\n", to_check.Name)

	address := net.JoinHostPort(to_check.Host, strconv.Itoa(to_check.Port))
	// fmt.Printf("Connect to address: %-30s \n", address)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	// fmt.Printf("Connected to address: %-30s \n", address)

	if err != nil {
		UpdateAvailability(idx, 0)
	} else {
		defer conn.Close()
		UpdateAvailability(idx, 1)
	}
	// fmt.Printf("Worker %s done\n", to_check.Name)
}

func CheckAll() {
	// fmt.Printf("CheckAll\n")
	// fmt.Printf("%v\n", Get())

	var wg sync.WaitGroup
	for idx, tc := range Checks {
		wg.Add(1)
		go worker(tc, idx, &wg)
	}
	wg.Wait()
}
