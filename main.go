package main

import (
	logConfig "github.com/mfirmanakbar/go-logging/log"
	log "github.com/sirupsen/logrus"
)

func main() {
	logConfig.InitializeLogging("/tmp/test1.log")
	log.Println("Something Happened ...")
	log.Fatalf("What Happened ?")
}
