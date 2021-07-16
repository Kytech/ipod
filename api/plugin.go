package api

type Plugin interface {
	PluginMain(ipodState interface{})
}
