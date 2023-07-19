package multiplatform

import "golang.org/x/sys/unix"

const (
	// Family type definitions
	FAMILY_ALL = unix.AF_UNSPEC
	FAMILY_V4  = unix.AF_INET

	RT_FILTER_PROTOCOL = 2
)
