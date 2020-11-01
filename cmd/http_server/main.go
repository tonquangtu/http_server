package main

import (
	"fmt"
	"github.com/tonquangtu/http_server/internal/router"
	"net/http"
)

func main() {
	r := router.NewRouter()
	if err := http.ListenAndServe(":9000", r); err != nil {
		fmt.Println("Cannot start server at port 9000")
		panic(err)
	}
}
