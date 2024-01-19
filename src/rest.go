package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
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
		w.Header().Set("user-get-error", "missing username")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := req.URL.Query().Get("user")
	var usr fitnessUser
	inst := db.QueryRow("SELECT * FROM users WHERE id = ?", user)
	err := inst.Scan(&usr.ID, &usr.FirstName, &usr.LastName, &usr.Height,
		&usr.FtOrCm, &usr.Weight, &usr.LbOrKg)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("User does not exist: %s\n", user)
			w.Header().Set("user-get-error", "user does not exist")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Printf("User: %s, Error: %s\n", user, err)
		w.Header().Set("user-get-error", "error retrieving")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userData := map[string]map[string]interface{}{}
	userData[usr.ID] = map[string]interface{}{}
	userData[usr.ID]["firstName"] = usr.FirstName
	userData[usr.ID]["lastName"] = usr.LastName
	if usr.FtOrCm == 0 {
		userData[usr.ID]["height"] = strconv.Itoa(usr.Height) + " in"
	} else {
		userData[usr.ID]["height"] = strconv.Itoa(int(float64(usr.Height)*2.54)) + " cm"
	}
	if usr.LbOrKg == 0 {
		userData[usr.ID]["weight"] = strconv.Itoa(usr.Weight) + " lbs"
	} else {
		userData[usr.ID]["weight"] = strconv.Itoa(int(float64(usr.Weight)*0.45359237)) + " kg"
	}
	fmt.Printf("User info: %s\n", user)
	jsonData, err := json.Marshal(userData[user])
	if err != nil {
		fmt.Printf("Could not marshal json: %s\n", err)
		w.Header().Set("user-get-error", "error marshaling")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(jsonData))
}

func postUser(w http.ResponseWriter, req *http.Request) {
	data := map[string]map[string]interface{}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		fmt.Printf("Could not decode: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for key, _ := range data {
		var exists bool
		row := db.QueryRow("select exists(select * from users where id=?)", key)
		if err = row.Scan(&exists); err != nil {
			fmt.Printf("Error scanning row: %s\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if exists {
			fmt.Printf("User already exists: %v\n", exists)
			w.Header().Set("user-post-error", "user exists")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//Does not exist, go ahead and post user
		_, err = db.Exec("insert into users (id, first_name, last_name, height, feetOrInches, userWeight, lbOrKg)values (?, ?, ?, ?, 0, ?, 0)",
			key, data[key], data[key]["firstName"], data[key]["lastName"],
			data[key]["height"], data[key]["weight"])

		if err != nil {
			fmt.Printf("Error inserting user: %s\n", err)
			w.Header().Set("user-post-error", "error during insertion")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func getWorkout(w http.ResponseWriter, req *http.Request) {
	data 
}

func recordWorkout(userId string) bool, int{
	if(userId==nil)
	{
		fmt.Printf("User does not exist")
		return false, 0
	}

	




}
func recordExercise(workoutId int, workoutType string) bool{
	if(workoutId==nil || workoutType==nil){
		fmt.Printf("Bad Inputs")
		return false
	}
	_, err := db.Exec("insert into exercises (id, workout_id, workout_type) values (0, ?,?)",
		workoutId, workoutType)

	if err!=nil{
		fmt.Printf("Error inserting exercise: %s", err)
		return false
	}

	fmt.Printf("Successful insert")
	return true


}

func recordExcerciseSet(exerciseId int, weight int, reps int) bool{
	if(exerciseId==nil){
		return false
	}
	if(weight==nil || reps==nil || weight <=0 || reps <0){
		return false
	}
	_, err := db.Exec("insert into exercise_sets (id, exercise_id, weight_amt, reps) values (0, ?, ?, ?)",
		exerciseId, weight, reps)
	if err!=nil{
		fmt.Printf("Error inserting exercise setL %s\n", err)
		return false
	}

	fmt.Printf("Successful insert")
	return true
}
