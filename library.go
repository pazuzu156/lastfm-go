package lastfm

type libraryAPI struct {
	params *apiParams
}

//library.getArtists
func (api libraryAPI) GetArtists(args map[string]interface{}) (result LibraryGetArtists, err error) {
	defer func() { appendCaller(err, "lastfm.Library.GetArtists") }()
	err = callGet("library.getartists", api.params, args, &result, P{
		"plain": []string{"user", "limit", "page"},
	})
	return
}
