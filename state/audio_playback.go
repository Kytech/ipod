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

func (ps *playbackState) SetPlayerState(pState extremote.PlayerState) {
	ps.pbStateMux.Lock()
	defer ps.pbStateMux.Unlock()
	ps.playerState = pState
}

func (ps *playbackState) ChapterName() string {
	ps.chNameMux.RLock()
	defer ps.chNameMux.RUnlock()
	return ps.chapterName
}

func (ps *playbackState) SetChapterName(chName string) {
	ps.chNameMux.Lock()
	defer ps.chNameMux.Unlock()
	ps.chapterName = chName
}

func (ps *playbackState) TrackTitle() string {
	ps.trTitleMux.RLock()
	defer ps.trTitleMux.RUnlock()
	return ps.trackTitle
}

func (ps *playbackState) SetTrackTitle(trTitle string) {
	ps.trTitleMux.Lock()
	defer ps.trTitleMux.Unlock()
	ps.trackTitle = trTitle
}

func (ps *playbackState) TrackArtist() string {
	ps.trArtistMux.RLock()
	defer ps.trArtistMux.RUnlock()
	return ps.artistName
}

func (ps *playbackState) SetTrackArtist(trArtist string) {
	ps.trArtistMux.Lock()
	defer ps.trArtistMux.Unlock()
	ps.artistName = trArtist
}

func (ps *playbackState) TrackAlbum() string {
	ps.trAlbumMux.RLock()
	defer ps.trAlbumMux.RUnlock()
	return ps.albumName
}

func (ps *playbackState) SetTrackAlbum(trAlbum string) {
	ps.trAlbumMux.Lock()
	defer ps.trAlbumMux.Unlock()
	ps.albumName = trAlbum
}

func (ps *playbackState) ShuffleMode() extremote.ShuffleMode {
	ps.shModeMux.RLock()
	defer ps.shModeMux.RUnlock()
	return ps.shuffleMode
}

func (ps *playbackState) SetShuffleMode(shMode extremote.ShuffleMode) {
	ps.shModeMux.Lock()
	defer ps.shModeMux.Unlock()
	ps.shuffleMode = shMode
}

func (ps *playbackState) RepeatMode() extremote.RepeatMode {
	ps.rptModeMux.RLock()
	defer ps.rptModeMux.RUnlock()
	return ps.repeatMode
}

func (ps *playbackState) SetRepeatMode(rptMode extremote.RepeatMode) {
	ps.rptModeMux.Lock()
	defer ps.rptModeMux.Unlock()
	ps.repeatMode = rptMode
}
