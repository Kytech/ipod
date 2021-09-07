package ipod_defaults

import (
	"sync"

	"github.com/oandrew/ipod/api"
)

var IpodPlugin = api.Plugin{
	Name:       "default",
	EntryPoint: PluginMain,
}

func PluginMain(ipod api.Ipod, wg *sync.WaitGroup) {
	ipod.SetPlayerStatePlaying()
	ipod.SetTrackArtist("Test Artist")
	wg.Done()
}
