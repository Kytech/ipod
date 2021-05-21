package state

type IpodState struct {
	devGeneral
	playbackState
}

func (is *IpodState) Name() string {
	return "ipod-gadget"
}
