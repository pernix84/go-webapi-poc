package models

type AuthUser struct {
	ID         int    `json:id`
	FullName   string `json:fullName`
	LoginCount int    `json:loginCount`
	OAuthId    string `json:oauthId`
}
