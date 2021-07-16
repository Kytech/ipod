package api

type Ipod interface {
	SetPlayerStatePlaying()
	SetTrackTitle(string)
	SetTrackArtist(string)
	SetTrackAlbum(string)
}
