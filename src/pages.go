package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, req *http.Request) {
	tpl.Execute(w, nil)
}

func about(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Nicholas and Nathan fitness page!</h1>"))
}

func searchUser(w http.ResponseWriter, req *http.Request) {
	hasUser := req.URL.Query().Has("user")
	if !hasUser {
		w.Header().Set("x-missing-field", "user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	qUser := req.URL.Query().Get("user")
	serverPort := 3333
	requestUrl := fmt.Sprintf("http://localhost:%d/user?user=%s", serverPort, qUser)
	res, err := handleGetRequest(requestUrl)
	if err != nil {
		errCode := fmt.Sprintf("Could not get user : %s\n", err)
		w.Write([]byte(errCode))
		return
	}
	w.Write(res)
}

func makeUser(w http.ResponseWriter, req *http.Request) {
	qUser := req.PostFormValue("user")
	if qUser == "" {
		w.Header().Set("x-missing-field", "user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	qName := req.PostFormValue("name")
	if qName == "" {
		w.Header().Set("x-missing-field", "name")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	qHeight := req.PostFormValue("height")
	if qHeight == "" {
		w.Header().Set("x-missing-field", "height")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := map[string]map[string]interface{}{}
	data[qUser] = map[string]interface{}{}
	data[qUser]["name"] = qName
	data[qUser]["height"] = qHeight
	serverPort := 3333
	requestUrl := fmt.Sprintf("http://localhost:%d/user", serverPort)
	err := handlePostRequest(requestUrl, data)
	if err != nil {
		errCode := fmt.Sprintf("Could not get user : %s\n", err)
		w.Write([]byte(errCode))
		return
	}
	w.Write([]byte("Sucessful Post!"))

}
