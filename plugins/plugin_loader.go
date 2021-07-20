package plugin

import (
	"fmt"
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
			Message: fmt.Sprintf("Plugin declaration struct IpodPlugin for plugin %s does not exist or is of incorrect type.", pluginFile),
			Plugin:  "plugin-loader",
		})
	}
	go plugin.EntryPoint(ipod)
}
