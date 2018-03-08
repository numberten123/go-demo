package go_log

import (  
    "log"  
    "os"  
)

func Info(msg string, value interface{}){  
    fileName := "../log/info.log"  
    logFile,err  := os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
    defer logFile.Close()  
    if err != nil {  
        log.Fatalln("open file error !")  
    }  
    debugLog := log.New(logFile,"[Info]",log.LstdFlags) 
    debugLog.SetPrefix("[Info]")  
    debugLog.Println(msg, value) 
}

func Error(msg string, value interface{}){  
    fileName := "../log/error.log"  
    logFile,err  := os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
    defer logFile.Close()  
    if err != nil {  
        log.Fatalln("open file error !")  
    }  
    debugLog := log.New(logFile,"[Error]",log.LstdFlags)
    debugLog.Println(msg, value)
}