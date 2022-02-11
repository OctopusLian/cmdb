/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-02-11 13:08:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-02-11 13:11:28
 */
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"agent/config"
	"agent/ens"

	"agent/plugins"
	_ "agent/plugins/init"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)
	req.Debug = false

	configReader := viper.New()
	configReader.SetConfigName("agent")
	configReader.SetConfigType("yaml")
	configReader.AddConfigPath("etc/")

	err := configReader.ReadInConfig()
	if err != nil {
		logrus.Error("读取配置出错:", err)
		os.Exit(-1)
	}

	gconf, err := config.NewConfig(configReader)
	if err != nil {
		logrus.Error("读取配置出错:", err)
		os.Exit(-1)
	}
	defer func() {
		os.Remove(gconf.PidFile)
	}()
	log, err := os.OpenFile(gconf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		logrus.Error("打开日志文件出错:", err)
		os.Exit(-1)
	}
	defer func() {
		log.Close()
	}()

	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{})
	// logrus.SetOutput(log)
	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent启动")

	plugins.DefaultManager.Init(gconf)

	ens.NewENS(gconf).Start()
	plugins.DefaultManager.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logrus.WithFields(logrus.Fields{
		"PID":  gconf.PID,
		"UUID": gconf.UUID,
	}).Info("Agent退出")
}
