package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	. "os"
)

func InitializeLogging(logFile string) {

	var file, err = OpenFile(logFile, O_RDWR|O_CREATE|O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)

	//log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})
}
