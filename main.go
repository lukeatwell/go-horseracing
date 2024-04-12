package main

import "errors"
import 	"database/sql"
import "os"
import "fmt"
import "github.com/lukeatwell/go-horseracing/database"

import _ "github.com/mattn/go-sqlite3"

func main() {
	_, err := os.Stat("./file.db")
	if errors.Is(err, os.ErrNotExist) {
		os.Create("./file.db")
	} else {
		db, err := sql.Open("sqlite3", "./file.db")
		if err != nil {
			fmt.Println(err)
		}

		defer db.Close()
	}	
}