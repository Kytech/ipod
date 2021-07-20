package plugin

import (
	"fmt"
)

type PluginError struct {
	Message string
	Plugin  string
}

func (err *PluginError) Error() string {
	return fmt.Sprintf("Error in plugin %s: %s", err.Plugin, err.Message)
}
