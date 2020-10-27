package main

import (
	"github.com/koko2824/go_crud/db"
	"github.com/koko2824/go_crud/server"
)

func main() {
	db.Init()
	server.Init()
}