package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//START OMIT
func Greet(name string, response_chan chan string) {
	secs := rand.Intn(10)
	log.Printf("Sleeping %d seconds before greeting %s", secs, name)
	time.Sleep(time.Duration(secs) * time.Second)
	greeting := fmt.Sprintf("Greetings, %s!", name)
	response_chan <- greeting
	close(response_chan)
}

func main() {

	ac := make(chan string)
	bc := make(chan string)
	go Greet("Alice", ac)
	go Greet("Bob", bc)
	for {
		select {
		case greeting, ok := <-ac:
			if !ok {
				ac = nil
			}
			log.Print(greeting)
		case greeting, ok := <-bc:
			if !ok {
				bc = nil
			}
			log.Print(greeting)
		}

		if ac == nil && bc == nil {
			break
		}
	}
}

//END OMIT
