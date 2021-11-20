package main

import (
	"log"
	view "money-manager/router"
)

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func main() {
	view.StartService()
}
