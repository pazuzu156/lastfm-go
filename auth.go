package lastfm

import (
	"net/url"
)

// Login : Mobile app style.
func (api *API) Login(username, password string) (err error) {
	defer func() { appendCaller(err, "lastfm.Login") }()

	var result AuthGetMobileSession
	args := P{"username": username, "password": password}
	if err = callPostWithoutSession("auth.getmobilesession", api.params, args, &result, P{
		"plain": []string{"username", "password"},
	}); err != nil {
		return
	}
	api.params.sk = result.Key
	//api.creds.username = result.Name
	return
}

// GetToken : Desktop app style.
func (api API) GetToken() (token string, err error) {
	defer func() { appendCaller(err, "lastfm.GetToken") }()

	var result AuthGetToken
	if err = callGet("auth.gettoken", api.params, nil, &result, P{}); err != nil {
		return
	}
	token = result.Token
	return
}

// GetAuthTokenURL returns auth token url.
func (api API) GetAuthTokenURL(token string) (uri string) {
	urlParams := url.Values{}
	urlParams.Add("api_key", api.params.apikey)
	urlParams.Add("token", token)
	uri = constructURL(URIBrowserBase, urlParams)
	return
}

// GetAuthRequestURL : Web app style
func (api API) GetAuthRequestURL(callback string) (uri string) {
	urlParams := url.Values{}
	urlParams.Add("api_key", api.params.apikey)
	if callback != "" {
		urlParams.Add("cb", callback)
	}
	uri = constructURL(URIBrowserBase, urlParams)
	return
}

// LoginWithToken : Desktop and Web app style
func (api *API) LoginWithToken(token string) (err error) {
	defer func() { appendCaller(err, "lastfm.LoginWithToken") }()

	var result AuthGetSession
	args := P{"token": token}
	if err = callPostWithoutSession("auth.getsession", api.params, args, &result, P{"plain": []string{"token"}}); err != nil {
		return
	}
	api.params.sk = result.Key
	//api.params.username = result.Name
	return
}
