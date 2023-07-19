package xlbcli

// TypeMeta describes an individual object in an API response or request
// with strings representing the type of the object and its API schema version.
// Structures that are versioned or persisted should inline TypeMeta.
type TypeMeta struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

type ObjectMeta struct {
	BGP      bool   `yaml:"bgp,omitempty"`
	Block    int    `yaml:"block,omitempty"`
	HostName string `yaml:"hostname,omitempty"`
	VlanID   int    `yaml:"vid,omitempty"`
	VxlanID  int    `yaml:"vxlanID,omitempty"`
}
