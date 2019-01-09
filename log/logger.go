package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"scaffold_go/conf"
)
// Create a new instance of the logger. You can have any number of instances.

func New()(lg *logrus.Logger){
	var log = logrus.New()
	G := conf.GetConf()
	file, err := os.OpenFile(G.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	if G.Formater == "text"{
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	if G.Formater == "json"{
		log.SetFormatter(&logrus.TextFormatter{})
	}
	log.SetFormatter(&logrus.TextFormatter{})
	//log.SetLevel(logrus.WarnLevel)
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