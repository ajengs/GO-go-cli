package main

import (
  "encoding/json"
  "fmt"
  "os"
  "io/ioutil"
)

type User struct {
  Name string
  Password string
  Phone string
  Email string
}

func userF() string {
  return "users.json"
}

func (u User) print() {
  fmt.Println("Name: ", u.Name)
  fmt.Println("Password: ", u.Password)
  fmt.Println("Phone: ", u.Phone)
  fmt.Println("Email: ", u.Email)
}

// Convert to json
func (u User) toJson() []byte {
  bs, err := json.Marshal(u)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return bs
}

// Save to file in json format
func (u User) Save() error {
  return ioutil.WriteFile(userF(), u.toJson(), 0644)
}

// Load user from file
func UserFromFile() User {
  bs, err := ioutil.ReadFile(userF())
  if err != nil {
    return User{}
  }
  return userFromJson(bs)
}

// Read from json
func userFromJson(bs []byte) User {
  var u User
  err := json.Unmarshal(bs, &u)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return u
}