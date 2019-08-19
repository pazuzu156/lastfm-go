package lastfm

import "encoding/xml"

//album.addTags (empty)

// AlbumGetInfo <- album.getInfo
type AlbumGetInfo struct {
	XMLName     xml.Name `xml:"album"`
	Name        string   `xml:"name"`
	Artist      string   `xml:"artist"`
	ID          string   `xml:"id"`
	Mbid        string   `xml:"mbid"`
	URL         string   `xml:"url"`
	ReleaseDate string   `xml:"releasedate"`
	Images      []struct {
		Size string `xml:"size,attr"`
		URL  string `xml:",chardata"`
	} `xml:"image"`
	Listeners     string `xml:"listeners"`
	PlayCount     string `xml:"playcount"`
	UserPlayCount string `xml:"userplaycount"`
	Tracks        []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		Duration   string `xml:"duration"`
		Mbid       string `xml:"Mbid"`
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
	} `xml:"tracks>track"`
	TopTags []struct {
		Name string `xml:"name"`
		URL  string `xml:"url"`
	} `xml:"toptags>tag"`
	Wiki struct {
		Published string `xml:"published"`
		Summary   string `xml:"summary"`
		Content   string `xml:"content"`
	} `xml:"wiki"`
}

// AlbumGetTags <- album.getTags
type AlbumGetTags struct {
	XMLName xml.Name `xml:"tags"`
	Artist  string   `xml:"artist,attr"`
	Album   string   `xml:"album,attr"`
	Tags    []struct {
		Name string `xml:"name"`
		URL  string `xml:"url"`
	} `xml:"tag"`
}

// AlbumGetTopTags <- album.getTopTags
type AlbumGetTopTags struct {
	XMLName xml.Name `xml:"toptags"`
	Aritist string   `xml:"artist,attr"`
	Album   string   `xml:"album,attr"`
	Tags    []struct {
		Name  string `xml:"name"`
		Count string `xml:"count"`
		URL   string `xml:"url"`
	} `xml:"tag"`
}

//album.removeTag (empty)

// AlbumSearch <- album.search
type AlbumSearch struct {
	XMLName    xml.Name `xml:"results"`
	OpenSearch string   `xml:"opensearch,attr"`
	For        string   `xml:"for,attr"`
	Query      struct {
		Role        string `xml:"role,attr"`
		SearchTerms string `xml:"searchTrems,attr"`
		StartPage   int    `xml:"startPage,attr"`
	} `xml:"Query"`
	TotalResults int `xml:"totalResults"`
	StartIndex   int `xml:"startIndex"`
	ItemsPerPage int `xml:"itemsPerPage"`
	AlbumMatches []struct {
		Name   string `xml:"name"`
		Artist string `xml:"artist"`
		ID     string `xml:"id"`
		URL    string `xml:"url"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		Streamable string `xml:"streamable"`
		Mbid       string `xml:"mbid"`
	} `xml:"albummatches>album"`
}
