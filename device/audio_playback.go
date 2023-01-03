package device

import (
	"sync"
	"sync/atomic"

	uberatomic "github.com/oandrew/ipod/device/internal/atomic"
	extremote "github.com/oandrew/ipod/lingo-extremote"
)

type devPlaybackStatus struct {
	// Track info
	trackTitle uberatomic.String
	artistName uberatomic.String
	albumName  uberatomic.String

	// Playback status
	trackLength   atomic.Uint32
	trackPosition atomic.Uint32
	playerState   atomic.Uint32

	shModeMux   sync.RWMutex
	shuffleMode extremote.ShuffleMode

	rptModeMux sync.RWMutex
	repeatMode extremote.RepeatMode

	// Track chapters
	chapterName uberatomic.String
}

func (ps *devPlaybackStatus) PlaybackStatus() (trackLength, trackPos uint32, state extremote.PlayerState) {
	return ps.trackLength.Load(), ps.trackPosition.Load(), extremote.PlayerState(ps.playerState.Load())
}

func (ps *devPlaybackStatus) SetPlayerStatePlaying() {
	ps.playerState.Store(uint32(extremote.PlayerStatePlaying))
}

func (ps *devPlaybackStatus) SetPlayerStatePaused() {
	ps.playerState.Store(uint32(extremote.PlayerStatePaused))
}

func (ps *devPlaybackStatus) SetPlayerStateStopped() {
	ps.playerState.Store(uint32(extremote.PlayerStateStopped))
}

func (ps *devPlaybackStatus) SetPlayerStateError() {
	ps.playerState.Store(uint32(extremote.PlayerStateError))
}

func (ps *devPlaybackStatus) ChapterName() string {
	return ps.chapterName.Load()
}

func (ps *devPlaybackStatus) SetChapterName(chName string) {
	ps.chapterName.Store(chName)
}

func (ps *devPlaybackStatus) TrackTitle() string {
	return ps.trackTitle.Load()
}

func (ps *devPlaybackStatus) SetTrackTitle(trTitle string) {
	ps.trackTitle.Store(trTitle)
}

func (ps *devPlaybackStatus) TrackArtist() string {
	return ps.artistName.Load()
}

func (ps *devPlaybackStatus) SetTrackArtist(trArtist string) {
	ps.artistName.Store(trArtist)
}

func (ps *devPlaybackStatus) TrackAlbum() string {
	return ps.albumName.Load()
}

func (ps *devPlaybackStatus) SetTrackAlbum(trAlbum string) {
	ps.albumName.Store(trAlbum)
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
