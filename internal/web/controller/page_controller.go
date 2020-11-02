package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"pagestore/internal/repo"
)

func HandlePage(w http.ResponseWriter, r *http.Request) {
	var (
		res *JsonResult
	)

	if r.Method == "GET" {
		res = findOnePage(getQuery(r, "pid"))
	} else if r.Method == "POST" {
		res = saveOnePage(getBody(r))
	} else if r.Method == "DELETE" {

	}

	js, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func getQuery(r *http.Request, key string) (val string) {
	queries, _ := url.ParseQuery(r.URL.RawQuery)
	res := queries[key]
	if len(res) > 0 {
		return res[0]
	}
	return
}

func getBody(r *http.Request) *repo.Page {
	var res repo.Page
	_ = json.NewDecoder(r.Body).Decode(&res)
	return &res
}

func findOnePage(pid string) *JsonResult {
	res := &JsonResult{}

	if len(pid) == 0 {
		return res
	}

	paper, _ := repo.GlobalRepo.FindOne(pid)
	res.Data = paper

	return res
}

func saveOnePage(page *repo.Page) *JsonResult {
	res := &JsonResult{}

	err := repo.GlobalRepo.SaveOne(page)
	if err != nil {
		res.Code = 1
		res.Msg = fmt.Sprintf("save page error: %s", err)
	}

	return res
}

func HandlePapers(w http.ResponseWriter, r *http.Request) {
	var (
		res *JsonResult
	)

	if r.Method == "GET" {
		query := getQuery(r, "query")

		res = queryPages(query)
	}

	js, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func queryPages(query string) *JsonResult {
	res := &JsonResult{}

	pages, err := repo.GlobalRepo.Query(query)
	if err != nil {
		res.Code = 1
		res.Msg = fmt.Sprintf("Query pages error: %s", err)
	} else {
		res.Data = pages
	}

	return res
}
