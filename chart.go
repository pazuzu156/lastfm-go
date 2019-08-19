package lastfm

type chartAPI struct {
	params *apiParams
}

//chart.getTopArtists
func (api chartAPI) GetTopArtists(args map[string]interface{}) (result ChartGetTopArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopArtists") }()
	err = callGet("chart.gettopartists", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}

//chart.getTopTags
func (api chartAPI) GetTopTags(args map[string]interface{}) (result ChartGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopTags") }()
	err = callGet("chart.gettoptags", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}

//chart.getTopTracks
func (api chartAPI) GetTopTracks(args map[string]interface{}) (result ChartGetTopTracks, err error) {
	defer func() { appendCaller(err, "lastfm.Chart.GetTopTracks") }()
	err = callGet("chart.gettoptracks", api.params, args, &result, P{
		"plain": []string{"page", "limit"},
	})
	return
}
