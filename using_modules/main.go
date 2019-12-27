package main

// go mod init

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.Info("this is some info")
	log.Warn("a warning")
	//log.Fatal("This is bad...")
}
