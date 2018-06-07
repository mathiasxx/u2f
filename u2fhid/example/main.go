package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/mathiasxx/u2f/u2fhid"
)

func main() {
	msg := []byte(strings.Repeat("echo", 100))
	devices, err := u2fhid.Devices()
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range devices {
		dev, err := u2fhid.Open(d)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("opened", d.Path)
		res, err := dev.Ping([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		if !bytes.Equal(res, msg) {
			log.Fatalf("expected %x, got %x", msg, res)
		}
		log.Println("successfully pinged", d.Path)
	}
}
