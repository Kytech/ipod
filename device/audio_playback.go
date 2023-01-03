package device

import (
	"sync"
	"sync/atomic"

	uberatomic "github.com/oandrew/ipod/device/internal/atomic"
)

// Common playback states accross all lingos

type PlaybackState byte

const (
	PlaybackStopped        PlaybackState = 0x00
	PlaybackPlaying        PlaybackState = 0x01
	PlaybackPaused         PlaybackState = 0x02
	PlaybackFastForwarding PlaybackState = 0x03
	PlaybackRewinding      PlaybackState = 0x04
	PlaybackError          PlaybackState = 0xff
)

type ShuffleState byte

const (
	ShuffleOff    ShuffleState = 0x00
	ShuffleTracks ShuffleState = 0x01
	ShuffleAlbums ShuffleState = 0x02
)

type RepeatState byte

const (
	RepeatOff RepeatState = 0x00
	RepeatOne RepeatState = 0x01
	RepeatAll RepeatState = 0x02
)

type devPlaybackStatus struct {
	// Track info - TODO: Make these a single struct
	trackTitle  uberatomic.String
	artistName  uberatomic.String
	albumName   uberatomic.String
	trackLength atomic.Uint32

	// Playback status
	trackPosition atomic.Uint32
	playerState   atomic.Uint32

	shModeMux   sync.RWMutex
	shuffleMode ShuffleState

	rptModeMux sync.RWMutex
	repeatMode RepeatState

	// Track chapters
	chapterName uberatomic.String
}

// TODO: Refactor this to seperate getters instead of one
func (ps *devPlaybackStatus) PlaybackStatus() (state PlaybackState, trackLength, trackPos uint32) {
	return PlaybackState(ps.playerState.Load()), ps.trackLength.Load(), ps.trackPosition.Load()
}

func (ps *devPlaybackStatus) SetPlayerStatePlaying() {
	ps.playerState.Store(uint32(PlaybackPlaying))
}

func (ps *devPlaybackStatus) SetPlayerStatePaused() {
	ps.playerState.Store(uint32(PlaybackPaused))
}

func (ps *devPlaybackStatus) SetPlayerStateStopped() {
	ps.playerState.Store(uint32(PlaybackStopped))
}

func (ps *devPlaybackStatus) SetPlayerStateError() {
	ps.playerState.Store(uint32(PlaybackError))
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

func (ps *devPlaybackStatus) ShuffleMode() ShuffleState {
	ps.shModeMux.RLock()
	defer ps.shModeMux.RUnlock()
	return ps.shuffleMode
}

func (ps *devPlaybackStatus) SetShuffleMode(shMode ShuffleState) {
	ps.shModeMux.Lock()
	defer ps.shModeMux.Unlock()
	ps.shuffleMode = shMode
}

func (ps *devPlaybackStatus) RepeatMode() RepeatState {
	ps.rptModeMux.RLock()
	defer ps.rptModeMux.RUnlock()
	return ps.repeatMode
}

func (ps *devPlaybackStatus) SetRepeatMode(rptMode RepeatState) {
	ps.rptModeMux.Lock()
	defer ps.rptModeMux.Unlock()
	ps.repeatMode = rptMode
}
