package main

import (
  "database/sql"
  //"fmt"
  "log"
  _ "github.com/mattn/go-sqlite3"
  "sync"
  "github.com/amchang/golang-cli/web"
)

type WebCrawler struct {
  foo int
  mu sync.Mutex
}

func database() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table if not exists foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
  log.Println("New Database")
}

func main() {
  /*
  crawl_targets := WebCrawler {}
  c := make(chan int)
  go func(chan int) {
    fmt.Println("Hello world")
    database()
    c <- 1
  }(c)
  <- c*/
  web.Server()
}
