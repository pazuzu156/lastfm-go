package lastfm

const (
	// URIAPISecBase base api uri (secured)
	URIAPISecBase = "https://ws.audioscrobbler.com/2.0/"

	// URIAPIBase base api uri
	URIAPIBase = "http://ws.audioscrobbler.com/2.0/"

	// URIBrowserBase base browser api uri
	URIBrowserBase = "https://www.last.fm/api/auth/"
)

// P is used for adding parameters to API method calls.
type P map[string]interface{}

// API is the base Last.fm API struct for each endpoint.
type API struct {
	params  *apiParams
	Album   *albumAPI
	Artist  *artistAPI
	Chart   *chartAPI
	Geo     *getAPI
	Library *libraryAPI
	Tag     *tagAPI
	Track   *trackAPI
	User    *userAPI
}

type apiParams struct {
	apikey    string
	secret    string
	sk        string
	useragent string
}

// New creates a new Last.fm API instance.
func New(key, secret string) (api *API) {
	params := apiParams{key, secret, "", ""}
	api = &API{
		params:  &params,
		Album:   &albumAPI{&params},
		Artist:  &artistAPI{&params},
		Chart:   &chartAPI{&params},
		Geo:     &getAPI{&params},
		Library: &libraryAPI{&params},
		Tag:     &tagAPI{&params},
		Track:   &trackAPI{&params},
		User:    &userAPI{&params},
	}
	return
}

// SetSession sets API session key
func (api *API) SetSession(sessionkey string) {
	api.params.sk = sessionkey
}

// GetSessionKey returns API session key
func (api API) GetSessionKey() (sk string) {
	sk = api.params.sk
	return
}

// SetUserAgent sets user agent for API calls
func (api *API) SetUserAgent(useragent string) {
	api.params.useragent = useragent
}
