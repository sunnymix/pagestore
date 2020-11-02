package web

import (
	"fmt"
	"net/http"
	"pagestore/internal/web/controller"
)

func Start(port int) {
	http.HandleFunc("/api/page", controller.HandlePage)
	http.HandleFunc("/api/pages", controller.HandlePapers)

	fmt.Printf("start web server at port: %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
