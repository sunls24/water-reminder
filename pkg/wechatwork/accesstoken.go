package wechatwork

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
	"water-reminder/pkg/httpclient"
	"water-reminder/pkg/wechatwork/constant"
)

type AccessToken interface {
	CheckToken() error
}

type accessToken struct {
	app        *application
	token      string
	expiration int64
}

func NewAccessToken(app *application) (AccessToken, error) {
	log.Debug("NewAccessToken app:", *app)
	token := &accessToken{app: app}
	return token, token.CheckToken()
}

func (at accessToken) CheckToken() error {
	resp, err := httpclient.Get(fmt.Sprintf(constant.URLGetToken, at.app.companyId, at.app.secret))
	if err != nil {
		return errors.Wrap(err, "httpclient.Get")
	}
	at.token, err = resp.GetString(constant.KeyAccessToken)
	if err != nil {
		return err
	}
	expiresIn, err := resp.GetInt(constant.KeyExpiresIn)
	if err != nil {
		return err
	}
	at.expiration = time.Now().Unix() + int64(expiresIn)

	log.Debugf("%s: %s, expiration: %d", constant.KeyAccessToken, at.token, at.expiration)
	return nil
}
