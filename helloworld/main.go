package main

import (
	"log"
	"webserver"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	webserver.Start_webserver()
}
