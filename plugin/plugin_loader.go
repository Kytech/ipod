package plugin

import (
	"plugin"

	"github.com/oandrew/ipod/state"
)

func runPlugin(pluginFile string, ipod *state.IpodState) {
	p, err := plugin.Open(pluginFile)
	if err != nil {
		panic(err)
	}
	pluginMain, err := p.Lookup("PluginMain")
	if err != nil {
		panic(err)
	}
	go pluginMain.(func(interface{}))(ipod)
}
