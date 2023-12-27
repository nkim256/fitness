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
	
# Welcome to your CDK TypeScript project

This is a blank project for CDK development with TypeScript.

The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run test`    perform the jest unit tests
* `npx cdk deploy`  deploy this stack to your default AWS account/region
* `npx cdk diff`    compare deployed stack with current state
* `npx cdk synth`   emits the synthesized CloudFormation template
