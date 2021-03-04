package logx

import (
	"fmt"
	"log"
	"os"
	"time"
)

//getFilename 确定log文件夹存在，并返回要创建的log文件名及路径
func getFilename() string {
	dirInfo, err := os.Stat("log")
	if err != nil {
		err = os.Mkdir("log", os.ModeDir)
		if err != nil {
			log.Fatalln("Failed to make dir", err)
		}
	} else {
		if !dirInfo.IsDir() {
			log.Fatalln("log is exist but it is not a dir")
		}
	}
	return fmt.Sprintf("log/%s.log", time.Now().Format("20060102"))
}

//Info 创建前缀为INFO的日志
func Info(args ...interface{}) {
	file, err := os.OpenFile(getFilename(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger := log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(args...)
}

//Danger 创建前缀为ERROR的日志
func Danger(args ...interface{}) {
	file, err := os.OpenFile(getFilename(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger := log.New(file, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(args...)
}

//Warning 创建前缀为WARNING的日志
func Warning(args ...interface{}) {
	file, err := os.OpenFile(getFilename(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger := log.New(file, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(args...)
}
