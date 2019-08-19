package lastfm

type getAPI struct {
	params *apiParams
}

//geo.getTopArtists
func (api getAPI) GetTopArtists(args map[string]interface{}) (result GeoGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetTopArtists") }()
	err = callGet("geo.gettopartists", api.params, args, &result, P{
		"plain": []string{"country", "limit", "page"},
	})
	return
}

//geo.getTopTracks
func (api getAPI) GetTopTracks(args map[string]interface{}) (result GeoGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Geo.GetTopTracks") }()
	err = callGet("geo.gettoptracks", api.params, args, &result, P{
		"plain": []string{"country", "location", "limit", "page"},
	})
	return
}
