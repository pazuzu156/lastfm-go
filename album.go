package lastfm

type albumAPI struct {
	params *apiParams
}

//album.addTags
func (api albumAPI) AddTags(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Album.AddTags") }()
	err = callPost("album.addtags", api.params, args, nil, P{
		"plain": []string{"artist", "album", "tags"},
	})
	return
}

//album.getInfo
func (api albumAPI) GetInfo(args map[string]interface{}) (result AlbumGetInfo, err error) {
	defer func() { appendCaller(err, "lastfm.Album.GetInfo") }()
	err = callGet("album.getinfo", api.params, args, &result, P{
		"plain": []string{"artist", "album", "mbid", "autocorrect", "username", "lang"},
	})
	return
}

//album.getTags
func (api albumAPI) GetTags(args map[string]interface{}) (result AlbumGetTags, err error) {
	defer func() { appendCaller(err, "lastfm.Album.GetTags") }()
	if _, ok := args["user"]; !ok && api.params.sk != "" {
		err = callPost("album.gettags", api.params, args, &result, P{
			"plain": []string{"artist", "album", "mbid", "autocorrect"},
		})
	} else {
		err = callGet("album.gettags", api.params, args, &result, P{
			"plain": []string{"artist", "album", "mbid", "autocorrect", "user"},
		})
	}
	return
}

//album.getTopTags
func (api albumAPI) GetTopTags(args map[string]interface{}) (result AlbumGetTopTags, err error) {
	defer func() { appendCaller(err, "lastfm.Album.GetTopTags") }()
	err = callGet("album.gettoptags", api.params, args, &result, P{
		"plain": []string{"artist", "album", "autocorrect", "mbid"},
	})
	return
}

//album.removeTag
func (api albumAPI) RemoveTag(args map[string]interface{}) (err error) {
	defer func() { appendCaller(err, "lastfm.Album.RemoveTag") }()
	err = callPost("album.removetag", api.params, args, nil, P{
		"plain": []string{"artist", "album", "tag"},
	})
	return
}

//album.search
func (api albumAPI) Search(args map[string]interface{}) (result AlbumSearch, err error) {
	defer func() { appendCaller(err, "lastfm.Album.Search") }()
	err = callGet("album.search", api.params, args, &result, P{
		"plain": []string{"limit", "page", "album"},
	})
	return
}
