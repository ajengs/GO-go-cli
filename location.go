package main

import (
  "encoding/json"
  "fmt"
  "os"
  "io/ioutil"
)

type Location struct {
  Name string
  Coord [2]float64
}

func locationF() string {
  return "locations.json"
}

func (l Location) print() {
  fmt.Println("Name: ", l.Name)
  fmt.Println("Coordinates: ", l.Coord)
}

// Convert to json
func (l Location) toJson() []byte {
  bs, err := json.Marshal(l)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return bs
}

// Save to file in json format
func (l Location) Save() error {
  return ioutil.WriteFile(locationF(), l.toJson(), 0644)
}

// Load location from file
func locationsFromFile() []Location {
  bs, err := ioutil.ReadFile(locationF())
  if err != nil {
    return []Location{}
  }

  return locationsFromJson(bs)
}

// Read from json
func locationsFromJson(bs []byte) []Location {
  var ls []Location
  err := json.Unmarshal(bs, &ls)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return ls
}

func getLocation(n string) Location {
  var ls []Location
  ls = locationsFromFile()
  var loc Location
  for _, l := range ls {
    if l.Name == n {
      loc.Name = l.Name
      loc.Coord = l.Coord
    }
  }
  return loc
}