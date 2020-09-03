package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/aulang/site/util"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type OAuth struct {
	clientId        string
	clientSecret    string
	redirectUrl     string // https://aulang.cn/site/admin/login
	codeMode        bool   // code/token
	codeKey         string // code
	accessTokenKey  string // access_token
	refreshTokenKey string // refresh_token
	expiresInKey    string // expires_in
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type User struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	AccessToken
}

const (
	Bearer        = "Bearer "
	Authorization = "Authorization"

	TokenUrl     = "https://aulang.cn/oauth/token"
	ProfileUrl   = "https://aulang.cn/oauth/api/profile"
	AuthorizeUrl = "https://aulang.cn/oauth/authorize?client_id=%s&response_type=%s&state=%s&redirect_uri=%s"
)

func New(clientId, clientSecret string) *OAuth {
	return &OAuth{
		clientId:        clientId,
		clientSecret:    clientSecret,
		redirectUrl:     "https://aulang.cn/site/admin/login",
		codeMode:        true,
		codeKey:         "code",
		accessTokenKey:  "access_token",
		refreshTokenKey: "refresh_token",
		expiresInKey:    "expires_in",
	}
}

func (o *OAuth) AuthorizeUrl() string {
	responseType := "code"
	if !o.codeMode {
		responseType = "token"
	}

	state := util.RandString(6)
	redirectUrl := url.QueryEscape(o.redirectUrl)

	return fmt.Sprintf(AuthorizeUrl, o.clientId, responseType, state, redirectUrl)
}

func (o *OAuth) Token(context *context.Context) *AccessToken {
	if o.codeMode {
		code := o.getCode(context)
		if code == "" {
			return nil
		}

		token, err := o.obtainAccessToken(code)
		if err != nil {
			log.Println("获取Token失败！", err)
		}

		return token
	} else {
		return o.getAccessToken(context)
	}
}

func (o *OAuth) User(token *AccessToken) (*User, error) {
	return o.obtainUser(token.AccessToken)
}

func (o *OAuth) getCode(context *context.Context) string {
	code := context.URLParam(o.codeKey)

	return code
}

// 简化模式
func (o *OAuth) getAccessToken(context *context.Context) *AccessToken {
	accessToken := context.URLParam(o.codeKey)

	if accessToken == "" {
		return nil
	}

	refreshToken := context.URLParam(o.refreshTokenKey)
	expiresIn, err := context.URLParamInt(o.expiresInKey)

	if err != nil {
		// 30分钟
		expiresIn = 30 * 60
	}

	return &AccessToken{AccessToken: accessToken, RefreshToken: refreshToken, ExpiresIn: expiresIn}
}

// 认证码模式
func (o *OAuth) obtainAccessToken(code string) (token *AccessToken, err error) {
	data := url.Values{}
	data.Set("client_id", o.clientId)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_secret", o.clientSecret)
	data.Set("redirect_uri", o.redirectUrl)

	resp, err := http.PostForm(TokenUrl, data)
	if err != nil {
		return token, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	err = json.Unmarshal(body, &token)
	if err != nil {
		log.Println("获取Token失败", string(body))
		return token, err
	}

	return token, err
}

func (o *OAuth) obtainUser(accessToken string) (user *User, err error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", ProfileUrl, nil)
	if err != nil {
		return user, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return user, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("获取User失败", string(body))
		return user, err
	}

	return user, err
}
