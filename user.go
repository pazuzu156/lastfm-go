package lastfm

type userAPI struct {
	params *apiParams
}

//user.getArtistTracks
func (api userAPI) GetArtistTracks(args map[string]interface{}) (result UserGetArtistTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetArtistTracks") }()
	err = callGet("user.getartisttracks", api.params, args, &result, P{
		"plain": []string{"user", "artist", "startTimeStamp", "page", "endTimeStamp"},
	})
	return
}

//user.getFriends
func (api userAPI) GetFriends(args map[string]interface{}) (result UserGetFriends, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetFriends") }()
	err = callGet("user.getfriends", api.params, args, &result, P{
		"plain": []string{"user", "recenttracks", "limit", "page"},
	})
	return
}

//user.getInfo
func (api userAPI) GetInfo(args map[string]interface{}) (result UserGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetInfo") }()
	if _, ok := args["user"]; !ok && api.params.sk != "" {
		err = callPost("user.getinfo", api.params, args, &result, P{})
	} else {
		err = callGet("user.getinfo", api.params, args, &result, P{
			"plain": []string{"user"},
		})
	}
	return
}

//user.getLovedTracks
func (api userAPI) GetLovedTracks(args map[string]interface{}) (result UserGetLovedTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetLovedTracks") }()
	err = callGet("user.getlovedtracks", api.params, args, &result, P{
		"plain": []string{"user", "limit", "page"},
	})
	return
}

//user.getPersonalTags
func (api userAPI) GetPersonalTags(args map[string]interface{}) (result UserGetPersonalTags, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetPersonalTags") }()
	err = callGet("user.getPersonalTags", api.params, args, &result, P{
		"plain": []string{"user", "tag", "taggingtype", "limit", "page"},
	})
	return
}

//user.getRecentTracks
func (api userAPI) GetRecentTracks(args map[string]interface{}) (result UserGetRecentTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetRecentTracks") }()
	err = callGet("user.getrecenttracks", api.params, args, &result, P{
		"plain": []string{"user", "limit", "page", "from", "extended", "to"},
	})
	return
}

//user.getTopAlbums
func (api userAPI) GetTopAlbums(args map[string]interface{}) (result UserGetTopAlbums, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopAlbums") }()
	err = callGet("user.gettopalbums", api.params, args, &result, P{
		"plain": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getTopArtists
func (api userAPI) GetTopArtists(args map[string]interface{}) (result UserGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopArtists") }()
	err = callGet("user.gettopartists", api.params, args, &result, P{
		"plain": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getTopTags
func (api userAPI) GetTopTags(args map[string]interface{}) (result UserGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopTags") }()
	err = callGet("user.gettoptags", api.params, args, &result, P{
		"plain": []string{"user", "limit"},
	})
	return
}

//user.getTopTracks
func (api userAPI) GetTopTracks(args map[string]interface{}) (result UserGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetTopTracks") }()
	err = callGet("user.gettoptracks", api.params, args, &result, P{
		"plain": []string{"user", "period", "limit", "page"},
	})
	return
}

//user.getWeeklyAlbumChart
func (api userAPI) GetWeeklyAlbumChart(args map[string]interface{}) (result UserGetWeeklyAlbumChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyAlbumChart") }()
	err = callGet("user.getweeklyalbumchart", api.params, args, &result, P{
		"plain": []string{"user", "from", "to"},
	})
	return
}

//user.getWeeklyArtistChart
func (api userAPI) GetWeeklyArtistChart(args map[string]interface{}) (result UserGetWeeklyArtistChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyArtistChart") }()
	err = callGet("user.getweeklyartistchart", api.params, args, &result, P{
		"plain": []string{"user", "from", "to"},
	})
	return
}

//user.getWeeklyChartList
func (api userAPI) GetWeeklyChartList(args map[string]interface{}) (result UserGetWeeklyChartList, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyChartList") }()
	err = callGet("user.getweeklychartlist", api.params, args, &result, P{
		"plain": []string{"user"},
	})
	return
}

//user.getWeeklyTrackChart
func (api userAPI) GetWeeklyTrackChart(args map[string]interface{}) (result UserGetWeeklyTrackChart, err error) {
	defer func() { appendCaller(err, "lastfm.User.GetWeeklyTrackChart") }()
	err = callGet("user.getweeklytrackchart", api.params, args, &result, P{
		"plain": []string{"user", "from", "to"},
	})
	return
}
