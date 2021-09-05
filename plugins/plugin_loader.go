package plugins

import (
	"fmt"
	"plugin"
	"sync"

	"io/ioutil"
	"path/filepath"

	"github.com/oandrew/ipod/api"
	"github.com/oandrew/ipod/state"
)

func LoadPluginsDir(pluginDirPath string) []error {
	files, err := ioutil.ReadDir(pluginDirPath)
	if err != nil {
		return err
	}

	var pluginLoadErrors []error
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".so" {
			err := openPlugin(filepath.Join(pluginDirPath, file.Name()))
			if err != nil {
				pluginLoadErrors = append(pluginLoadErrors, err)
			}
		}
	}

	return pluginLoadErrors
}

func openPlugin(pluginFilePath string) error {
	p, err := plugin.Open(pluginFilePath)
	if err != nil {
		return fmt.Errorf("error in plugin %s: %w", pluginFilePath, err)
	}
	pluginInfo, err := p.Lookup("IpodPlugin")
	if err != nil {
		return fmt.Errorf("error in plugin %s: %w", pluginFilePath, err)
	}
	plugin, ok := pluginInfo.(*api.Plugin)
	if !ok {
		return fmt.Errorf("error in plugin %s: Plugin declaration does not match required fields or is of incorrect struct type", pluginFilePath)
	}
	registerPlugin(plugin)
	return nil
}

func InitPlugins(ipod *state.IpodState) {
	defaultPlugin, plugins := getRegisteredPlugins()
	if defaultPlugin != nil {
		var defaultWaitGroup sync.WaitGroup
		defaultWaitGroup.Add(1)
		go defaultPlugin.EntryPoint(ipod, &defaultWaitGroup)
		defaultWaitGroup.Wait()
	}

	var waitGroup sync.WaitGroup
	for _, plugin := range plugins {
		waitGroup.Add(1)
		go plugin.EntryPoint(ipod, &waitGroup)
	}

	waitGroup.Wait()
}
