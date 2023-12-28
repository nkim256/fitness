package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	restMux := http.NewServeMux()
	restMux.HandleFunc("/user", user)
	pageMux := http.NewServeMux()
	pageMux.Handle("/static/", http.StripPrefix("/static/", fs))
	pageMux.HandleFunc("/", indexHandler)
	pageMux.HandleFunc("/searchUser", searchUser)
	pageMux.HandleFunc("/about", about)
	pageMux.HandleFunc("/makeUser", makeUser)
	ctx, cancelCtx := context.WithCancel(context.Background())
	restServer := &http.Server{
		Addr:    ":3333",
		Handler: restMux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "serverAddr", l.Addr().String())
			return ctx
		},
	}
	pageServer := &http.Server{
		Addr:    ":4444",
		Handler: pageMux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "serverAddr", l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := restServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed \n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func() {
		err := pageServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed \n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()
	<-ctx.Done()
	serverPort := 3333
	requestUrl := fmt.Sprintf("http://localhost:%d/user?user=nkim256", serverPort)
	data := map[string]map[string]interface{}{}
	data["nkim256"] = map[string]interface{}{}
	data["nkim256"]["name"] = "Nathan Kim"
	data["nkim256"]["height"] = "172 in"
	data["nickdraggy"] = map[string]interface{}{}
	data["nickdraggy"]["name"] = "Nicholas Kim"
	data["nickdraggy"]["height"] = "170 in"
	handlePostRequest(requestUrl, data)
	handleGetRequest(requestUrl)
}

func handlePostRequest(requestUrl string, data map[string]map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return err
	}
	jsonReader := bytes.NewReader(jsonData)
	req, err := http.NewRequest(http.MethodPost, requestUrl, jsonReader)
	if err != nil {
		fmt.Printf("Error creating this httpRequest: %s\n", err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return err
	}
	fmt.Printf("Client got a response.\n")
	fmt.Printf("client response code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		return err
	}
	fmt.Printf("client response body: %s\n", resBody)
	return err
}

func handleGetRequest(requestUrl string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		fmt.Printf("Error creating this httpRequest: %s\n", err)
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return nil, err
	}
	fmt.Printf("Client got a response.\n")
	fmt.Printf("client response code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		return nil, err
	}
	fmt.Printf("client response body: %s\n", resBody)
	return resBody, err
}
