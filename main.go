package main

import (
	"log"
	"net"
	"strings"
	"time"

	"github.com/williamfhe/godivert"
)

var fbIps []net.IP
var notFbIps []net.IP

func checkPacket(wd *godivert.WinDivertHandle, packetChan <-chan *godivert.Packet) {
	for packet := range packetChan {

		if isFacebookIP(packet.DstIP()) || isFacebookIP(packet.SrcIP()) {
			log.Println("Packet artificially slowed down")
			time.Sleep(time.Millisecond * 255)
		}
		packet.Send(wd)
	}
}

func isFacebookIP(ip net.IP) bool {
	for _, checkedIP := range fbIps {
		if ip.Equal(checkedIP) {
			return true
		}
	}

	for _, checkedIP := range notFbIps {
		if ip.Equal(checkedIP) {
			return false
		}
	}

	hosts, err := net.LookupAddr(ip.String())
	if err != nil {
		return false
	}

	for _, host := range hosts {
		if strings.Contains(host, "facebook.com") {
			fbIps = append(fbIps, ip)
			return true
		}
	}
	notFbIps = append(notFbIps, ip)
	return false
}

func main() {
	winDivert, err := godivert.NewWinDivertHandle("tcp.DstPort == 443")
	if err != nil {
		panic(err)
	}
	defer winDivert.Close()

	packetChan, err := winDivert.Packets()
	if err != nil {
		panic(err)
	}

	go checkPacket(winDivert, packetChan)

	block := make(chan string)
	<-block
}
