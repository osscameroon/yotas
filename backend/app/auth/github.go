package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubAccessToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func NewGithubAccessToken(code string) (*GithubAccessToken, error) {
	accessUrl := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		env.GithubClientId, env.GithubSecretKey, code,
	)

	req, _ := http.NewRequest(http.MethodPost, accessUrl, nil)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	var accessToken GithubAccessToken
	json.NewDecoder(res.Body).Decode(&accessToken)

	return &accessToken, nil
}

type GithubUser struct {
	Login            string `json:"login"`
	Id               string `json:"id"`
	AvatarUrl        string `json:"avatar_url"`
	Url              string `json:"url"`
	HtmlUrl          string `json:"html_url"`
	FollowingUrl     string `json:"following_url"`
	FollowersUrl     string `json:"followers_url"`
	GistsUrl         string `json:"gists_url"`
	StarredUrl       string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	OrganizationsUrl string `json:"organizations_url"`
	ReposUrl         string `json:"repos_url"`
	Name             string `json:"name"`
	Company          string `json:"company"`
	Blog             string `json:"blog"`
	Hireable         bool   `json:"hireable"`
	Bio              string `json:"bio"`
	TwitterUsername  string `json:"twitter_username"`
	PublicRepos      int    `json:"public_repos"`
	PublicGists      int    `json:"public_gists"`
	Followers        int    `json:"followers"`
	Following        int    `json:"following"`
}

func NewGithubUser(accessToken *GithubAccessToken) (*GithubUser, error) {
	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", accessToken.TokenType+" "+accessToken.AccessToken)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	var githubUser GithubUser
	json.NewDecoder(res.Body).Decode(&githubUser)

	return &githubUser, nil
}

type GithubEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}

func NewGithubEmail(accessToken *GithubAccessToken) (*GithubEmail, error) {
	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/user/emails", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", accessToken.TokenType+" "+accessToken.AccessToken)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	var emails []GithubEmail
	json.NewDecoder(res.Body).Decode(&emails)

	for _, value := range emails {
		if value.Primary {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("there is no primary email")
}

type GithubProfilePageData struct {
	User  GithubUser
	Email GithubEmail
}