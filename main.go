package main

import (
	"home/kimura/go-packet-capture/cap"
	"log"
)

func main() {
	err := cap.PacketCapture()
	if err != nil {
		log.Fatal(err)
	}
}
