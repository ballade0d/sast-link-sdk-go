package oauth

import "container/list"

type Response[T any] struct {
	Data    T      `json:"Data"`
	ErrCode int    `json:"ErrCode"`
	ErrMsg  string `json:"ErrMsg"`
	Success bool   `json:"Success"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type Badge struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type User struct {
	UserId   string    `json:"userId"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Badge    Badge     `json:"badge"`
	Bio      string    `json:"bio"`
	Dep      string    `json:"dep"`
	Hide     list.List `json:"hide"`
	Link     list.List `json:"link"`
	Nickname string    `json:"nickname"`
	Org      string    `json:"org"`
}
