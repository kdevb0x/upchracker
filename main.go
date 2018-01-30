package upchracker

import (
    "log"
    "golang.com/x/mobile"
    "golang.com/x/mobile/app"
    "github.com/google/gopacket/pcap"
)
 func init() {

 }

 func findDevices() {
    devices, err := pcap.FindAllDevs()
    if err != nil {
	log.Fatal
    }

    for _, dev := range

 }
