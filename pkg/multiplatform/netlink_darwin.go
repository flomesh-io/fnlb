package multiplatform

import (
	"github.com/vishvananda/netlink"
)

// RouteListFiltered gets a list of routes in the system filtered with specified rules.
// All rules must be defined in RouteFilter struct
func RouteListFiltered(family int, filter *netlink.Route, filterMask uint64) ([]netlink.Route, error) {
	panic("Unsupported!")
}
