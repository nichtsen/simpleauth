package simpleauth

import (
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
		// r.BasicAuth()
		cred = r.Header.Get("Authorization")
		if cred == "" {
			goto unAuth
		}
		auths = strings.Split(cred, " ")
		if auths[0] == "Basic" && encodeString([]byte(s.cred)) == auths[1] {
			goto Auth
		}
	unAuth:
		w.Header().Set("WWW-Authenticate", "Basic realm=\"test@google.com\"")
		w.WriteHeader(http.StatusUnauthorized)
		return
	Auth:
		handler := http.FileServer(http.Dir(s.root))
		handler.ServeHTTP(w, r)
	})
	http.ListenAndServe(":30808", nil)
}
