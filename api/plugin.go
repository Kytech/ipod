package api

type EntryFunction func(Ipod, chan bool)

type Plugin struct {
	Name       string
	EntryPoint EntryFunction
}
