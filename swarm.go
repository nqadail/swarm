package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"swarm/hive"
)

func Scout(url string, numRequests int) {
	for i := 0; i < numRequests; i++ {
		bumble := hive.NewBee(url, i)
		dance := bumble.FindNectar()
		fmt.Printf("Bee #%d: %d bytes read, code = %d\n",
			dance.Bee().Id(), dance.Length(), dance.Code())
	}
}

func Swarm(url string, numBees int) {
	danceParty := make(chan *hive.Dance)
	for i := 0; i < numBees; i++ {
		bumble := hive.NewBee(url, i)
		go func() { danceParty <- bumble.FindNectar() }()
	}
	for waitingFor := numBees; waitingFor > 0; waitingFor-- {
		dance := <-danceParty // blocks waiting on a dance
		fmt.Printf("Bee #%d: %d bytes read, code = %d\n",
			dance.Bee().Id(), dance.Length(), dance.Code())
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("\nusage: swarm <url> <num-bees>\n\n")
		return
	}

	url := os.Args[1]
	count, _ := strconv.Atoi(os.Args[2])

	scoutStart := time.Now()
	fmt.Printf("Sending out scouts...\n\n")
	Scout(url, count)
	fmt.Printf("\nScouting Completed in %.3fs.\n\n", time.Since(scoutStart).Seconds())

	swarmStart := time.Now()
	fmt.Printf("Sending swarm...\n\n")
	Swarm(url, count)

	fmt.Printf("\nScouting Completed in: %.3fs.", time.Since(scoutStart).Seconds())
	fmt.Printf("\nSwarm Completed in:    %.3fs.\n\n", time.Since(swarmStart).Seconds())
}
