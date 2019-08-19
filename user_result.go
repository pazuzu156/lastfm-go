package lastfm

import "encoding/xml"

// UserGetArtistTracks <- user.getArtistTracks
type UserGetArtistTracks struct {
	User       string `xml:"user,attr"`
	Artist     string `xml:"artist,attr"`
	Items      string `xml:"items,attr"`
	Total      int    `xml:"total,attr"`
	Page       int    `xml:"page,attr"`
	PerPage    int    `xml:"perPage,attr"`
	TotalPages int    `xml:"totalPages,attr"`
	Tracks     []struct {
		Artist struct {
			Mbid string `xml:"mbid,attr"`
			Name string `xml:",chardata"`
		} `xml:"artist"`
		Name       string `xml:"name"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:",chardata"`
		} `xml:"streamable"`
		Mbid  string `xml:"mbid"`
		Album struct {
			Mbid string `xml:"mbid,attr"`
			Name string `xml:",chardata"`
		} `xml:"album"`
		URL    string `xml:"url"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		Date struct {
			Uts string `xml:"uts,attr"`
			Str string `xml:",chardata"`
		} `xml:"date"`
	} `xml:"track"`
}

// UserGetFriends <- user.getFriends
type UserGetFriends struct {
	XMLName    xml.Name `xml:"friends"`
	For        string   `xml:"for,attr"` //username
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Friends    []struct {
		ID         string `xml:"id"`
		Name       string `xml:"name"`
		RealName   string `xml:"realname"`
		URL        string `xml:"url"`
		Country    string `xml:"country"`
		Age        string `xml:"age"`
		Gender     string `xml:"gender"`
		Subscriber string `xml:"subscriber"`
		Type       string `xml:"type"`
		PlayCount  string `xml:"playcount"`
		Playlists  string `xml:"playlists"`
		Bootstrap  string `xml:"bootstrap"` // ?
		Registered struct {
			Unixtime string `xml:"unixtime,attr"`
			Time     string `xml:",chardata"`
		} `xml:"registered"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		ScrobbleSource struct {
			Name  string `xml:"name"`
			Image string `xml:"image"`
			URL   string `xml:"url"`
		} `xml:"scrobblesource"`
		RecentTrack struct {
			Date   string `xml:"date,attr"`
			Uts    string `xml:"uts,attr"`
			Artist struct {
				Name string `xml:"name"`
				Mbid string `xml:"mbid"`
				URL  string `xml:"url"`
			} `xml:"artist"`
			Album struct {
				Name string `xml:"name"`
				Mbid string `xml:"mbid"`
				URL  string `xml:"url"`
			} `xml:"album"`
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			URL  string `xml:"url"`
		} `xml:"recenttrack"`
	} `xml:"user"`
}

// UserGetInfo <- user.getInfo
type UserGetInfo struct {
	XMLName    xml.Name `xml:"user"`
	ID         string   `xml:"id"`
	Name       string   `xml:"name"`
	RealName   string   `xml:"realname"`
	URL        string   `xml:"url"`
	Country    string   `xml:"country"`
	Age        string   `xml:"age"`
	Gender     string   `xml:"gender"`
	Subscriber string   `xml:"subscriber"`
	PlayCount  string   `xml:"playcount"`
	Playlists  string   `xml:"playlists"`
	Bootstrap  string   `xml:"bootstrap"`
	Registered struct {
		Unixtime string `xml:"unixtime,attr"`
		Time     string `xml:",chardata"`
	} `xml:"registered"`
	Type   string `xml:"type"` //user type: stuff, moderator, user...
	Images []struct {
		Size string `xml:"size,attr"`
		URL  string `xml:",chardata"`
	} `xml:"image"`
}

// UserGetLovedTracks <- user.getLovedTracks
type UserGetLovedTracks struct {
	XMLName    xml.Name `xml:"lovedtracks"`
	User       string   `xml:"user,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Name string `xml:"name"`
		Mbid string `xml:"mbid"`
		URL  string `xml:"url"`
		Date struct {
			Uts  string `xml:"uts,attr"`
			Date string `xml:",chardata"`
		} `xml:"date"`
		Artist struct {
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			URL  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:",chardata"`
		} `xml:"streamable"`
	} `xml:"track"`
}

