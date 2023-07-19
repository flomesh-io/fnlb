package ippool

import (
	"errors"
	"k8s.io/klog/v2"
	"net"
	"sync"

	tk "github.com/cybwan/fsmxlb/pkg/xlblib"
)

type IPPool struct {
	CIDR    string
	NetCIDR *net.IPNet
	IPAlloc *tk.IPAllocator
	mutex   sync.Mutex
	Shared  bool
}

// Initailize IP Pool
func NewIPPool(ipa *tk.IPAllocator, CIDR string, Shared bool) (*IPPool, error) {
	ipa.AddIPRange(tk.IPClusterDefault, CIDR)

	_, ipn, err := net.ParseCIDR(CIDR)
	if err != nil {
		return nil, errors.New("CIDR parse failed")
	}

	return &IPPool{
		CIDR:    CIDR,
		NetCIDR: ipn,
		IPAlloc: ipa,
		mutex:   sync.Mutex{},
		Shared:  Shared,
	}, nil
}

// GetNewIPAddr generate new IP and add key(IP) in IP Pool.
// If IP is already in pool, try to generate next IP.
// Returns nil If all IPs in the subnet are already in the pool.
func (i *IPPool) GetNewIPAddr(sIdent uint32, proto string) net.IP {

	i.mutex.Lock()
	defer i.mutex.Unlock()

	if !i.Shared {
		sIdent = 0
	}

	newIP, err := i.IPAlloc.AllocateNewIP(tk.IPClusterDefault, i.CIDR, sIdent, proto)
	if err != nil {
		return nil
	}

	klog.Infof("Allocate ServiceIP %s:%v", newIP.String(), sIdent)

	return newIP
}

// ReturnIPAddr return IPaddress in IP Pool
func (i *IPPool) ReturnIPAddr(ip string, sIdent uint32, proto string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if !i.Shared {
		sIdent = 0
	}

	klog.Infof("Release ServiceIP %s:%v", ip, sIdent)

	i.IPAlloc.DeAllocateIP(tk.IPClusterDefault, i.CIDR, sIdent, ip, proto)
}

// ReserveIPAddr reserve this IPaddress in IP Pool
func (i *IPPool) ReserveIPAddr(ip string, sIdent uint32, proto string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if !i.Shared {
		sIdent = 0
	}

	klog.Infof("Reserve ServiceIP %s:%v", ip, sIdent)

	i.IPAlloc.ReserveIP(tk.IPClusterDefault, i.CIDR, sIdent, ip, proto)
}

// CheckAndReserveIP check and reserve this IPaddress in IP Pool
func (i *IPPool) CheckAndReserveIP(ip string, sIdent uint32, proto string) bool {
	IP := net.ParseIP(ip)
	if IP != nil && i.NetCIDR.Contains(IP) {
		i.ReserveIPAddr(ip, sIdent, proto)
		return true
	}

	return false
}
