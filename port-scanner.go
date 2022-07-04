package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var hostname string
var start_port, end_port, open_ports int
var wg sync.WaitGroup

func main() {
	init_input()
}

func init_input() {
	fmt.Print("Host Target > ")
	fmt.Scan(&hostname)
	fmt.Print("Starting Port to Scan (i.e. 80) > ")
	fmt.Scan(&start_port)
	fmt.Print("End Port to Scan (i.e. 56000) > ")
	fmt.Scan(&end_port)
	fmt.Println("Running scan on " + hostname + " ...")

	port_range := end_port - start_port
	end_port_set_1 := (port_range / 10) + start_port
	end_port_set_2 := (port_range / 10) + end_port_set_1
	end_port_set_3 := (port_range / 10) + end_port_set_2
	end_port_set_4 := (port_range / 10) + end_port_set_3
	end_port_set_5 := (port_range / 10) + end_port_set_4
	end_port_set_6 := (port_range / 10) + end_port_set_5
	end_port_set_7 := (port_range / 10) + end_port_set_6
	end_port_set_8 := (port_range / 10) + end_port_set_7
	end_port_set_9 := (port_range / 10) + end_port_set_8

	wg.Add(10)

	go scan(hostname, start_port, end_port_set_1)
	go scan(hostname, (end_port_set_1 + 1), end_port_set_2)
	go scan(hostname, (end_port_set_2 + 1), end_port_set_3)
	go scan(hostname, (end_port_set_3 + 1), end_port_set_4)
	go scan(hostname, (end_port_set_4 + 1), end_port_set_5)
	go scan(hostname, (end_port_set_5 + 1), end_port_set_6)
	go scan(hostname, (end_port_set_6 + 1), end_port_set_7)
	go scan(hostname, (end_port_set_7 + 1), end_port_set_8)
	go scan(hostname, (end_port_set_8 + 1), end_port_set_9)
	go scan(hostname, (end_port_set_9 + 1), end_port)

	wg.Wait()
	fmt.Println("")
	fmt.Println("===================================")
	fmt.Printf("Total %d port is Open.\n", open_ports)
	fmt.Println("===================================")

}

func scan(host string, start_port int, end_port int) {
	for i := start_port; i <= end_port; i++ {
		target := fmt.Sprintf("%s:%d", hostname, i)

		_, err := net.DialTimeout("tcp", target, 5*time.Second)

		if err != nil {
			continue
		}

		fmt.Printf("Port %d is Open.\n", i)
		open_ports += 1
	}

	wg.Done()
}
