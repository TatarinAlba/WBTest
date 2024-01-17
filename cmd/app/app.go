package main

import (
	"github.com/TatarinAlba/WBTest/config"
	"github.com/TatarinAlba/WBTest/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	var cfg config.Config
	err := config.ParseConfig(&cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := app.Run(&cfg); err != nil {
		logrus.Errorf(err.Error())
	}
	logrus.Info("Successfully exited!")
}
