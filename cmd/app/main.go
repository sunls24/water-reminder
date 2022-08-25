package main

import (
	log "github.com/sirupsen/logrus"
	"water-reminder/config"
	"water-reminder/pkg/wechatwork"
)

func init() {
	// 设置日志打印格式
	format := new(log.TextFormatter)
	format.FullTimestamp = true
	format.TimestampFormat = "06-01-02 15:04:05"
	log.SetFormatter(format)
}

func main() {
	cfg := config.NewConfig()
	app, err := wechatwork.NewApplication(cfg.CompanyId, cfg.Secret, cfg.AgentId)
	if err != nil {
		log.Fatal(err)
	}
	if err = app.SendMessage(wechatwork.NewTextMessage("first message ~")); err != nil {
		log.Fatal(err)
	}
}
