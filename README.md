# Gin-Chat-Backend
A chat application backend built with Gin framework in Go, interfacing with Google Firestore.

## Setup

Ensure you have Go installed on your system.

### Install Dependencies

From the project's root directory, install the necessary dependencies:

```bash
cd backend
go mod tidy
```

### Configure Environment Variables
Store your Firestore credentials JSON file in the main project directory and set the environment variable to its path:

```bash
export GOOGLE_APPLICATION_CREDENTIALS='/absolute/path/to/firestore_key.json'
```

Replace '/absolute/path/to/firestore_key.json' with the actual path to your Firestore credentials file (received via email)

Update the Firestore service configuration in 'services/firestoreService.go' to reflect the path to your credentials file.

## Run the application

To start the server, run the following command from the backend directory:

```bash
go run main.go
```