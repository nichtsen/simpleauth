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
