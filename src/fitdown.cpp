#include <iostream>
#include <string>
#include <cstdlib>
#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm> 
using namespace std;

struct WorkoutLog{
    string userName;
    string exercise;
    int sets;
    vector<int> reps;
    vector<int> weight;
    string date;
};


struct UserData{
    string userName;
    float userWeight;
    string userHeight;
    int userAge;

    int maxBench;
    int maxSquat;
    int maxDeadlift;
};


class FitnessTracker{
    private: 
        vector<UserData> users;
        string fileName;

        void saveUsers() const{
            ofstream outFile(fileName);

            if(!outFile){
                cout << "Error opening file!" << endl;
                return;
            }

            for(const UserData& user : users){ //itterate thru vector of users and send to output file to save
                outFile << user.userName << " "
                        << user.userWeight << " "
                        << user.userHeight << " "
                        << user.userAge << " "
                        << user.maxBench << " "
                        << user.maxSquat << " "
                        << user.maxDeadlift << endl;       
            }
        }

        void loadUsers(){
            ifstream inFile(fileName);

            if(!inFile){
                cout << "Error opening file" << endl;
                return;
            }

            while(!inFile.eof()){
                UserData newUser;
                //grab info from saved file "load users"
                inFile >> newUser.userName >> newUser.userWeight >> newUser.userHeight >> newUser.userAge >> newUser.maxBench >> newUser.maxSquat >> newUser.maxDeadlift;

                if(!inFile.fail()){
                    users.push_back(newUser);
                }
            }
        }


        vector<UserData>::iterator findUserByName(const string& name){
            return find_if(users.begin(), users.end(), [&name](const UserData& user) {  //algorithm to find user by name
                return user.userName == name;
            });
        }


    public:
        //constructor to initialize filename
        FitnessTracker(const string& filename) : fileName(filename){
            loadUsers();
        }

        //destructor to save users to file when program ends
        ~FitnessTracker(){
            saveUsers();
        }

        void addOrUpdateUsersFromInput(){
            string name, height;
            float weight;
            int age, bench, deadlift, squat;

            cout << "\n-----------------" << endl;
            cout << "Enter user name: " << endl;
            cout << "-----------------" << endl;
            cin >> name;

            //check if user is already in file
            auto existingUser = findUserByName(name);
            if(existingUser != users.end()){ //found existing user
                cout << "\nUser found. Updating information." << endl;
                cout << "-----------------------------------" << endl;
                cout << "Enter new weight: ";
                cin >> weight;

                cout << "Enter new height: ";
                cin >> height;

                cout << "Enter new age: ";
                cin >> age;

                cout << "Enter NEW bench PR: ";
                cin >> bench;

                cout << "Enter NEW squat PR: ";
                cin >> squat;

                cout << "Enter NEW deadlift PR: ";
                cin >> deadlift;

                //congrats on progress
                if(existingUser->userWeight > weight){ //weight loss
                    cout << "\nCongrats on loosing " << existingUser->userWeight - weight << " pounds!" << endl;
                }

                if(existingUser->maxBench < bench){  //bench pr
                    cout << "\nCongrats on NEW bench pr, welcome to the " << bench << " club!" << endl;
                }

                if(existingUser->maxSquat < squat){  // squat pr
                    cout << "\nCongrats on NEW squat pr, welcome to the " << squat << " club!" << endl;
                }

                if(existingUser->maxDeadlift < deadlift){ //deadlift pr
                    cout << "\nCongrats on NEW deadlift pr, welcome to the " << deadlift << " club!" << endl;
                }


                // Update user information
                existingUser->userWeight = weight;
                existingUser->userHeight = height;
                existingUser->userAge = age;
                existingUser->maxBench = bench;
                existingUser->maxSquat = squat;
                existingUser->maxDeadlift = deadlift;

            }
            else{
                //user does not exist add a new user
                cout << "\nUser not found. Adding new user." << endl;
                cout << "----------------------------------" << endl;
                cout << "Enter weight: ";
                cin >> weight;

                cout << "Enter height: ";
                cin >> height;

                cout << "Enter age: ";
                cin >> age;

                cout << "Enter bench PR: ";
                cin >> bench;

                cout << "Enter squat PR: ";
                cin >> squat;

                cout << "Enter deadlift PR: ";
                cin >> deadlift;

                addUsers(name,weight,height,age,bench,squat,deadlift);
            }
        }
    
