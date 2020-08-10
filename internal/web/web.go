package web

import (
	"fmt"
	"net/http"
	"paperstore/internal/web/controller"
)

func Start(port int) {
	http.HandleFunc("/api/paper", controller.HandlePaper)
	http.HandleFunc("/api/papers", controller.HandlePapers)

	fmt.Printf("start web server at port: %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
