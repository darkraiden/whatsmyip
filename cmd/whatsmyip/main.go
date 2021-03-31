package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darkraiden/whatsmyip"
)

func main() {
	ip, err := whatsmyip.Get(&http.Client{})
	if err != nil {
		log.Fatalf("Failed to fetch the public IP address: %s", ip)
	}
	fmt.Println(ip)
}
