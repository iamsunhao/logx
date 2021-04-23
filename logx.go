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
		err = os.Mkdir("log", 0666)
		if err != nil {
			log.Fatalln("创建 log 文件夹失败。", err)
		}
	} else {
		if !dirInfo.IsDir() {
			log.Fatalln("存在 log，但并不是目录。")
		}
	}
	return fmt.Sprintf("log/%s.log", time.Now().Format("20060102"))
}

//Info 创建前缀为INFO的日志
func Info(args ...interface{}) {
	file, err := os.OpenFile(getFilename(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalln("打开日志文件失败。", err)
	}
	logger := log.New(file, "INFO ", log.Ldate|log.Ltime)
	logger.Println(args...)
}

//Error 创建前缀为ERROR的日志
func Error(args ...interface{}) {
	file, err := os.OpenFile(getFilename(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalln("打开日志文件失败。", err)
	}
	logger := log.New(file, "ERROR ", log.Ldate|log.Ltime)
	logger.Fatalln(args...)
}

//Warning 创建前缀为WARNING的日志
func Warning(args ...interface{}) {
	file, err := os.OpenFile(getFilename(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalln("打开日志文件失败。", err)
	}
	logger := log.New(file, "WARNING ", log.Ldate|log.Ltime)
	logger.Println(args...)
}
