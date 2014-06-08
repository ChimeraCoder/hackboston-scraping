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
}

func main() {

	ac := make(chan string)
	bc := make(chan string)
	go Greet("Alice", ac)
	go Greet("Bob", bc)
	select {
	case greeting := <-ac:
		log.Print(greeting)
	case greeting := <-bc:
		log.Print(greeting)
	}
}

//END OMIT
