package main

import (
	"github.com/golang/glog"
	"github.com/josue/chalenge_golang/internal/app/api"
)

func main() {
	glog.Info("Starting service")
	api.LoadApiV1()
}
