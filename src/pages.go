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
	res, err, _ := handleGetRequest(requestUrl)
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
		w.Write([]byte("Missing user val"))
		return
	}
	qFirstName := req.PostFormValue("firstname")
	if qFirstName == "" {
		w.Write([]byte("Missing first name val"))
		return
	}
	qLastName := req.PostFormValue("lastname")
	if qLastName == "" {
		w.Write([]byte("Missing last name val"))
		return
	}
	qHeight := req.PostFormValue("height")
	if qHeight == "" {
		w.Write([]byte("Missing height val"))
		return
	}
	qWeight := req.PostFormValue("weight")
	if qWeight == "" {
		w.Write([]byte("Missing weight val"))
		return
	}
	data := map[string]map[string]interface{}{}
	data[qUser] = map[string]interface{}{}
	data[qUser]["height"] = qHeight
	data[qUser]["weight"] = qWeight
	data[qUser]["firstName"] = qFirstName
	data[qUser]["lastName"] = qLastName

	serverPort := 3333
	requestUrl := fmt.Sprintf("http://localhost:%d/user", serverPort)
	res, err, resCode := handlePostRequest(requestUrl, data)
	if err != nil {
		errCode := fmt.Sprintf("Could not get user : %s\n", err)
		w.Write([]byte(errCode))
		return
	} else if resCode != 200 {
		w.Write([]byte(res))
		return
	}
	w.Write([]byte("Sucessful Post!"))
}

func functionalUser(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Fn functionalUser called\n")
	hasUser := req.URL.Query().Has("user")
	if !hasUser {
		w.Header().Set("x-missing-field", "user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	qUser := req.URL.Query().Get("user")
	//Should also be grabbing user info, but for now just request the workouts
	qUserStruct := []FitnessUser{
		{
			ID:        qUser,
			FirstName: "Nathan",
			LastName:  "Kim",
			Height:    72,
			FtOrCm:    0,
			Weight:    190,
			LbOrKg:    0,
		},
	}

	var tmplFile = "user.html"
	tmpl := template.Must(template.ParseFiles(tmplFile))

	err := tmpl.Execute(w, qUserStruct)
	if err != nil {
		panic(err)
	}
}
