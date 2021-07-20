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
	plugin, ok := pluginInfo.(*api.Plugin)
	if !ok {
		panic(PluginError{
			Message: "Plugin declaration struct \"IpodPlugin\" does not exist or is of incorrect type.",
			Plugin:  pluginFile,
		})
	}
	initFinished := make(chan bool)
	go plugin.EntryPoint(ipod, initFinished)
	<-initFinished
}
