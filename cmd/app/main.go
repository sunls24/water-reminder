package main

import log "github.com/sirupsen/logrus"

func init() {
	// 设置日志打印格式
	format := new(log.TextFormatter)
	format.FullTimestamp = true
	format.TimestampFormat = "06-01-02 15:04:05"
	log.SetFormatter(format)
}
