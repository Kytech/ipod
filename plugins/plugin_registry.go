package plugins

import (
	"github.com/oandrew/ipod/api"
)

var pluginRegistry = make(map[string]*api.Plugin)
