package xlbcli

import (
	"fmt"
	"sort"
)

type FDB struct {
	CommonAPI
}

type FDBModGet struct {
	FdbAttr []FDBMod `json:"fdbAttr"`
}

type FDBMod struct {
	// Dev - name of the related device
	Dev string `json:"dev" yaml:"dev"`
	// MacAddress - Hardware address
	MacAddress string `json:"macAddress" yaml:"macAddress"`
}

type ConfigurationFDBFile struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec       FDBMod `yaml:"spec"`
}

func (fdb FDBMod) Key() string {
	return fmt.Sprintf("%s|%s", fdb.Dev, fdb.MacAddress)
}

func (FDBresp FDBModGet) Sort() {
	sort.Slice(FDBresp.FdbAttr, func(i, j int) bool {
		return FDBresp.FdbAttr[i].Key() < FDBresp.FdbAttr[j].Key()
	})
}