        void addUsers(const string& name, float weight, string height, int age, int bench, int squat, int deadlift){
            UserData newUser;
            newUser.userName = name;
            newUser.userWeight = weight;
            newUser.userHeight = height;
            newUser.userAge = age;
            newUser.maxBench = bench;
            newUser.maxSquat = squat;
            newUser.maxDeadlift = deadlift;

            users.push_back(newUser);
            saveUsers();
        }

        void removeUser(){
            string name;
            cout << "\nEnter name of user to be removed: " << endl;
            cin >> name;
            auto userToRemove = findUserByName(name);

            if(userToRemove != users.end()){
                users.erase(userToRemove);
                cout << "Sorry to see you leave, " << name << endl;
                saveUsers();
            }
            else{
                cout << "User " << name << " was not found, try another name" << endl;
            }
        }


        void printUserInfo(){
            string currentName;
            cout << "\nEnter your name to see your stats: ";
            cin >> currentName;
            bool userFound = false;

            for(UserData& user : users){
                if(user.userName == currentName){
                    cout << "\n---------------------" << endl;
                    cout << "Name: " << user.userName << endl;
                    cout << "Weight: " << user.userWeight << endl;
                    cout << "Height: " << user.userHeight << endl;
                    cout << "Age: " << user.userAge << endl;
                    cout << "-----------------" << endl;
                    cout << "Current bench PR: " << user.maxBench << endl;
                    cout << "Current squat PR: " << user.maxSquat << endl;
                    cout << "Current deadlift PR: " << user.maxDeadlift << endl;
                    cout << "---------------------" << endl;
                    userFound = true;
                    break;
                }
            }
            if(!userFound){
                cout << "User " << currentName << " was not found" << endl;
            }
        }
};


class LogWorkout{
    private:
        vector<WorkoutLog> logs;
        string logName;

        void saveLog() const{
            ofstream workoutFile(logName);
    
            if(!workoutFile){
                cout << "Error opening workout file" << endl;
                return;
            }
        
            for(const WorkoutLog& workout : logs){ //itterate thru vector of workouts and send to output file to save
                workoutFile << workout.userName << " "
                << workout.date << " "
                << workout.exercise << " "
                << workout.sets << " ";
                for(const auto &r : workout.reps){
                    workoutFile << r << " ";
                }
                for(const auto &w : workout.weight){
                    workoutFile << w << " ";
                }
                workoutFile << endl;
            }
        }

        void loadLog(){
            ifstream workinFile(logName);

            if(!workinFile){
                cout << "Error opening workout file" << endl;
                return;
            }

            while(!workinFile.eof()){
                WorkoutLog newWorkout;
                int repz;
                int weightz;
                //grab info from saved file "load workouts"
                workinFile >> newWorkout.userName >> newWorkout.date >> newWorkout.exercise >> newWorkout.sets;
                for(int i = 0; i < newWorkout.sets; i++){
                    workinFile >> repz;
                    newWorkout.reps.push_back(repz);
                }

                for(int i = 0; i < newWorkout.sets; i++){
                    workinFile >> weightz;
                    newWorkout.weight.push_back(weightz);
                }
                
                //  workinFile >> newWorkout.date;

                if(!workinFile.fail()){
                    logs.push_back(newWorkout);
                }
            }
        }

    public:
        //constructor to initialize filename
        LogWorkout(const string& logname) : logName(logname){
            loadLog();
        }

        //destructor to save workouts to file when program ends
        ~LogWorkout(){
            saveLog();
        }

