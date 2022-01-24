package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Logger struct{
	Trace *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

var myLogger Logger

func main(){
	// printLog();
	// logOnFile();
	Test_Custom_Logger_File();
	logInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	myLogger.Info.Println("Starting the application...")
	myLogger.Trace.Println("Something noteworthy happened")
	myLogger.Warn.Println("There is something you should know about")
	myLogger.Error.Println("Something went wrong")
}

func printLog(){
	log.Println("Logging") // 출력 : 2022/01/13 15:34:27 Logging

	log.SetFlags(0)
	log.Println("Logging2") // 출력 : Logging2

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile) // 출력 : INFO : 2022/01/13 15:39:10 C:/Users/antline/Documents/TGM/golang/study/src/220113_logging.go:13: Logging
	log.SetPrefix("INFO : ")
	log.Println("Logging")
}

func logOnFile(){
	log.SetFlags( log.Ldate | log.Ltime )
	log.SetPrefix("")
	logFile, err := os.OpenFile("logfile.txt",os.O_CREATE | os.O_WRONLY | os.O_APPEND,0666)

	if err!=nil{
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Print("test msg")
	log.Println("End of Programe")
}

var logger *log.Logger

func Test_Custem_Logger() {
	logger = log.New(os.Stdout, "INFO : ",log.LstdFlags)
	logger.Println("Logging")
}

func Test_Custom_Logger_File(){
	// 로그 파일 오픈
	logFile, err := os.OpenFile("custom_Logger.txt",os.O_CREATE | os.O_WRONLY | os.O_APPEND,0666)
	if err!=nil{
		panic(err)
	}
	defer logFile.Close()

	logger = log.New(logFile,"INFO: ",log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("test msg")
	log.Println("End of Programe")
}

func logInit(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	myLogger.Trace = log.New(traceHandle, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile)
	myLogger.Info = log.New(infoHandle, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	myLogger.Warn = log.New(warningHandle, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	myLogger.Error = log.New(errorHandle, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}