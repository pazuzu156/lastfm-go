package lastfm

type trackAPI struct {
	params *apiParams
}

//track.addTags
func (api trackAPI) AddTags(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.AddTags") }()
	err = callPost("track.addtags", api.params, args, nil, P{
		"plain": []string{"artist", "track", "tags"},
	})
	return
}

//track.getCorrection
func (api trackAPI) GetCorrection(args map[string]interface{}) (result TrackGetCorrection, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Correction") }()
	err = callGet("track.getcorrection", api.params, args, &result, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//track.getInfo
func (api trackAPI) GetInfo(args map[string]interface{}) (result TrackGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetInfo") }()
	err = callGet("track.getinfo", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "username", "autocorrect"},
	})
	return
}

//track.getSimilar
func (api trackAPI) GetSimilar(args map[string]interface{}) (result TrackGetSimilar, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetSimilar") }()
	err = callGet("track.getsimilar", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "limit", "autocorrect"},
	})
	return
}

//track.getTags
func (api trackAPI) GetTags(args map[string]interface{}) (result TrackGetTags, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetTags") }()
	if _, ok := args["users"]; !ok && api.params.sk != "" {
		err = callPost("track.gettags", api.params, args, &result, P{
			"plain": []string{"artist", "track", "mbid", "autocorrect"},
		})
	} else {
		err = callGet("track.gettags", api.params, args, &result, P{
			"plain": []string{"artist", "track", "mbid", "user", "autocorrect"},
		})
	}
	return
}

//track.getTopTags
func (api trackAPI) GetTopTags(args map[string]interface{}) (result TrackGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Track.GetTopTags") }()
	err = callGet("track.gettoptags", api.params, args, &result, P{
		"plain": []string{"artist", "track", "mbid", "autocorrect"},
	})
	return
}

//track.love
func (api trackAPI) Love(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.Love") }()
	err = callPost("track.love", api.params, args, nil, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//track.removeTag
func (api trackAPI) RemoveTag(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.RemoveTag") }()
	err = callPost("track.removetag", api.params, args, nil, P{
		"plain": []string{"artist", "track", "tag"},
	})
	return
}

//track.scrobble
func (api trackAPI) Scrobble(args map[string]interface{}) (result TrackScrobble, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Scrobble") }()
	err = callPost("track.scrobble", api.params, args, &result, P{
		"indexing": []string{"artist", "track", "timestamp", "album", "context", "streamId", "chosenByUser", "trackNumber", "mbid", "albumArtist", "duration"},
	})
	return
}

//track.search
func (api trackAPI) Search(args map[string]interface{}) (result TrackSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Track.Search") }()
	err = callGet("track.search", api.params, args, &result, P{
		"plain": []string{"artist", "track", "limit", "page"},
	})
	return
}

//track.unlove
func (api trackAPI) UnLove(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Track.UnLove") }()
	err = callPost("track.unlove", api.params, args, nil, P{
		"plain": []string{"artist", "track"},
	})
	return
}

//track.updateNowPlaying
func (api trackAPI) UpdateNowPlaying(args map[string]interface{}) (result TrackUpdateNowPlaying, err error) {
	defer func() { appendCaller(err, "lastfm.Track.UpdateNowPlaying") }()
	err = callPost("track.updatenowplaying", api.params, args, &result, P{
		"plain": []string{"artist", "track", "album", "trackNumber", "context", "mbid", "duration", "albumArtist"},
	})
	return
}
