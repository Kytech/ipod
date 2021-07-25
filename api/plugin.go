package api

import (
	"sync"
)

type EntryFunction func(Ipod, *sync.WaitGroup)

type Plugin struct {
	Name       string
	EntryPoint EntryFunction
}
