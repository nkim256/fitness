# fitness
fitness app
General Requirements
- User has a profile 
	- name
	- height
	- Weight
- Each user can log a workout
- Workout
	- Consists of a series of exercises
	- exercises consist of sets
	- sets consist of reps
	- This can be used to calculate total volume of weight lifted 
- User can create friend groups
	- Friends in groups are ranked by total volume by week, month, year
	- MAYBE: Chat per friends group

Make sure you have go and mysql installed.

Start mysql cli, and run the create_database.sql command.
Then head into src folder and run go build -> go run .
http://localhost:4444/user?user=nkim256
To try it out.

 For Now to use run git clone git@github.com:nkim256/fitness.git
