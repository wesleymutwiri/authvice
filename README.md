# Authvice

A very simple authentication system in golang.

## Installation

1. Use dockerfile provided by running the following commands
    
    i. Clone Repository and change directory into the folder
    ```bash
    $ git clone https://github.com/wesleymutwiri/authvice.git && cd authvice
    ```
    ii. Change database settings in main.go to point to your own running instance of postgres
    iii. Build the file using docker creating a tag go-docker-optimized
    ```bash
    $ docker build -t go-docker-optimized -f Dockerfile .
    ```
    iv. Run the docker container once finished exposing the port to your local computer's port
    ```bash
    $ docker run -p 127.0.0.1:10000:10000 go-docker-optimized:latest
    ```
    v. The application will be running on port 10000 and can be accessed through: [localhost:10000](http://localhost:10000)

2. Run the code directly if you have go installed on the machine

    i. Clone Repository and change directory into the folder
    ```bash
    $ git clone https://github.com/wesleymutwiri/authvice.git && cd authvice
    ```
    ii. Build the application into a single binary via:
    ```bash
    $ go build -o main .
    ```
    iii. Run the executable go file by:
    ```bash
    $ ./main
    ```

3. With docker-compose (easiest, most recommended)
    i. Clone Repository and change directory into the folder
    ```bash
    $ git clone https://github.com/wesleymutwiri/authvice.git && cd authvice
    ```
    ii. Create a .env file with credentials for the database docker container
    ```bash
    $ echo "POSTGRES_USER=user \nPOSTGRES_PASSWORD=password \nPOSTGRES_DB=authvice" > .env
    ```
    iii. Ensure docker-compose is installed on your machine and run the following command:
    ```bash
    $ docker-compose up --build
    ```
    iv. Access the application on your local computer's port 10000

## Running tests
You can run the tests by simply running the following command in the root folder of the application. PS - If you have zsh installed refuse to accept autocorrect:
```bash 
go test ./... -v 
```

## Features

[] Create user
[] Login user and receive jwt token
[] Add User profile information
[] Edit User profile details
[] Login User with Github
[] Login User with Google
[] Login User with Facebook
[] Login User with Twitter
[] Add twilio/Africa's Talking for sending OTP via phone
[] Add two factor authentication using QR code
[] Add two factor authentication using SMS and calls
[] Create groups
[] Add users to certain groups
[] Change group permissions
[] Remove groups 
[] Edit groups
[] Enforce permissions for users
[] Ability to add or remove various functionality as one wishes.

## Use

Simply adding authentication to other applications and learning Go of course.