        //addWorkout function
        void addWorkout(){
            string name, date, exercise;
            int sets, repz;
            vector<int> reps;
            vector<int> weights;
            int weightz;

            cout << "Enter name: " << endl;
            cin >> name;

            cout << "Enter today's date: " << endl;
            cin >> date;

            cout << "Enter exercise name: " << endl;
            cin >> exercise;

            cout << "Enter number of sets: " << endl;
            cin >> sets;

            for(int i = 1; i <= sets; i++){
                cout << "Enter reps completed for set " << i << endl;
                cin >> repz;
                reps.push_back(repz);

                cout << "Enter weight lifted for set " << i << endl;
                cin >> weightz;
                weights.push_back(weightz);
                
            }

            // cout << "Enter number of reps: " << endl;
            // cin >> reps;

            

            // cout << "Enter weight: " << endl;
            // cin >> weight;

            WorkoutLog newWorkout;
            newWorkout.userName = name;
            newWorkout.date = date;
            newWorkout.exercise = exercise;
            newWorkout.reps = reps;
            newWorkout.sets = sets;
            newWorkout.weight = weights;

            logs.push_back(newWorkout);
            saveLog();
        }

        //remove workout function

        //print function
        void printWorkout(){
            string userName;
            cout << "Enter username to see your lifts: " << endl;
            cin >> userName;
            cout << endl;

            bool userFound = false;

            for(WorkoutLog& workout : logs){
                if(workout.userName == userName){
                    // cout << "\n" << userName << "'s workouts: " << endl;
                    // cout << "---------------------" << endl;
                    cout << "\n" << workout.date << ": |" << workout.exercise <<"|" << endl;
                    for(int i = 0; i < workout.sets; i++){
                        cout << " set "<< i+1 << ": weight lifted: " << workout.weight[i] <<   " reps executed: " << workout.reps[i] << endl;
                        cout << "---------------------" << endl;
                    }

                    // cout << " weight: " << workout.weight << endl;
                    userFound = true;
                    // break;
                }
            }
            if(!userFound){
                cout << "User " << userName << " was not found" << endl;
            }
        }

        //pid is date 
        // if same date and same workout add to sets n reps accordingly
            

};


int main(){
    FitnessTracker fitnesstracker("user_data.txt");
    LogWorkout logworkout("workout_data.txt");
    int choice;

    do{ //loop until user no longer wants to update or check stats
        cout << "\n\n-----------------" << endl;
        cout << "      MENU" << endl;
        cout << "-----------------" << endl;
        cout << "(1) New user || Returning user" << endl;
        cout << "-----------------" << endl;
        cout << "(2) See your stats" << endl;
        cout << "-----------------" << endl;
        cout << "(3) Remove user" << endl;
        cout << "-----------------" << endl;
        cout << "(4) log workout" << endl;
        cout << "-----------------" << endl;
        cout << "(5) See your workouts" << endl;
        cout << "-----------------" << endl;
        cout << "(6) Exit" << endl;
        cout << "-----------------" << endl;

        cout << "\nENTER CHOICE: " << endl;
        cin >> choice;
        cout <<"-"<< endl;

        if(choice == 1){
            fitnesstracker.addOrUpdateUsersFromInput(); //checks whether name is in file if it is you can modify your stats  if not new user is created
        }
        if(choice == 2){
            fitnesstracker.printUserInfo();  //asks for name you want the info of and itterates vector for that names data.
        }
        if(choice == 3){
            fitnesstracker.removeUser();  //remove certain user from the list
        }
        if(choice == 4){
            logworkout.addWorkout();
        }
        if(choice == 5){
            logworkout.printWorkout();
        }
    } while(choice != 6);


    return 0;
}

//input name... if name already exists allows user to update stats
                //if new name then new user is allocated and stats are logged


//whenever returning user losses weight or PR'd a congratulations note will display


//enter name to see your stats

//remove user


//ONE PROBLEM THE CONSOLE KEEPS GETTING PUSHED DOWN CAUSE SO MUCH TEXT. MAYBE EVEN HARD TO READ AFTER A COUPLE OF CHOICES
