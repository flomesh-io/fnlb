package xlbnlp

import (
	cmn "github.com/cybwan/fsmxlb/pkg/common"
	nlp "github.com/vishvananda/netlink"
)

var hooks cmn.NetHookInterface

func NlpRegister(hook cmn.NetHookInterface) {
	hooks = hook
}

func GetLinkNameByIndex(index int) (string, error) {
	brLink, err := nlp.LinkByIndex(index)
	if err != nil {
		return "", err
	}
	return brLink.Attrs().Name, nil
}