// UserGetPersonalTags <- user.getPersonalTags
type UserGetPersonalTags struct {
	XMLName    xml.Name `xml:"taggings"`
	User       string   `xml:"user,attr"`
	Tag        string   `xml:"tag,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Artists    []struct {
		Name       string `xml:"name"`
		Mbid       string `xml:"mbid"`
		URL        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artists>artist"`
	Tracks []struct {
		Name       string `xml:"name"`
		Duration   string `xml:"duration"`
		Mbid       string `xml:"mbid"`
		URL        string `xml:"url"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:"streamable"`
		} `xml:"streamable"`
		Artist struct {
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			URL  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"tracks>track"`
}

// UserGetRecentTracks <- user.getRecentTracks
type UserGetRecentTracks struct {
	XMLName    xml.Name `xml:"recenttracks"`
	User       string   `xml:"user,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		NowPlaying string `xml:"nowplaying,attr,omitempty"`
		Artist     struct {
			Name string `xml:",chardata"`
			Mbid string `xml:"mbid,attr"`
		} `xml:"artist"`
		Name       string `xml:"name"`
		Streamable string `xml:"streamable"`
		Mbid       string `xml:"mbid"`
		Album      struct {
			Name string `xml:",chardata"`
			Mbid string `xml:"mbid,attr"`
		} `xml:"album"`
		URL    string `xml:"url"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		Date struct {
			Uts  string `xml:"uts,attr"`
			Date string `xml:",chardata"`
		} `xml:"date"`
	} `xml:"track"`
}

// UserGetTopAlbums <- user.getTopAlbums
type UserGetTopAlbums struct {
	XMLName    xml.Name `xml:"topalbums"`
	User       string   `xml:"user,attr"`
	Type       string   `xml:"type,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Albums     []struct {
		Rank      string `xml:"rank,attr"`
		Name      string `xml:"name"`
		PlayCount string `xml:"playcount"`
		Mbid      string `xml:"mbid"`
		URL       string `xml:"url"`
		Artist    struct {
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			URL  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"album"`
}

// UserGetTopArtists <- user.getTopArtists
type UserGetTopArtists struct {
	XMLName    xml.Name `xml:"topartists"`
	User       string   `xml:"user,attr"`
	Type       string   `xml:"type,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Artists    []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		PlayCount  string `xml:"playcount"`
		Mbid       string `xml:"mbid"`
		URL        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

// UserGetTopTags <- user.getTopTags
type UserGetTopTags struct {
	XMLName xml.Name `xml:"toptags"`
	User    string   `xml:"user,attr"`
	Tags    []struct {
		Name  string `xml:"name"`
		Count string `xml:"count"`
		URL   string `xml:"url"`
	} `xml:"tag"`
}

// UserGetTopTracks <- user.getTopTracks
type UserGetTopTracks struct {
	XMLName    xml.Name `xml:"toptracks"`
	User       string   `xml:"user,attr"`
	Type       string   `xml:"type,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		Duration   string `xml:"duration"`
		PlayCount  string `xml:"playcount"`
		Mbid       string `xml:"mbid"`
		URL        string `xml:"url"`
		Streamable struct {
			FullTrack  string `xml:"fulltrack,attr"`
			Streamable string `xml:",chardata"`
		} `xml:"streamable"`
		Artist struct {
			Name string `xml:"name"`
			Mbid string `xml:"mbid"`
			URL  string `xml:"url"`
		} `xml:"artist"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"track"`
}

// UserGetWeeklyAlbumChart <- user.getWeeklyAlbumChart
type UserGetWeeklyAlbumChart struct {
	XMLName xml.Name `xml:"weeklyalbumchart"`
	User    string   `xml:"user,attr"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
	Albums  []struct {
		Rank      string `xml:"rank,attr"`
		Name      string `xml:"name"`
		Mbid      string `xml:"mbid"`
		PlayCount string `xml:"playcount"`
		URL       string `xml:"url"`
	} `xml:"album"`
}

// UserGetWeeklyArtistChart <- user.getWeeklyArtistChart
type UserGetWeeklyArtistChart struct {
	XMLName xml.Name `xml:"weeklyartistchart"`
	User    string   `xml:"user,attr"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
	Artists []struct {
		Rank      string `xml:"rank,attr"`
		Name      string `xml:"name"`
		Mbid      string `xml:"mbid"`
		PlayCount string `xml:"playcount"`
		URL       string `xml:"url"`
	} `xml:"artist"`
}

// UserGetWeeklyChartList <- user.getWeeklyChartList
type UserGetWeeklyChartList struct {
	XMLName xml.Name `xml:"weeklychartlist"`
	User    string   `xml:"user,attr"`
	Charts  []struct {
		From string `xml:"from,attr"`
		To   string `xml:"to,attr"`
	} `xml:"chart"`
}

// UserGetWeeklyTrackChart <- user.getWeeklyTrackChart
type UserGetWeeklyTrackChart struct {
	XMLName xml.Name `xml:"weeklytrackchart"`
	User    string   `xml:"user,attr"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
	Tracks  []struct {
		Rank   string `xml:"rank,attr"`
		Artist struct {
			Name string `xml:",chardata"`
			Mbid string `xml:"mbid,attr"`
		} `xml:"artist"`
		Name      string `xml:"name"`
		Mbid      string `xml:"mbid"`
		PlayCount string `xml:"playcount"`
		Images    []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		URL string `xml:"url"`
	} `xml:"track"`
}
