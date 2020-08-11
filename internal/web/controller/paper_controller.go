package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"paperstore/internal/repo"
)

func HandlePaper(w http.ResponseWriter, r *http.Request) {
	var (
		res *JsonResult
	)

	if r.Method == "GET" {
		res = findOnePaper(getQuery(r, "pid"))
	} else if r.Method == "POST" {
		res = saveOnePaper(getBody(r))
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

func getBody(r *http.Request) *repo.Paper {
	var res repo.Paper
	_ = json.NewDecoder(r.Body).Decode(&res)
	return &res
}

func findOnePaper(pid string) *JsonResult {
	res := &JsonResult{}

	if len(pid) == 0 {
		return res
	}

	paper, _ := repo.GlobalRepo.FindOne(pid)
	res.Data = paper

	return res
}

func saveOnePaper(paper *repo.Paper) *JsonResult {
	res := &JsonResult{}

	err := repo.GlobalRepo.SaveOne(paper)
	if err != nil {
		res.Code = 1
		res.Msg = fmt.Sprintf("save paper error: %s", err)
	}

	return res
}

func HandlePapers(w http.ResponseWriter, r *http.Request) {
	var (
		res *JsonResult
	)

	if r.Method == "GET" {
		res = findPagePapers()
	}

	js, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func findPagePapers() *JsonResult {
	res := &JsonResult{}

	papers, err := repo.GlobalRepo.FindPage()
	if err != nil {
		res.Code = 1
		res.Msg = fmt.Sprintf("find page papers error: %s", err)
	} else {
		res.Data = papers
	}

	return res
}
