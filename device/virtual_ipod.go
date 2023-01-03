// Package device provides a virtual iPod device that can be controlled by
// iPod Accessory Protocol commands.
package device

import "github.com/sirupsen/logrus"

// VirtualIpod is a virtual iPod device used to simulate a real iPod that is
// connected over iPod Accessory Protocol (iAP). It provides an interface for
// connecting to events sent over iAP. It also enables sending data and
// messages to the other device on the iAP connection. VirtualIpod is safe for
// concurrent use.
type VirtualIpod struct {
	devGeneral
	devPlaybackStatus
}

// NewVirtualIpod creates a new VirtualIpod with the default logger.
func NewVirtualIpod() *VirtualIpod {
	return NewVirtualIpodWithLogger(logrus.StandardLogger())
}

// NewVirtualIpodWithLogger creates a new VirtualIpod with the specified logger.
func NewVirtualIpodWithLogger(logger *logrus.Logger) *VirtualIpod {
	return &VirtualIpod{
		devGeneral: devGeneral{
			Logger: logger,
		},
	}
}

func (is *VirtualIpod) Name() string {
	return "ipod-gadget"
}
