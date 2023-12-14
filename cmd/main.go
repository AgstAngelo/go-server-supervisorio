package main

import (
	"fmt"
	"net/http"
	"github.com/AgstAngelo/go-server-supervisorio/internal/router"
)

func main() {
	r := router.NewRouter()


	http.ListenAndServe(":80", r)
}