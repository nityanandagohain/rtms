package main

import (
	"net/http"

	"github.com/nityanandagohain/rtms/pkg/web"
)

func main() {
	h := web.NewHandler("localhost:6379", "")
	http.ListenAndServe(":3000", h)
}
