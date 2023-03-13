package common

import (
	"github.com/tortuecucu/pathfinder/pkg/core"
	"golang.org/x/sys/windows/registry"
)

func getRegistryStringValue(key registry.Key, path string, name string) (string, error) {
	k, err := registry.OpenKey(key, path, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue(name)
	if err != nil {
		return "", err
	}
	return s, nil
}

type ProxyPacUrl struct {
}

func (p ProxyPacUrl) Run(exe *core.FactCollection) {
	const factKey string = "proxy.proxypac.autoconfigurl"
	v, err := getRegistryStringValue(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, "AutoConfigURL")
	if err != nil {
		exe.Facts[factKey] = core.NewFact(p, err, factKey)
	} else {
		exe.Facts[factKey] = core.NewFact(p, v, factKey)
	}

}
func (p ProxyPacUrl) Name() string {
	return "ProxyPacUrl"
}

func AddCoreCollectors(plan core.Plan) {
	plan.AddCollector(ProxyPacUrl{})

	//TODO: try to ping server
	//TODO: try to get the file

}
