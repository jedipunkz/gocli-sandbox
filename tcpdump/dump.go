package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	device := "en0"
	filter := "tcp and port 443"

	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	handle, err := pcap.OpenLive(
		device, int32(0xFFFF), true, pcap.BlockForever,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Printf("%s\n", cyan(packet))
		fmt.Printf("%s", yellow(hex.Dump(packet.Data())))
	}
}
