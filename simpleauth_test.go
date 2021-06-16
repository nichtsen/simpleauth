package simpleauth

import "testing"

func TestServe(t *testing.T) {
	s := &Server{
		cred: "a:b",
	}
	s.Serve()
}
