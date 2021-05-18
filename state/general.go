package state

import (
	"bytes"
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/oandrew/ipod"
	general "github.com/oandrew/ipod/lingo-general"

	"github.com/fullsailor/pkcs7"
)

type devGeneral struct {
	uimode general.UIMode
	tokens []general.FIDTokenValue
}

var _ general.DeviceGeneral = &devGeneral{}

func (d *devGeneral) UIMode() general.UIMode {
	return d.uimode
}

func (d *devGeneral) SetUIMode(mode general.UIMode) {
	d.uimode = mode
}

func (d *devGeneral) Name() string {
	return "ipod-gadget"
}

func (d *devGeneral) SoftwareVersion() (major uint8, minor uint8, rev uint8) {
	return 7, 1, 2
}

func (d *devGeneral) SerialNum() string {
	return "abcd1234"
}

func (d *devGeneral) LingoProtocolVersion(lingo uint8) (major uint8, minor uint8) {
	switch lingo {
	case ipod.LingoGeneralID:
		return 1, 9
	case ipod.LingoDisplayRemoteID:
		return 1, 5
	case ipod.LingoExtRemoteID:
		return 1, 12
	case ipod.LingoDigitalAudioID:
		return 1, 2
	default:
		return 1, 1
	}
}

func (d *devGeneral) LingoOptions(lingo uint8) uint64 {
	switch lingo {
	case ipod.LingoGeneralID:
		return 0x000000063DEF73FF

	default:
		return 0
	}
}

func (d *devGeneral) PrefSettingID(classID uint8) uint8 {
	return 0
}

func (d *devGeneral) SetPrefSettingID(classID uint8, settingID uint8, restoreOnExit bool) {
}

func (d *devGeneral) SetEventNotificationMask(mask uint64) {

}

func (d *devGeneral) EventNotificationMask() uint64 {
	return 0
}

func (d *devGeneral) SupportedEventNotificationMask() uint64 {
	return 0
}

func (d *devGeneral) CancelCommand(lingo uint8, cmd uint16, transaction uint16) {

}

func (d *devGeneral) MaxPayload() uint16 {
	return 65535
}

func (d *devGeneral) StartIDPS() {
	d.tokens = make([]general.FIDTokenValue, 0)
}

func (d *devGeneral) SetToken(token general.FIDTokenValue) error {
	d.tokens = append(d.tokens, token)
	return nil
}

func (d *devGeneral) EndIDPS(status general.AccEndIDPSStatus) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Tokens:\n")
	for _, token := range d.tokens {

		fmt.Fprintf(&buf, "* Token: %T\n", token.Token)
		//log.Printf("New token: %T", token.Token)
		switch t := token.Token.(type) {
		case *general.FIDIdentifyToken:

		case *general.FIDAccCapsToken:
			for _, c := range general.AccCaps {
				if t.AccCapsBitmask&uint64(c) != 0 {
					fmt.Fprintf(&buf, "Capability: %v\n", c)
				}
			}
		case *general.FIDAccInfoToken:
			key := general.AccInfoType(t.AccInfoType).String()
			fmt.Fprintf(&buf, "%s: %s\n", key, spew.Sdump(t.Value))

		case *general.FIDiPodPreferenceToken:

		case *general.FIDEAProtocolToken:

		case *general.FIDBundleSeedIDPrefToken:

		case *general.FIDScreenInfoToken:

		case *general.FIDEAProtocolMetadataToken:

		case *general.FIDMicrophoneCapsToken:

		}

	}
	log.Print(buf.String())
}

func (d *devGeneral) AccAuthCert(cert []byte) {
	pkcs, err := pkcs7.Parse(cert)
	if err != nil {
		log.Error(err)
		return
	}
	if len(pkcs.Certificates) >= 1 {
		cn := pkcs.Certificates[0].Subject.CommonName
		log.Infof("cert: CN=%s", cn)
	}

}
