package oauth

import (
	"crypto/tls"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type OAuth struct {
	profileUrl string
}

const (
	Bearer        = "Bearer "
	ACCESS_TOKEN  = "access_token"
	Authorization = "Authorization"
)

func New() iris.Handler {
	return (&OAuth{
		profileUrl: "https://aulang.cn/oauth/api/profile",
	}).Serve
}

func (o *OAuth) Serve(ctx iris.Context) {
	accessToken := o.getAccessToken(ctx)
	if accessToken == "" {
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	user := o.getSessionUser(ctx)
	if user != nil && user.GetAuthorization() == accessToken {
		ctx.Next()
		return
	}

	user, err := o.obtainUser(accessToken)
	if err != nil {
		log.Println("获取User失败：", err)
		ctx.StopWithStatus(http.StatusUnauthorized)
		return
	}

	o.setSessionUser(ctx, *user)
	ctx.SetUser(user)
	ctx.Next()
}

func (o *OAuth) getAccessToken(ctx iris.Context) string {
	accessToken := ctx.URLParam(ACCESS_TOKEN)
	authorization := ctx.GetHeader(Authorization)

	if authorization != "" {
		accessToken = strings.Replace(authorization, Bearer, "", 1)
	}

	return accessToken
}

func (o *OAuth) obtainUser(accessToken string) (user *SimpleUser, err error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{Transport: transport}

	req, err := http.NewRequest("GET", o.profileUrl, nil)
	if err != nil {
		return user, err
	}

	req.Header.Set(Authorization, "Bearer "+accessToken)

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
	if err != nil || user.ID == "" {
		log.Println("Profile接口调用失败！", string(body))
		return user, err
	}

	user.AccessToken = accessToken

	return user, err
}

func (o *OAuth) setSessionUser(ctx iris.Context, user SimpleUser) {
	session := sessions.Get(ctx)
	if session != nil {
		session.Set("SESSION_USER", user)
	}
}

func (o *OAuth) getSessionUser(ctx iris.Context) *SimpleUser {
	session := sessions.Get(ctx)

	if session != nil {
		data := session.Get("SESSION_USER")
		if data != nil {
			if user, ok := data.(SimpleUser); ok {
				return &user
			}
		}
	}

	return nil
}
