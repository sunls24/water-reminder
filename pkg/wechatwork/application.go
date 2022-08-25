package wechatwork

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Application interface {
}

type application struct {
	companyId string
	secret    string

	token AccessToken
}

func NewApplication(companyId, secret string) (Application, error) {
	log.Debugf("NewApplication companyId: %s, secret: %s", companyId, secret)
	if len(companyId) == 0 {
		return nil, errors.New("企业 ID 不能为空")
	}
	if len(secret) == 0 {
		return nil, errors.New("应用 Secret 不能为空")
	}
	app := &application{companyId: companyId, secret: secret}
	token, err := NewAccessToken(app)
	if err != nil {
		return nil, errors.Wrap(err, "NewAccessToken")
	}
	app.token = token
	return app, nil
}
