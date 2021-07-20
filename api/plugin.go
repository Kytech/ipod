package api

type EntryFunction func(Ipod)

type Plugin struct {
	EntryPoint EntryFunction
}
