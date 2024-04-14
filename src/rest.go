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
	var usr FitnessUser
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

func getWorkoutsMock(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Called")
	data := map[string]map[string]interface{}{}
	data["workout_1"] = map[string]interface{}{}
	data["workout_1"]["user_id"] = "nkim256"
	data["workout_1"]["workout_date"] = "1-1-24"
	data["workout_2"] = map[string]interface{}{}
	data["workout_2"]["user_id"] = "nkim256"
	data["workout_2"]["workout_date"] = "1-2-24"
	jsonData, err := json.Marshal(data)
	if err!=nil{
		fmt.Printf("Error making JSon: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(jsonData))
	
}

func getUserWorkouts(w http.ResponseWriter, req *http.Request){
	fmt.Printf("Called getUserWorkouts")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	hasUser := req.URL.Query().Has("user")
	if !hasUser {
		w.Header().Set("x-missing-field", "user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hasQuery := req.URL.Query().Has("numQuery")
	if !hasQuery {
		w.Header().Set("x-missing-field", "user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := req.URL.Query().Get("user")
	numQuery, err := strconv.Atoi(req.URL.Query().Get("numQuery"))
	if err!=nil{
		fmt.Printf("numQuery was not an Int")
		return
	}

	rows, err := db.Query("select * from workouts where user_id=? order by workout_date desc", user)
	var workoutLog []Workout	
	for rows.Next() {
		if numQuery ==0 {
			break
		}
		var workoutInst Workout
		err = rows.Scan(&workoutInst.ID, &workoutInst.UserID, &workoutInst.WorkoutDate)
		if err!=nil{
			fmt.Printf("Error retrieving row: %s" , err)
				return
		}
		numQuery -=1
		workoutLog = append(workoutLog, workoutInst)	
	}

	marshalData, err := json.Marshal(workoutLog)
	fmt.Printf(string(marshalData))
	fmt.Fprintf(w, string(marshalData))

	
}

func recordWorkout(w http.ResponseWriter, req *http.Request){
	fmt.Printf("called recordWorkout\n")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	hasUser := req.URL.Query().Has("user")
	if !hasUser{
		fmt.Printf("No user found")
		w.Header().Set("x-missing-field", "user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := req.URL.Query().Get("user")
	recordWorkoutUser(user)

}
func recordWorkoutUser(userId string) (bool, int64){
	fmt.Printf("Calling record workout for: %s\n", userId)
	if(userId==""){
		fmt.Printf("userId was a null value\n")
		return false, -1		
	}
	currTime := time.Now()
    workout_date := fmt.Sprintf("%d-%d-%d %d:%d:%d",
								currTime.Year(),
								currTime.Month(),
								currTime.Day(),
								currTime.Hour(),
								currTime.Minute(),
								currTime.Second())
	res , err := db.Exec("insert into workouts (id, user_id, workout_date) values (0, ?, ?)",
		userId, workout_date)
	if err!= nil{
		fmt.Printf("Error inserting workout: %s\n", err)
		return false, -1
	}
	id, err := res.LastInsertId()
	if err!=nil{
		fmt.Printf("Error retrieving ID: %s\n", err)
		return false, -1
	}
	fmt.Printf("Workout ID: %d\n", id)
	fmt.Printf("Successful insert of workout\n")
	return true, id
}
func recordExercise(workoutId int, workoutType string) bool{
	if(workoutType==""){
		fmt.Printf("Bad Inputs\n")
		return false
	}
	_, err := db.Exec("insert into exercises (id, workout_id, workout_type) values (0, ?,?)",
		workoutId, workoutType)

	if err!=nil{
		fmt.Printf("Error inserting exercise: %s\n", err)
		return false
	}

	fmt.Printf("Successful insert\n")
	return true


}

func recordExcerciseSet(exerciseId int, weight int, reps int) bool{
	if(weight <=0 || reps <0){
		return false
	}
	_, err := db.Exec("insert into exercise_sets (id, exercise_id, weight_amt, reps) values (0, ?, ?, ?)",
		exerciseId, weight, reps)
	if err!=nil{
		fmt.Printf("Error inserting exercise setL %s\n", err)
		return false
	}

	fmt.Printf("Successful insert\n")
	return true
}
