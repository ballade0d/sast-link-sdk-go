package oauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
)

type Client struct {
	ClientID     string
	ClientSecret string
	HostName     string
	HttpClient   *http.Client
}

func NewClient(clientID, clientSecret string, hostName string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HostName:     hostName,
		HttpClient:   &http.Client{},
	}
}

func (c *Client) AccessToken(code string, redirectURI string, codeVerifier string) (*Token, error) {
	url := AccessToken.GetHttp(c.HostName)
	var reqBody bytes.Buffer
	mw := multipart.NewWriter(&reqBody)
	_ = mw.WriteField("code", code)
	_ = mw.WriteField("code_verifier", codeVerifier)
	_ = mw.WriteField("grant_type", GRANT_TYPE_AUTHORIZATION_CODE)
	_ = mw.WriteField("redirect_uri", redirectURI)
	_ = mw.WriteField("client_id", c.ClientID)
	_ = mw.WriteField("client_secret", c.ClientSecret)
	err := mw.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, &reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "multipart/form-data;boundary="+mw.Boundary())

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(respBody)
	if err != nil {
		return nil, err
	}

	token := &Response[Token]{}
	err = json.Unmarshal(respBody, token)
	if err != nil {
		return nil, err
	}
	if !token.Success {
		return nil, errors.New(token.ErrMsg)
	}
	return &token.Data, nil
}

func (c *Client) RefreshToken(refreshToken string) (*Token, error) {
	url := RefreshToken.GetHttp(c.HostName)
	var reqBody bytes.Buffer
	mw := multipart.NewWriter(&reqBody)
	_ = mw.WriteField("refresh_token", refreshToken)
	_ = mw.WriteField("grant_type", GRANT_TYPE_REFRESH_TOKEN)
	err := mw.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, &reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "multipart/form-data;boundary="+mw.Boundary())

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(respBody)
	if err != nil {
		return nil, err
	}

	token := &Response[Token]{}
	err = json.Unmarshal(respBody, token)
	if err != nil {
		return nil, err
	}
	if !token.Success {
		return nil, errors.New(token.ErrMsg)
	}
	return &token.Data, nil
}

func (c *Client) UserInfo(accessToken string) (*User, error) {
	url := UserInfo.GetHttp(c.HostName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(respBody)
	if err != nil {
		return nil, err
	}

	user := &Response[User]{}
	err = json.Unmarshal(respBody, user)
	if err != nil {
		return nil, err
	}
	if !user.Success {
		return nil, errors.New(user.ErrMsg)
	}
	return &user.Data, nil
}
