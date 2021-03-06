package main

import (
	"fmt"
	"io"
	"log"
	"money-manager/controller"
	"money-manager/util"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

type PlainFormatter struct {
    TimestampFormat string
    LevelDesc []string
}

func init() { 
	log.SetFlags(log.Lshortfile | log.LstdFlags) 
	logrusInit()
}

func main() {
	runSqlMigrationUP()
	route := gin.Default()
	// gin.SetMode(gin.ReleaseMode)  					//set for release mode
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	controller.Route(route)
}

func runSqlMigrationUP() {
	migrations := &migrate.FileMigrationSource{
		Dir: "sqlMigration/",
	}

	db, err := util.ConnectMySQL()
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		logrus.Error("Error occcured:", err)
		return
	}

	logrus.Info(fmt.Sprintf("Applied %d migrations!\n", n))
}

func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
    timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	level := f.LevelDesc[entry.Level]

	if level != "INFO" {
		return []byte(fmt.Sprintf("[PROJECT-debug] %s [%s] %s~%v \t %s\n", timestamp, level, entry.Caller.File, 
		entry.Caller.Line, entry.Message)), nil
	} else {
		return []byte(fmt.Sprintf("[PROJECT-debug] %s [%s] %s\n", timestamp, level, entry.Message)), nil
	}
}

// logger format
func logrusInit() {
	plainFormatter := new(PlainFormatter)
    plainFormatter.TimestampFormat = "2006-01-02 15:04:05"
    plainFormatter.LevelDesc = []string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG"}

	path := "./log"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	logFile := path + "/logger.log"
	f, err := os.OpenFile(logFile, os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		logrus.Error(err.Error())
	}
	mw := io.MultiWriter(os.Stdout, f)
	logrus.SetReportCaller(true)
	logrus.SetOutput(mw)
	logrus.SetFormatter(plainFormatter)
}