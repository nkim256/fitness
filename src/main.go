package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type fitnessUser struct {
	ID        string
	FirstName string
	LastName  string
	Height    int
	FtOrCm    int
	Weight    int
	LbOrKg    int
}

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "yesyesyes",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "fitness",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Printf("Failure connecting to database: %s", err)
		return
	}

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
}

func handlePostRequest(requestUrl string, data map[string]map[string]interface{}) ([]byte, error, int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return []byte("Could not marshal json"), err, 0
	}
	jsonReader := bytes.NewReader(jsonData)
	req, err := http.NewRequest(http.MethodPost, requestUrl, jsonReader)
	if err != nil {
		fmt.Printf("Error creating this httpRequest: %s\n", err)
		return []byte("Eror creating httpRequest"), err, 0
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return []byte(req.Header.Get("user-post-error")), err, res.StatusCode
	}
	fmt.Printf("Client got a response.\n")
	fmt.Printf("client response code: %d\n", res.StatusCode)

	return []byte(req.Header.Get("user-post-error")), err, res.StatusCode
}

func handleGetRequest(requestUrl string) ([]byte, error, int) {
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		fmt.Printf("Error creating this httpRequest: %s\n", err)
		return nil, err, 0
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return nil, err, 0
	}
	fmt.Printf("Client got a response.\n")
	fmt.Printf("client response code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		return nil, err, res.StatusCode
	} else if res.StatusCode != 200 {
		return []byte(res.Header.Get("user-get-error")), err, res.StatusCode
	}
	fmt.Printf("client response body: %s\n", resBody)
	return resBody, err, res.StatusCode
}
