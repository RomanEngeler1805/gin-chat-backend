# gin-chat-backend
## Setup

### Install the requirements
Navigate into the backend folder
```
go build
```

### Set environment variables
Save the via email received firestore credentials in the main project folder
Set the environment variable
```
export GOOGLE_APPLICATION_CREDENTIALS='PATH_TO_CREDENTIALS.json'
```
Change the path to the firestore credentials in the servies/firestoreService.go file

## Run the application
Navigate into the /backend folder and run
```
go run main.go
```
