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
	flag.Parse()
	fmt.Println(root, user, pwd)
	server := simpleauth.New(user+":"+pwd, root)
	server.Serve()
}
