package lastfm

type artistAPI struct {
	params *apiParams
}

//artist.addTags
func (api artistAPI) AddTags(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Artist.AddTags") }()
	err = callPost("artist.addtags", api.params, args, nil, P{
		"plain": []string{"artist", "tags"},
	})
	return
}

//artist.getCorrection
func (api artistAPI) GetCorrection(args map[string]interface{}) (result ArtistGetCorrection, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetCorrection") }()
	err = callGet("artist.getcorrection", api.params, args, &result, P{
		"plain": []string{"artist"},
	})
	return
}

//artist.getInfo
func (api artistAPI) GetInfo(args map[string]interface{}) (result ArtistGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetInfo") }()
	if _, ok := args["username"]; !ok && api.params.sk != "" {
		err = callPost("artist.getinfo", api.params, args, &result, P{
			"plain": []string{"artist", "mbid", "lang", "autocorrect"},
		})
	} else {
		err = callGet("artist.getinfo", api.params, args, &result, P{
			"plain": []string{"artist", "mbid", "lang", "autocorrect", "username"},
		})
	}
	return
}

//artist.getSimilar
func (api artistAPI) GetSimilar(args map[string]interface{}) (result ArtistGetSimilar, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetSimilar") }()
	err = callGet("artist.getsimilar", api.params, args, &result, P{
		"plain": []string{"limit", "artist", "autocorrect", "mbid"},
	})
	return
}

//artist.getTags
func (api artistAPI) GetTags(args map[string]interface{}) (result ArtistGetTags, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTags") }()
	if _, ok := args["user"]; !ok && api.params.sk != "" {
		err = callGet("artist.gettags", api.params, args, &result, P{
			"plain": []string{"artist", "mbid", "user", "autocorrect"},
		})
	} else {
		err = callGet("artist.gettags", api.params, args, &result, P{
			"plain": []string{"artist", "mbid", "user", "autocorrect"},
		})
	}
	return
}

//artist.getTopAlbums
func (api artistAPI) GetTopAlbums(args map[string]interface{}) (result ArtistGetTopAlbums, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTopAlbums") }()
	err = callGet("artist.gettopalbums", api.params, args, &result, P{
		"plain": []string{"artist", "mbid", "autocorrect", "page", "limit"},
	})
	return
}

//artist.getTopTags
func (api artistAPI) GetTopTags(args map[string]interface{}) (result ArtistGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTopTags") }()
	err = callGet("artist.gettoptags", api.params, args, &result, P{
		"plain": []string{"artist", "mbid", "autocorrect"},
	})
	return
}

//artist.getTopTracks
func (api artistAPI) GetTopTracks(args map[string]interface{}) (result ArtistGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.GetTopTracks") }()
	err = callGet("artist.gettoptracks", api.params, args, &result, P{
		"plain": []string{"artist", "mbid", "autocorrect", "page", "limit"},
	})
	return
}

//artist.removeTag
func (api artistAPI) RemoveTag(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Artist.RemoveTag") }()
	err = callPost("artist.removetag", api.params, args, nil, P{
		"plain": []string{"artist", "tag"},
	})
	return
}

//artist.search
func (api artistAPI) Search(args map[string]interface{}) (result ArtistSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Artist.Search") }()
	err = callGet("artist.search", api.params, args, &result, P{
		"plain": []string{"limit", "page", "artist"},
	})
	return
}
