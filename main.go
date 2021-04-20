package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	name, password, email string
}

type Role struct {
	name string
}

type Profile struct {
	user                User
	firstName, lastName string
	role                Role
}

func (p Profile) getUsername() (username string) {
	return p.user.name
}

func (p Profile) checkPassword(somePassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(p.user.password), []byte(somePassword)); err != nil {
		return false, err
	}
	return true, nil
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println(w)
	var profile Profile
	body := json.NewDecoder(r.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err := body.Decode(&profile)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(body)
	fmt.Println(&body)
	w.WriteHeader(http.StatusCreated)
}

// func(u User) createUser(user struct) (bool, error) {

// }

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func main() {
	// text := []byte("some password")
	// hash, err := bcrypt.GenerateFromPassword(text, bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Hash to be stored: ", hash)
	// getHash := "some password"
	// hashfromDatabase := hash
	// if err := bcrypt.CompareHashAndPassword(hashfromDatabase, []byte(getHash)); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Password was correct")
	// fmt.Println("Creating new user")
	// newUser := Profile{
	// 	user: User{
	// 		name:     "wesley",
	// 		password: string(hash),
	// 		email:    "www@gmail.com",
	// 	},
	// 	firstName: "Wesley",
	// 	lastName:  "Mutwiri",
	// 	role: Role{
	// 		name: "super_admin",
	// 	},
	// }
	// fmt.Printf("%v \n", newUser)
	// fmt.Println("New User successfully created here")
	// fmt.Println("Username: ", newUser.getUsername())
	// fmt.Printf("Checking Password: ")
	// value, err := newUser.checkPassword("some password")
	// fmt.Println(value, err)
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/user", CreateUser)
	log.Fatal(http.ListenAndServe(":10000", router))
}
