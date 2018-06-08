package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/services"
	"os"
	"log"
)

func main() {
	file:=setupLogFile()
	defer file.Close()
	dataBase.Connect()
	services.Start()
}

func setupLogFile()  *os.File{
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	println("All errors will be in the log.txt. Read it if you think that something is wrong.")
	log.Println("Recording of the log file has started...")
	return logFile
}
