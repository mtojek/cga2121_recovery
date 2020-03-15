package main

import (
	"flag"
	"log"
)

func main() {
	hostPort := flag.String("hostPort", "192.168.0.1", "host and port")
	user := flag.String("user", "admin", "user")
	passwordMinLength := flag.Int("passwordMinLength", 1, "minimal password length")
	passwordPattern := flag.String("passwordPattern", "aA1_", "character patterns")
	numWorkers := flag.Int("numWorkers", 10, "number of workers")
	flag.Parse()

	seq, err := newSequence(*passwordMinLength, *passwordPattern)
	if err != nil {
		log.Fatal(err)
	}

	bf := newBruteForce(*hostPort, *user, *numWorkers, seq)
	bf.Do()
}
