package main

import (
	"flag"
	"fmt"

	"github.com/nichtsen/simpleauth"
)

func main() {
	var (
		root string
		user string
		pwd  string
	)
	flag.StringVar(&root, "i", "index.html", "path to index.html to serve")
	flag.StringVar(&user, "u", "user", "user name")
	flag.StringVar(&pwd, "p", "pwd@2333", "password")
	fmt.Println(root, user, pwd)
	server := &simpleauth.Server{
		cred: user + ":" + "pwd",
		root: root,
	}
	server.Serve()
}
