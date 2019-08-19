package lastfm

import "encoding/xml"

// TagGetInfo <- tag.getinfo
type TagGetInfo struct {
	XMLName    xml.Name `xml:"tag"`
	Name       string   `xml:"name"`
	URL        string   `xml:"url"`
	Reach      string   `xml:"reach"`
	Taggings   string   `xml:"taggings"`
	Streamable string   `xml:"streamable"`
	Wiki       struct {
		Published string `xml:"published"`
		Summary   string `xml:"summary"`
		Content   string `xml:"content"`
	} `xml:"wiki"`
}

// TagGetSimilar <- tag.getSimilar
type TagGetSimilar struct {
	XMLName xml.Name `xml:"similartags"`
	Tag     string   `xml:"tag,attr"`
	Tags    []struct {
		Name       string `xml:"name"`
		URL        string `xml:"url"`
		Streamable string `xml:"streamable"`
	} `xml:"tag"`
}

// TagGetTopAlbums <- tag.getTopAlbums
type TagGetTopAlbums struct {
	XMLName    xml.Name `xml:"topalbums"`
	Tag        string   `xml:"tag,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Albums     []struct {
		Rank   string `xml:"rank,attr"`
		Name   string `xml:"name"`
		URL    string `xml:"url"`
		Artist struct {
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

// TagGetTopArtists <- tag.getTopArtists
type TagGetTopArtists struct {
	XMLName xml.Name `xml:"topartists"`
	Tag     string   `xml:"tag,attr"`
	//Total      string   `xml:"total,attr"`
	//Page       string   `xml:"page,attr"`
	//PerPage    string   `xml:"perPage,attr"`
	//TotalPages string   `xml:"totalPages"`
	Artists []struct {
		Rank       string `xml:"rank,attr"`
		Name       string `xml:"name"`
		URL        string `xml:"url"`
		Streamable string `xml:"streamable"`
		Images     []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
	} `xml:"artist"`
}

// TagGetTopTags <- tag.getTopTags
type TagGetTopTags struct {
	XMLName xml.Name `xml:"toptags"`
	Tags    []struct {
		Name  string `xml:"name"`
		Count string `xml:"count"`
		URL   string `xml:"url"`
	} `xml:"tag"`
}

// TagGetTopTracks <- tag.getTopTracks
type TagGetTopTracks struct {
	XMLName    xml.Name `xml:"toptracks"`
	Tag        string   `xml:"tag,attr"`
	Total      int      `xml:"total,attr"`
	Page       int      `xml:"page,attr"`
	PerPage    int      `xml:"perPage,attr"`
	TotalPages int      `xml:"totalPages,attr"`
	Tracks     []struct {
		Rank       string `xml:"rank,attr"`
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
	} `xml:"track"`
}

// TagGetWeeklyChartList <- tag.getWeeklyChartList
type TagGetWeeklyChartList struct {
	XMLName xml.Name `xml:"weeklychartlist"`
	Tag     string   `xml:"tag,attr"`
	Charts  []struct {
		From string `xml:"from,attr"`
		To   string `xml:"to,attr"`
	} `xml:"chart"`
}
