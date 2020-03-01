### Go Logging
a library for logging in Go using sirupsen-logrus

### How to use
- Add logrus dependency, type command on terminal : 
    ```
    go get github.com/sirupsen/logrus
    ```
- Example: using in main class, then run it.
    ```go
    package main
    
    import (
        logConfig "github.com/mfirmanakbar/go-logging/log"
        log "github.com/sirupsen/logrus"
    )
    
    func main() {
        logConfig.InitializeLogging("/tmp/test1.log")
        log.Println("Something Happened ...")
        log.Fatalf("What Happened ?")
    }
    ```
- Check the log on e.g Mac directory is `cat /tmp/test1.log`
    ```shell script
    firman@Cafe-Latte $ cat /tmp/test1.log
    {"level":"info","msg":"Something Happened ...","time":"2020-03-01T09:20:04+07:00"}
    {"level":"fatal","msg":"What Happened ?","time":"2020-03-01T09:20:04+07:00"}
    ```