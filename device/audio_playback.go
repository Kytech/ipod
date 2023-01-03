package device

import (
	"sync"

	uberatomic "github.com/oandrew/ipod/device/internal/atomic"
	extremote "github.com/oandrew/ipod/lingo-extremote"
)

type devPlaybackStatus struct {

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
	chapterName uberatomic.String
}

func (ps *devPlaybackStatus) PlaybackStatus() (trackLength, trackPos uint32, state extremote.PlayerState) {
	ps.pbStateMux.RLock()
	defer ps.pbStateMux.RUnlock()
	return ps.trackLength, ps.trackPosition, ps.playerState
}

func (ps *devPlaybackStatus) SetPlayerStatePlaying() {
	ps.pbStateMux.Lock()
	defer ps.pbStateMux.Unlock()
	ps.playerState = extremote.PlayerStatePlaying
}

func (ps *devPlaybackStatus) SetPlayerStatePaused() {
	ps.pbStateMux.Lock()
	defer ps.pbStateMux.Unlock()
	ps.playerState = extremote.PlayerStatePaused
}

func (ps *devPlaybackStatus) SetPlayerStateStopped() {
	ps.pbStateMux.Lock()
	defer ps.pbStateMux.Unlock()
	ps.playerState = extremote.PlayerStateStopped
}

func (ps *devPlaybackStatus) SetPlayerStateError() {
	ps.pbStateMux.Lock()
	defer ps.pbStateMux.Unlock()
	ps.playerState = extremote.PlayerStateError
}

func (ps *devPlaybackStatus) ChapterName() string {
	return ps.chapterName.Load()
}

func (ps *devPlaybackStatus) SetChapterName(chName string) {
	ps.chapterName.Store(chName)
}

func (ps *devPlaybackStatus) TrackTitle() string {
	ps.trTitleMux.RLock()
	defer ps.trTitleMux.RUnlock()
	return ps.trackTitle
}

func (ps *devPlaybackStatus) SetTrackTitle(trTitle string) {
	ps.trTitleMux.Lock()
	defer ps.trTitleMux.Unlock()
	ps.trackTitle = trTitle
}

func (ps *devPlaybackStatus) TrackArtist() string {
	ps.trArtistMux.RLock()
	defer ps.trArtistMux.RUnlock()
	return ps.artistName
}

func (ps *devPlaybackStatus) SetTrackArtist(trArtist string) {
	ps.trArtistMux.Lock()
	defer ps.trArtistMux.Unlock()
	ps.artistName = trArtist
}

func (ps *devPlaybackStatus) TrackAlbum() string {
	ps.trAlbumMux.RLock()
	defer ps.trAlbumMux.RUnlock()
	return ps.albumName
}

func (ps *devPlaybackStatus) SetTrackAlbum(trAlbum string) {
	ps.trAlbumMux.Lock()
	defer ps.trAlbumMux.Unlock()
	ps.albumName = trAlbum
}

func (ps *devPlaybackStatus) ShuffleMode() extremote.ShuffleMode {
	ps.shModeMux.RLock()
	defer ps.shModeMux.RUnlock()
	return ps.shuffleMode
}

func (ps *devPlaybackStatus) SetShuffleMode(shMode extremote.ShuffleMode) {
	ps.shModeMux.Lock()
	defer ps.shModeMux.Unlock()
	ps.shuffleMode = shMode
}

func (ps *devPlaybackStatus) RepeatMode() extremote.RepeatMode {
	ps.rptModeMux.RLock()
	defer ps.rptModeMux.RUnlock()
	return ps.repeatMode
}

func (ps *devPlaybackStatus) SetRepeatMode(rptMode extremote.RepeatMode) {
	ps.rptModeMux.Lock()
	defer ps.rptModeMux.Unlock()
	ps.repeatMode = rptMode
}
