package plugins

import (
	"github.com/oandrew/ipod/api"
)

var pluginRegistry = make(map[string]*api.Plugin)
var defaultPluginRef *api.Plugin

func registerPlugin(plugin *api.Plugin) {
	if plugin.Name == "default" {
		defaultPluginRef = plugin
	} else {
		pluginRegistry[plugin.Name] = plugin
	}
}

func getRegisteredPlugins() (defaultPlugin *api.Plugin, plugins []*api.Plugin) {
	defaultPlugin = defaultPluginRef
	plugins = make([]*api.Plugin, 0, len(pluginRegistry))
	for _, val := range pluginRegistry {
		plugins = append(plugins, val)
	}
	return
}

func GetRegisteredPluginNames() []string {
	plugins := make([]string, 0, len(pluginRegistry))
	for k := range pluginRegistry {
		plugins = append(plugins, k)
	}
	return plugins
}
