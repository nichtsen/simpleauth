package simpleauth

import (
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	cred string
	root string
}

func New(cred, root string) *Server {
	return &Server{
		cred: cred,
		root: root,
	}
}

func (s *Server) Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			cred  string
			auths []string
		)
		cred = r.Header.Get("Authorization")
		if cred == "" {
			goto unAuth
		}
		auths = strings.Split(cred, " ")
		fmt.Println(auths)
		if auths[0] == "Basic" && encodeString([]byte(s.cred)) == auths[1] {
			goto Auth
		} else {
			str := encodeString([]byte(s.cred))
			if len(str) < len(auths[1]) && str == auths[1][:len(str)] {
				goto Auth
			}
			fmt.Printf("%v: %v ", encodeString([]byte(s.cred)), auths[1])
		}

	unAuth:
		w.Header().Set("WWW-Authenticate", "Basic realm=\"test@google.com\"")
		w.WriteHeader(http.StatusUnauthorized)
		return
	Auth:
		http.ServeFile(w, r, s.root)
	})
	http.ListenAndServe(":30808", nil)
}
