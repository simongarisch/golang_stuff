// go run logrus_example.go
// https://github.com/sirupsen/logrus/issues/156
// https://github.com/sirupsen/logrus/issues/483
// https://stackoverflow.com/questions/52899535/using-logger-configs-in-multi-packages-best-practice-for-golang-productive
// https://www.xspdf.com/resolution/52899535.html
// https://www.loggly.com/blog/think-differently-about-what-to-log-in-go-best-practices-examined/
// https://adeshinahassan.com/posts/designing-and-understanding-logging-system-in-go-application/
package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const logFile = "logrus.log"

var logger = log.New() // create the logger

func log2file() {
	// with Json Formatter
	logger.Formatter = &log.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	//defer file.Close()
	logger.SetOutput(file)
}

func main() {
	log2file()
	logger.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	logWithFields()
	anotherExample()
}

func logWithFields() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func anotherExample() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
