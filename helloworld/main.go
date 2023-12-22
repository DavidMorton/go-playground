package main

import (
	"log"
	"webserver/counting_path"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	counting_path.Run_webserver()
}
