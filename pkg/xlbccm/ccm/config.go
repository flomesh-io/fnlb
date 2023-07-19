package ccm

import (
	"fmt"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
)

type CCMConfig struct {
	APIServerUrlStrList []string `yaml:"apiServerURL"`
	ExternalCIDR        string   `yaml:"externalCIDR"`
	SetBGP              bool     `yaml:"setBGP"`
	SetLBMode           int32    `yaml:"setLBMode"`
	ExclIPAM            bool     `yaml:"setExclIPAM"`
	APIServerUrlList    []*url.URL
}

func ReadCCMConfig(configBytes []byte) (CCMConfig, error) {
	o := CCMConfig{}

	if err := yaml.Unmarshal(configBytes, &o); err != nil {
		return o, fmt.Errorf("failed to unmarshal config. err: %v", err)
	}

	for _, u := range o.APIServerUrlStrList {
		apiURL, err := url.Parse(u)
		if err != nil {
			return o, err
		}

		klog.Infof("add fsmxlb API server %s", u)
		o.APIServerUrlList = append(o.APIServerUrlList, apiURL)
	}
	return o, nil
}

func ReadCCMEnvronment() (CCMConfig, error) {
	ccmConfigStr, ok := os.LookupEnv("FSM_CCM_CONFIG")
	if !ok {
		return CCMConfig{}, fmt.Errorf("not found FSM_CCM_CONFIG env")
	}

	return ReadCCMConfig([]byte(ccmConfigStr))
}
