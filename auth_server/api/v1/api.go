package v1

import (
	"fmt"
	"net/http"
)

func DoAPI(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.Path == "/api/v1/":
		doIndex(rw, req)
	case req.URL.Path == "/api/v1/acls":
		doACLs(rw, req)
	case req.URL.Path == "/api/v1/users":
		doUsers(rw, req)
	default:
		http.Error(rw, "Not found", http.StatusNotFound)
		return
	}
}

func doIndex(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text-html; charset=utf-8")
	fmt.Fprintf(rw, "<h1>API V1</h1>\n")
}

func doACLs(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text-html; charset=utf-8")
	fmt.Fprintf(rw, "<h1>API V1 ACLs</h1>\n")
}

func doUsers(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text-html; charset=utf-8")
	fmt.Fprintf(rw, "<h1>API V1 Users</h1>\n")
}
