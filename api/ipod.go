package api

type Ipod interface {
	Name() string

	SetPlayerStatePlaying()
	SetPlayerStatePaused()
	SetPlayerStateStopped()
	SetPlayerStateError()

	ChapterName() string
	SetChapterName(string)
	TrackTitle() string
	SetTrackTitle(string)
	TrackArtist() string
	SetTrackArtist(string)
	TrackAlbum() string
	SetTrackAlbum(string)
}
