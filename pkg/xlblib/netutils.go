// SPDX-License-Identifier: Apache 2.0
// Copyright (c) 2022 NetLOX Inc

package xlblib

import (
	"bufio"
	"encoding/binary"
	"net"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

// constants related to ifstats
const (
	RxBytes = iota
	RxPkts
	RxErrors
	RxDrops
	RxFifo
	RxFrame
	RxComp
	RxMcast
	TxBytes
	TxPkts
	TxErrors
	TxDrops
	TxFifo
	TxColls
	TxCarr
	TxComp
	MaxSidx
)

const (
	OsIfStatFile = "/proc/net/dev"
)

// IfiStat - Container of interface statistics
type IfiStat struct {
	Ifs [MaxSidx]uint64
}

// IsNetIPv4 - Check if net.IP is ipv4 address
func IsNetIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

// IsNetIPv6 - Check if net.IP is ipv6 address
func IsNetIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

// Ntohl - Network to host byte-order long
func Ntohl(i uint32) uint32 {
	return binary.BigEndian.Uint32((*(*[4]byte)(unsafe.Pointer(&i)))[:])
}

// Htonl - Host to network byte-order long
func Htonl(i uint32) uint32 {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return *(*uint32)(unsafe.Pointer(&b[0]))
}

// Htons - Host to network byte-order short
func Htons(i uint16) uint16 {
	var j = make([]byte, 2)
	binary.BigEndian.PutUint16(j[0:2], i)
	return *(*uint16)(unsafe.Pointer(&j[0]))
}

// Ntohs - Network to host byte-order short
func Ntohs(i uint16) uint16 {
	return binary.BigEndian.Uint16((*(*[2]byte)(unsafe.Pointer(&i)))[:])
}

// IPtonl - Convert net.IP to network byte-order long
func IPtonl(ip net.IP) uint32 {
	var val uint32

	if len(ip) == 16 {
		val = uint32(ip[12])
		val |= uint32(ip[13]) << 8
		val |= uint32(ip[14]) << 16
		val |= uint32(ip[15]) << 24
	} else {
		val = uint32(ip[0])
		val |= uint32(ip[1]) << 8
		val |= uint32(ip[2]) << 16
		val |= uint32(ip[3]) << 24
	}

	return val
}

// NltoIP - Convert network byte-order long to net.IP
func NltoIP(addr uint32) net.IP {
	var dip net.IP

	dip = append(dip, uint8(addr&0xff))
	dip = append(dip, uint8(addr>>8&0xff))
	dip = append(dip, uint8(addr>>16&0xff))
	dip = append(dip, uint8(addr>>24&0xff))

	return dip
}

// NetGetIfiStats - Get OS statistics for a given interface
func NetGetIfiStats(ifName string, ifs *IfiStat) int {
	file, err := os.Open(OsIfStatFile)
	if err != nil {
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := scanner.Text()
		ifi := strings.Split(content, ":")
		if len(ifi) > 1 {
			if ifi[0] != ifName {
				continue
			}

			ifSfs := strings.Fields(ifi[1])
			if len(ifSfs) >= MaxSidx {
				for i := 0; i < MaxSidx; i++ {
					val, err := strconv.ParseUint(ifSfs[i], 10, 64)
					if err == nil {
						ifs.Ifs[i] = val
					}
				}
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return -1
	}

	return 0
}
