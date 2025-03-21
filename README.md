# Pirana CLI
Pirana-cli is the open source friendly version of cli app from The Alfred Project, my personal use butler application for everyday cybersecurity use. This repository contains GO code for a cli application called ALFRED.

## Requirements
This repository contains only the CLI code. The CLI app makes request to a localhost backend which is connected to the database. The backend server must be available while running the cli app.

## Environment variables
The following contents are required in the environment variable file `.env`
```
CLI_PASS=
```

## Setting up the codebase
To set up the codebase for the first time, run the following command to install the required Go packages:
```sh
go mod tidy
```

## Building the application
Run the command
```sh
go build -o alfred
```

## Running the application
After building the application, run the executable file
```sh
./alfred
```