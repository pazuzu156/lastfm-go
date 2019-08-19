package lastfm

////////////
// auth.x //
////////////

// AuthGetMobileSession <- auth.getMobileSession
type AuthGetMobileSession struct {
	Name       string `xml:"name"` //username
	Key        string `xml:"key"`  //session key
	Subscriber bool   `xml:"subscriber"`
}

// AuthGetToken <- auth.getToken
type AuthGetToken struct {
	Token string `xml:",chardata"`
}

// AuthGetSession <- auth.getSession
type AuthGetSession AuthGetMobileSession
