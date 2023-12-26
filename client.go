package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	//"time"
	"io/ioutil"
	"os"
)

func main() {
	// c := http.Client{Timeout: time.Duration(1) * time.Second}
	// resp, err := c.Get("http://localhost:8080/fitness")
	// if err!= nil{
	// 	fmt.Printf("Error %s", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("Body: %s", body)
	serverPort := 3333
	requestUrl := fmt.Sprintf("http://localhost:%d/user?user=nkim256", serverPort)
	data := map[string]map[string]interface{}{}
	data["nkim256"] = map[string]interface{}{}
	data["nkim256"]["name"] = "Nathan Kim"
	data["nkim256"]["height"] = "172 in"
	data["nickdraggy"] = map[string]interface{}{}
	data["nickdraggy"]["name"] = "Nicholas Kim"
	data["nickdraggy"]["height"] = "170 in"
	handleRequest(requestUrl, data, http.MethodPost)
	handleRequest(requestUrl, data, http.MethodGet)

}

func handleRequest(requestUrl string, data map[string]map[string]interface{}, method string) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	jsonReader := bytes.NewReader(jsonData)
	req, err := http.NewRequest(method, requestUrl, jsonReader)
	if err != nil {
		fmt.Printf("Error creating this httpRequest: %s\n", err)
		os.Exit(1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Client got a response.\n")
	fmt.Printf("client response code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client response body: %s\n", resBody)
}
