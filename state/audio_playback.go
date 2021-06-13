package state

import (
	"sync"

	extremote "github.com/oandrew/ipod/lingo-extremote"
)

type playbackState struct {

	// Track info

	trTitleMux sync.RWMutex
	trackTitle string

	trArtistMux sync.RWMutex
	artistName  string

	trAlbumMux sync.RWMutex
	albumName  string

	// Playback status

	pbStateMux    sync.RWMutex
	trackLength   uint32
	trackPosition uint32
	playerState   extremote.PlayerState

	shModeMux   sync.RWMutex
	shuffleMode extremote.ShuffleMode

	rptModeMux sync.RWMutex
	repeatMode extremote.RepeatMode

	// Track chapters

	chNameMux   sync.RWMutex
	chapterName string
}

func (ps *playbackState) PlaybackStatus() (trackLength, trackPos uint32, state extremote.PlayerState) {
	ps.pbStateMux.RLock()
	defer ps.pbStateMux.RUnlock()
	return ps.trackLength, ps.trackPosition, ps.playerState
}

func (ps *playbackState) ChapterName() string {
	ps.chNameMux.RLock()
	defer ps.chNameMux.RUnlock()
	return ps.chapterName
}

func (ps *playbackState) TrackTitle() string {
	ps.trTitleMux.RLock()
	defer ps.trTitleMux.RUnlock()
	return ps.trackTitle
}

func (ps *playbackState) TrackArtist() string {
	ps.trArtistMux.RLock()
	defer ps.trArtistMux.RUnlock()
	return ps.artistName
}

func (ps *playbackState) TrackAlbum() string {
	ps.trAlbumMux.RLock()
	defer ps.trAlbumMux.RUnlock()
	return ps.albumName
}

func (ps *playbackState) ShuffleMode() extremote.ShuffleMode {
	ps.shModeMux.RLock()
	defer ps.shModeMux.RUnlock()
	return ps.shuffleMode
}

func (ps *playbackState) RepeatMode() extremote.RepeatMode {
	ps.rptModeMux.RLock()
	defer ps.rptModeMux.RUnlock()
	return ps.repeatMode
}
