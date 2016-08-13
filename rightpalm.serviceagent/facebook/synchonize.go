package facebook

import (
	"fmt"
	"net/http"

	"rightpalm/rightpalm.serviceagent/utils"
)

func Synchonize(accessToken string) string {
	url := fmt.Sprintf("%s/me?access_token=%s", GraphUrl, accessToken)
	res, err := http.Get(url)
	if err != nil {
		return ""
	}

	return utils.ReadHttpBody(res)
}

func DebugToken() string {
	url := fmt.Sprintf("%s/debug_token?input_token=%s&access_token=%s", GraphUrl, TestAccessToken, ClientSecret)
	res, err := http.Get(url)
	if err != nil {
		return ""
	}

	return utils.ReadHttpBody(res)
}
