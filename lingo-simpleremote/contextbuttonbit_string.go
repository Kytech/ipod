// Code generated by "stringer -type=ContextButtonBit"; DO NOT EDIT.

package simpleremote

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ContextButtonPlayPause-1]
	_ = x[ContextButtonVolumeUp-2]
	_ = x[ContextButtonVolumeDown-4]
	_ = x[ContextButtonNextTrack-8]
	_ = x[ContextButtonPreviousTrack-16]
	_ = x[ContextButtonNextAlbum-32]
	_ = x[ContextButtonPreviousAlbum-64]
	_ = x[ContextButtonStop-128]
	_ = x[ContextButtonPlayResume-256]
	_ = x[ContextButtonPause-512]
	_ = x[ContextButtonMuteToggle-1024]
	_ = x[ContextButtonNextChapter-2048]
	_ = x[ContextButtonPreviousChapter-4096]
	_ = x[ContextButtonNextPlaylist-8192]
	_ = x[ContextButtonPreviousPlaylist-16384]
	_ = x[ContextButtonShuffleSettingAdvance-32768]
	_ = x[ContextButtonRepeatSettingAdvance-65536]
	_ = x[ContextButtonPowerOn-131072]
	_ = x[ContextButtonPowerOff-262144]
	_ = x[ContextButtonBacklightfor30Seconds-524288]
	_ = x[ContextButtonBeginFastForward-1048576]
	_ = x[ContextButtonBeginRewind-2097152]
	_ = x[ContextButtonMenu-4194304]
	_ = x[ContextButtonSelect-8388608]
	_ = x[ContextButtonUpArrow-16777216]
	_ = x[ContextButtonDownArrow-33554432]
	_ = x[ContextButtonBacklightOff-67108864]
}

const _ContextButtonBit_name = "ContextButtonPlayPauseContextButtonVolumeUpContextButtonVolumeDownContextButtonNextTrackContextButtonPreviousTrackContextButtonNextAlbumContextButtonPreviousAlbumContextButtonStopContextButtonPlayResumeContextButtonPauseContextButtonMuteToggleContextButtonNextChapterContextButtonPreviousChapterContextButtonNextPlaylistContextButtonPreviousPlaylistContextButtonShuffleSettingAdvanceContextButtonRepeatSettingAdvanceContextButtonPowerOnContextButtonPowerOffContextButtonBacklightfor30SecondsContextButtonBeginFastForwardContextButtonBeginRewindContextButtonMenuContextButtonSelectContextButtonUpArrowContextButtonDownArrowContextButtonBacklightOff"

var _ContextButtonBit_map = map[ContextButtonBit]string{
	1:        _ContextButtonBit_name[0:22],
	2:        _ContextButtonBit_name[22:43],
	4:        _ContextButtonBit_name[43:66],
	8:        _ContextButtonBit_name[66:88],
	16:       _ContextButtonBit_name[88:114],
	32:       _ContextButtonBit_name[114:136],
	64:       _ContextButtonBit_name[136:162],
	128:      _ContextButtonBit_name[162:179],
	256:      _ContextButtonBit_name[179:202],
	512:      _ContextButtonBit_name[202:220],
	1024:     _ContextButtonBit_name[220:243],
	2048:     _ContextButtonBit_name[243:267],
	4096:     _ContextButtonBit_name[267:295],
	8192:     _ContextButtonBit_name[295:320],
	16384:    _ContextButtonBit_name[320:349],
	32768:    _ContextButtonBit_name[349:383],
	65536:    _ContextButtonBit_name[383:416],
	131072:   _ContextButtonBit_name[416:436],
	262144:   _ContextButtonBit_name[436:457],
	524288:   _ContextButtonBit_name[457:491],
	1048576:  _ContextButtonBit_name[491:520],
	2097152:  _ContextButtonBit_name[520:544],
	4194304:  _ContextButtonBit_name[544:561],
	8388608:  _ContextButtonBit_name[561:580],
	16777216: _ContextButtonBit_name[580:600],
	33554432: _ContextButtonBit_name[600:622],
	67108864: _ContextButtonBit_name[622:647],
}

func (i ContextButtonBit) String() string {
	if str, ok := _ContextButtonBit_map[i]; ok {
		return str
	}
	return "ContextButtonBit(" + strconv.FormatInt(int64(i), 10) + ")"
}