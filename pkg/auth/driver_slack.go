package auth

import (
	"encoding/json"
	"gopkg.in/danilopolani/gocialite.v1"
	"io/ioutil"
	"net/http"

	"github.com/danilopolani/gocialite/structs"
	"golang.org/x/oauth2/slack"
)

const slackDriverName = "slack"

func init() {
	gocialite.RegisterNewDriver(slackDriverName, SlackDefaultScopes, SlackUserFn, slack.Endpoint, SlackAPIMap, SlackUserMap)
}

// Decode a json or return an error
func jsonDecode(js []byte) (map[string]interface{}, error) {
	var decoded map[string]interface{}
	if err := json.Unmarshal(js, &decoded); err != nil {
		return nil, err
	}

	return decoded, nil
}

// SlackUserMap is the map to create the User struct
var SlackUserMap = map[string]string{
	"real_name":      "FullName",
	"first_name":     "FirstName",
	"last_name":      "LastName",
	"email":          "Email",
	"image_original": "Avatar",
}

// SlackAPIMap is the map for API endpoints
var SlackAPIMap = map[string]string{
	"endpoint":     "https://slack.com/api",
	"userEndpoint": "/users.profile.get",
	"authEndpoint": "/auth.test",
}

// SlackUserFn is a callback to parse additional fields for User
var SlackUserFn = func(client *http.Client, u *structs.User) {
	// Get user ID
	req, err := client.Get(SlackAPIMap["endpoint"] + SlackAPIMap["authEndpoint"])
	if err != nil {
		return
	}

	defer req.Body.Close()
	res, _ := ioutil.ReadAll(req.Body)
	data, err := jsonDecode(res)
	if err != nil {
		return
	}

	u.ID = data["user_id"].(string)

	// Fetch other user information
	userInfo := u.Raw["profile"].(map[string]interface{})
	u.Username = userInfo["display_name"].(string)
	u.FullName = userInfo["real_name"].(string)
	if userInfo["first_name"] != nil {
		u.FirstName = userInfo["first_name"].(string)
	}
	if userInfo["last_name"] != nil {
		u.LastName = userInfo["last_name"].(string)
	}
	u.Email = userInfo["email"].(string)
	if userInfo["image_original"] != nil {
		u.Avatar = userInfo["image_original"].(string)
	}
}

// SlackDefaultScopes contains the default scopes
var SlackDefaultScopes = []string{"users.profile:read"}
