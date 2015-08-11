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
		panic(err)
	}

	fmt.Println(fbx.DownloadsStats())
}
