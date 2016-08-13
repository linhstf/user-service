package facebook

// GenerateUrl returns authenticated url of facebook
func GenerateUrl() string {
	return FbConfig.AuthCodeURL("state")
}
