package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type OAuth struct {
	codeMode        bool
	codeKey         string
	accessTokenKey  string
	refreshTokenKey string
	expiresInKey    string
	tokenUrl        string
	profileUrl      string
	clientId        string
	clientSecret    string
	redirectUrl     string
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

func New() iris.Handler {
	oauth := OAuth{}

	return func(context *context.Context) {
		tokenBearer := context.GetHeader("Authorization")
		if tokenBearer == "" {
			// TODO 解析Token
			fmt.Println(tokenBearer)
		}

		if oauth.codeMode {
			code := oauth.getCode(context)
			if code == "" {

			}

			token, err := oauth.obtainAccessToken(code)
			if err != nil {
				// TODO 认证失败
			}

			user, err := oauth.obtainUser(token.AccessToken)
			if err != nil {
				// TODO 认证失败
			}

			// TODO 缓存用户信息
			fmt.Println(user.ID)
		} else {
			token := oauth.getAccessToken(context)
			if token != nil {
				// TODO 认证失败
			}

			user, err := oauth.obtainUser(token.AccessToken)
			if err != nil {
				// TODO 认证失败
			}

			// TODO 缓存用户信息
			fmt.Println(user.ID)
		}
	}
}

func (o *OAuth) getCode(context *context.Context) string {
	code := context.Params().Get(o.codeKey)

	return code
}

// 简化模式
func (o *OAuth) getAccessToken(context *context.Context) *AccessToken {
	accessToken := context.Params().Get(o.codeKey)

	if accessToken == "" {
		return nil
	}

	refreshToken := context.Params().Get(o.refreshTokenKey)
	expiresIn, err := context.Params().GetInt(o.expiresInKey)

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
	data.Set("redirect_uri", o.clientSecret)

	resp, err := http.PostForm(o.tokenUrl, data)
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
		return token, err
	}

	return token, err
}

func (o *OAuth) obtainUser(accessToken string) (user *User, err error) {
	client := http.Client{}

	req, err := http.NewRequest("POST", o.profileUrl, nil)
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

	err = json.Unmarshal(body, user)
	if err != nil {
		return user, err
	}

	return user, err
}
