package main

import (
	"log"
	"projection/cmd"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err :%v", err)
	}
}
