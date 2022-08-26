package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
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
	shanghai, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}

	app, err := wechatwork.NewApplication(cfg.CompanyId, cfg.Secret, cfg.AgentId)
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 0-120 minute
		<-time.After(time.Duration(rand.Intn(120)) * time.Minute)
		now := time.Now().In(shanghai).Format("06-01-02 15:04:05")
		if err = app.SendMessage(wechatwork.NewTextMessage(fmt.Sprintf("当前时间: %v", now))); err != nil {
			log.Fatal(err)
		}
	}

}
