package main

import "fmt"

type contact struct {
  phone string
  email string
}

type person struct {
  name string
  age int
  contact contact
}

func (p person) print() {
  fmt.Println("name: ", p.name)
  fmt.Println("age: ", p.age)
  fmt.Println("phone: ", p.contact.phone)
  fmt.Println("email: ", p.contact.email)
}
