package plugin

import (
	"plugin"

	"github.com/oandrew/ipod/api"
	"github.com/oandrew/ipod/state"
)

func runPlugin(pluginFile string, ipod *state.IpodState) {
	p, err := plugin.Open(pluginFile)
	if err != nil {
		panic(err)
	}
	pluginInfo, err := p.Lookup("IpodPlugin")
	if err != nil {
		panic(err)
	}
	plugin := *pluginInfo.(*api.Plugin)
	go plugin.EntryPoint(ipod)
}
