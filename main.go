package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/fitness", fitness)
	mux.HandleFunc("/user", user)
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "serverAddr", l.Addr().String())
			return ctx
		},
	}
	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "serverAddr", l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed \n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed \n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()
	<-ctx.Done()
}

func getRoot(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Printf("%s: got / request\n", ctx.Value("serverAddr"))
	io.WriteString(w, "Welcome to Root")
}

func fitness(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Printf("%s: got / fitness request\n", ctx.Value("serverAddr"))
	io.WriteString(w, "Welcome to fitness app!")
}

func user(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("server: %s\n", req.Method)
	fmt.Printf("server query id: %s\n", req.URL.Query().Get("id"))
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
		fmt.Printf("couldn't open file: %s\n", err)
		return
	}
	userData := map[string]map[string]interface{}{}
	err = json.Unmarshal(fileData, &userData)
	if err != nil {
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
		fmt.Printf("couldn't open file: %s\n", err)
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
		return
	}
}
