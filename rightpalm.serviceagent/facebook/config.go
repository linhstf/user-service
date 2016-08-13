package facebook

import (
	"fmt"

	"golang.org/x/oauth2"
)

// generate loginURL
var (
	ClientID        = "315627565448019"
	ClientSecret    = "92abb488a9349ead127b43d3fd3c3805"
	TestAccessToken = "EAAEfD8Tre1MBAEpDFL7VV8OC5VdCpgZBhURF7vEMC2UYZAbe4Pu1wlUg3c5ZCkacsviX9URaK0ajULvsId7LPBze3ZAWyOkH8VgZAg8ZANcuZCFkfiEum139xDpWBXu4y8hEZAZAmfXPmeFTLS4N8B3EnJRyG6aRcDQwZD"
	GraphUrl        = "https://graph.facebook.com/v2.7/"
	FbConfig        = &oauth2.Config{
		// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
		// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

		ClientID:     ClientID, // change this to yours
		ClientSecret: ClientSecret,
		RedirectURL:  "http://localhost:9000/Home/Login", // change this to your webserver adddress
		Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: fmt.Sprintf("%s/oauth/access_token", GraphUrl),
		},
	}
)
