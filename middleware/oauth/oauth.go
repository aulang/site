package oauth

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
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
	}

	user := ctx.User()
	if user != nil && user.GetAuthorization() == accessToken {
		return
	}

	user, err := o.obtainUser(accessToken)
	if err != nil {
		ctx.StopWithStatus(http.StatusUnauthorized)
	}

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
	client := http.Client{}

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
	if err != nil {
		log.Println("获取User失败", string(body))
		return user, err
	}

	user.AccessToken = accessToken

	return user, err
}
