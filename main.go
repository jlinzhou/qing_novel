package main

import (
	"github.com/sirupsen/logrus"
	"qing_novel/server"
)

func main() {

	logrus.Info("run")

	server.StartServer()
}
