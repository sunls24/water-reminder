package wechatwork

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

const (
	companyId = "ww48d2720d9851f7af"
	secret    = "SYOdciKqJlKSkNXfjpeifKPa0NsKhBuq4oyq7to1wbI"
)

func TestNewApplication(t *testing.T) {
	_, err := NewApplication(companyId, secret)
	if err != nil {
		t.Fatal(errors.Wrap(err, "NewApplication"))
	}
}
