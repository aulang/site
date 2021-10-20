package oauth

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

type OAuth struct {
	profileUrl string
}

type Profile struct {
	Id       string `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

const (
	Bearer        = "Bearer "
	AccessToken   = "access_token"
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

	user, err := o.obtainUser(accessToken)
	if err != nil {
		log.Println("获取User失败：", err)
		ctx.StopWithError(http.StatusUnauthorized, err)
		return
	}

	err = ctx.SetUser(user)
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Next()
}

func (o *OAuth) getAccessToken(ctx iris.Context) string {
	accessToken := ctx.URLParam(AccessToken)
	authorization := ctx.GetHeader(Authorization)

	if authorization != "" {
		accessToken = strings.TrimPrefix(authorization, Bearer)
	}

	return accessToken
}

func (o *OAuth) obtainUser(accessToken string) (user *iris.SimpleUser, err error) {
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

	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	profile := Profile{}
	err = json.Unmarshal(body, &profile)
	if err != nil || user.ID == "" {
		log.Println("Profile接口调用失败！", string(body))
		return user, err
	}

	user.Authorization = accessToken
	user.AuthorizedAt = time.Now()

	user.ID = profile.Id
	user.Username = profile.Nickname

	return user, err
}
