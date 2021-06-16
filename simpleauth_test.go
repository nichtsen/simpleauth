package simpleauth

import "testing"

func TestServe(t *testing.T) {
	s := &Server{
		cred: "a:bb",
		root: "index.html",
	}
	s.Serve()
}
