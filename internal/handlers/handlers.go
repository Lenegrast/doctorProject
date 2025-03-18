package handlers

import (
	"net/http"
)

func Example(rw http.ResponseWriter, res *http.Request) {
	rw.Write([]byte("asd"))
}
