package main

import (
	"plasticine/db"
	"plasticine/server"
)

func main() {
	server.NewServer(":80", db.NewDefaultDb()).Run()
}
