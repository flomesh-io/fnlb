package xlblib

import (
	"errors"
	"net"
	"runtime"
)

var ErrUnsupported = errors.New("SCTP is unsupported on " + runtime.GOOS + "/" + runtime.GOARCH)

// ArpPing - sends a arp request given the DIP, SIP and interface name
func ArpPing(DIP net.IP, SIP net.IP, ifName string) (int, error) {
	return 0, ErrUnsupported
}
