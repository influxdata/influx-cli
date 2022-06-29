package signin

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"syscall"

	"github.com/influxdata/influx-cli/v2/api"
	"golang.org/x/term"
)

func GetCookie(ctx context.Context, params api.ConfigParams, userPass string) (string, error) {
	bufUserPass, err := base64.StdEncoding.DecodeString(userPass)
	if err != nil {
		return "", err
	}

	userPass = strings.Trim(string(bufUserPass), "\x00")
	splitUserPass := strings.Split(userPass, ":")
	if len(splitUserPass) < 1 {
		return "", fmt.Errorf("bad config")
	}
	username := splitUserPass[0]
	var password string
	if len(splitUserPass) != 2 {
		fmt.Print("Please provide your password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}
		password = string(bytePassword)
		fmt.Println()
	} else {
		password = splitUserPass[1]
	}

	cfg := api.NewAPIConfig(params)
	client := api.NewAPIClient(cfg)
	ctx = context.WithValue(ctx, api.ContextBasicAuth, api.BasicAuth{
		UserName: username,
		Password: password,
	})
	res, err := client.SigninApi.PostSignin(ctx).ExecuteWithHttpInfo()
	if err != nil {
		return "", err
	}

	cookies := res.Cookies()
	if len(cookies) != 1 {
		return "", fmt.Errorf("failure getting session cookie: %w", err)
	}

	return cookies[0].Value, err
}
