package main

import (
  "encoding/json"
  "fmt"
  "os"
  "io/ioutil"
  // "time"
)

type Order struct {
  Origin string
  Destination string
  Distance float64
  Price float64
  // Timestamp time.Time
}

func orderF() string {
  return "orders.json"
}

func (o Order) print() {
  fmt.Println("Origin: ", o.Origin)
  fmt.Println("Destination: ", o.Destination)
  fmt.Println("Distance: ", o.Distance)
  fmt.Println("Price: ", o.Price)
}

// Convert to json
func (o Order) toJson() []byte {
  bs, err := json.Marshal(o)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return bs
}

func toJson(ors []Order) []byte {
  bs, err := json.Marshal(ors)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return bs
}

// Save to file in json format
func (o Order) Save() error {
  ors := ordersFromFile()
  ors = append(ors, o)
  return ioutil.WriteFile(orderF(), toJson(ors), 0644)
}

// Load from file
func ordersFromFile() []Order {
  bs, err := ioutil.ReadFile(orderF())
  if err != nil {
    return []Order{}
  }
  return ordersFromJson(bs)
}

// Read from json
func ordersFromJson(bs []byte) []Order {
  var ors []Order
  err := json.Unmarshal(bs, &ors)
  if err != nil {
    fmt.Println("ordersFromJson:", err)
    os.Exit(1)
  }
  return ors
}
