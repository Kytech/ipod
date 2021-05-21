package state

type playbackState struct {
	// Track info
	trackTitle string
	artistName string
	albumName  string

	// Extended track info
	genreName         string
	composerName      string
	albumTrackIndex   uint16
	discSetAlbumIndex uint16

	// Playback status
	trackLength   uint32
	trackPosition uint32
	playerState   byte
	shuffleMode   byte
	repeatMode    byte
	numTracks     uint32

	// Track chapters
	chapterIndex    int32
	chapterCount    int32
	chapterLength   uint32
	chapterPosition uint32
	chapterName     string
}
