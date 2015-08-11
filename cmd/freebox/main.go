package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/moul/go-freebox"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	fbx := freebox.New()

	err := fbx.Connect()
	if err != nil {
		logrus.Fatalf("fbx.Connect(): %v", err)
	}

	err = fbx.Authorize()
	if err != nil {
		logrus.Fatalf("fbx.Authorize(): %v", err)
	}

	err = fbx.Login()
	if err != nil {
		logrus.Fatalf("fbx.Login(): %v", err)
	}

	stats, err := fbx.DownloadsStats()
	if err != nil {
		logrus.Fatalf("fbx.DownloadsStats(): %v", err)
	}
	fmt.Println(stats)
}
