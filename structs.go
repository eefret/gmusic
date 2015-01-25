package gmusic

//AddPlaylist is used to add a new playlist
type AddPlaylist struct {
	id      string
	title   string
	success bool
}

//DeletePlaylist is used to delete a playlist
type DeletePlaylist struct {
	deleteId string
}

//Tune is an abstract type
type Tune struct {
	genre          string
	beatsPerMinute int16
	album          string
	id             string
	composer       string
	title          string
	albumArtist    string
	year           int16
	artist         string
	durationMillis int64
	deleted        bool
	playCount      int16
	rating         string
	comment        string
}

//Song is a single song
type Song struct {
	Tune
	totalTracks       int16
	subjectToCuration bool
	name              string
	totalDiscs        int16
	titleNorm         string
	albumNorm         string
	track             int16
	albumArtUrl       string
	url               string
	creationDate      float32
	albumArtistNorm   string
	artistNorm        string
	lastPlayed        float64
	metajamId         string
	songType          int16 //type
	disc              int16
}

//Playlist is a set of songs
type Playlist struct {
	playlistId         string
	title              string
	requestTime        float32
	continuationToken  string
	differentialUpdate bool
	playlist           []Song
}

type Playlists struct {
	playlists      []Playlist
	magicPlaylists []Playlist
}

type SongUrl struct {
	url string
}

type QueryResults struct {
	artists []Song
	albums  []Song
	songs   []Song
}
