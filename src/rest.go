package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func user(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("server: %s\n", req.Method)
	//fmt.Printf("server query id: %s\n", req.URL.Query().Get("id"))
	fmt.Printf("server content type: %s\n", req.Header.Get("content-type"))
	fmt.Printf("server: headers:\n")
	for headerName, headerValue := range req.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}
	if req.Method == http.MethodGet {
		getUser(w, req)
	}

	if req.Method == http.MethodPost {
		postUser(w, req)
	}

}

func getUser(w http.ResponseWriter, req *http.Request) {
	hasUser := req.URL.Query().Has("user")
	if !hasUser {
		w.Header().Set("x-missing-field", "user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileData, err := ioutil.ReadFile("userData.json")
	if err != nil {
		fmt.Printf("Couldn't open file: %s\n", err)
		return
	}
	userData := map[string]map[string]interface{}{}
	err = json.Unmarshal(fileData, &userData)
	if err != nil {
		fmt.Printf("Couldn't Unmarshal: %s\n", err)
		return
	}
	user := req.URL.Query().Get("user")
	_, ok := userData[user]
	if !ok {
		fmt.Printf("User does not exist... terminate\n")
		fmt.Fprintf(w, "User "+user+" does not exist\n")
		return
	} else {
		fmt.Printf("User info: %s\n", user)
		jsonData, err := json.Marshal(userData[user])
		if err != nil {
			fmt.Printf("Could not marshal json: %s\n", err)
			return
		}
		fmt.Fprintf(w, string(jsonData))
	}
}

func postUser(w http.ResponseWriter, req *http.Request) {
	data := map[string]map[string]interface{}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		fmt.Printf("Could not decode: %s\n", err)
		return
	}
	fileData, err := ioutil.ReadFile("userData.json")
	if err != nil {
		fmt.Printf("Couldn't open file: %s\n", err)
		return
	}
	userData := map[string]map[string]interface{}{}
	if len(fileData) != 0 {
		err = json.Unmarshal(fileData, &userData)
		if err != nil {
			fmt.Printf("Couldn't Unmarshal: %s\n", err)
			return
		}
	}

	for key, value := range data {
		userData[key] = value
	}
	outgoingJson, err := json.Marshal(userData)
	if err != nil {
		fmt.Printf("Could not marshal json: %s\n", err)
		return
	}

	err = ioutil.WriteFile("userData.json", outgoingJson, 0644)
	if err != nil {
		fmt.Printf("Could not write to file: %s\n", err)
		return
	}
}
