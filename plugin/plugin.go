package plugin

type Plugin interface {
	PluginMain(ipodState interface{})
}
