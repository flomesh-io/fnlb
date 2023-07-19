package ippool

import (
	"fmt"
	"net"
	"sync/atomic"
)

const (
	mask24 = uint32(0xFF << 24)
	mask16 = uint32(0xFF << 16)
	mask8  = uint32(0xFF << 8)
	mask0  = uint32(0xFF)
)

type IPGenerater struct {
	netCIDR      *net.IPNet
	uintMaskBits uint32
	counter      uint32
}

func InitIPGenerater(cidr string) (*IPGenerater, error) {
	_, ipn, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR format")
	}

	maskbit, _ := ipn.Mask.Size()
	uintMaskBits := uint32(0xFFFFFFFF >> maskbit)

	return &IPGenerater{
		netCIDR:      ipn,
		uintMaskBits: uintMaskBits,
	}, nil
}

func (i *IPGenerater) NextIP() net.IP {
	counter := i.counter
	for !atomic.CompareAndSwapUint32(&i.counter, counter, counter+1) {
		counter = i.counter
	}

	counterBit := counter & i.uintMaskBits

	return net.IP{
		byte(uint32(i.netCIDR.IP[0]) | (counterBit&mask24)>>24),
		byte(uint32(i.netCIDR.IP[1]) | (counterBit&mask16)>>16),
		byte(uint32(i.netCIDR.IP[2]) | (counterBit&mask8)>>8),
		byte(uint32(i.netCIDR.IP[3]) | (counterBit&mask0)>>0),
	}
}

func (i *IPGenerater) GetBroadcastIP() net.IP {
	return net.IP{
		byte(uint32(i.netCIDR.IP[0]) | (i.uintMaskBits)>>24),
		byte(uint32(i.netCIDR.IP[1]) | (i.uintMaskBits)>>16),
		byte(uint32(i.netCIDR.IP[2]) | (i.uintMaskBits)>>8),
		byte(uint32(i.netCIDR.IP[3]) | (i.uintMaskBits)>>0),
	}
}

func (i *IPGenerater) GetNetwork() net.IP {
	return i.netCIDR.IP
}

func (i *IPGenerater) CheckIPAddressInSubnet(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return i.netCIDR.Contains(ip)
}
