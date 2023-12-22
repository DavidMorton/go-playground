package main

import (
	"log"
	"webserver/drawing"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	drawing.Main()
}
