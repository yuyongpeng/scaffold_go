package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"scaffold_go/conf"
)

var Log *logrus.Logger = logrus.New()

// Create a new instance of the logger. You can have any number of instances.
func New()(lg *logrus.Logger){
	log := logrus.New()
	G := conf.GetConf()
	//log.SetOutput(os.Stdout)
	file, err := os.OpenFile(G.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	fmt.Println(G.Formater)
	if G.Formater == "text"{
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	if G.Formater == "json"{
		log.SetFormatter(&logrus.TextFormatter{})
	}
	log.SetFormatter(&logrus.TextFormatter{})

	loglevel := G.Loglevel

	if loglevel == "FatalLevel"{
		log.SetLevel(logrus.FatalLevel)
	}
	if loglevel == "ErrorLevel"{
		log.SetLevel(logrus.ErrorLevel)
	}
	if loglevel == "WarnLevel"{
		log.SetLevel(logrus.WarnLevel)
	}
	if loglevel == "InfoLevel"{
		log.SetLevel(logrus.InfoLevel)
	}
	if loglevel == "DebugLevel"{
		log.SetLevel(logrus.DebugLevel)
	}
	if loglevel == "TraceLevel"{
		log.SetLevel(logrus.TraceLevel)
	}

	return log
}

//func main() {
//	// The API for setting attributes is a little different than the package level
//	// exported logger. See Godoc.
//	log.Out = os.Stdout
//
//	// You could set this to any `io.Writer` such as a file
//	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
//	// if err == nil {
//	//  log.Out = file
//	// } else {
//	//  log.Info("Failed to log to file, using default stderr")
//	// }
//
//	log.WithFields(logrus.Fields{
//		"animal": "walrus",
//		"size":   10,
//	}).Info("A group of walrus emerges from the ocean")
//}