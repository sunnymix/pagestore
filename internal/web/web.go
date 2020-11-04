package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pagestore/internal/web/controller"
)

type ApiDef struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

var ApiMap map[string]string

func Start(port int) {
	http.HandleFunc("/api/page", controller.HandlePage)
	http.HandleFunc("/api/pages", controller.HandlePages)
	http.HandleFunc("/", ApiMapDefault)

	fmt.Printf("Start web server at port: %d\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func ApiMapDefault(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(ApiMap)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}

func init() {
	ApiMap = make(map[string]string)
	ApiMap["page_url"] = "/api/page"
	ApiMap["pages_url"] = "/api/pages"
}
