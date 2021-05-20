package state

type IpodState struct {
	devGeneral
}

func (is *IpodState) Name() string {
	return "ipod-gadget"
}
