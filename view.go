package main

import (
  // "encoding/json"
  "fmt"
  "bufio"
  "os"
)

func registrationView() (u User) {
  fmt.Print("Full Name: ")

  fmt.Scanln(&u.Name)

  fmt.Print("Password: ")
  fmt.Scanln(&u.Password)

  fmt.Print("Email: ")
  fmt.Scanln(&u.Email)

  fmt.Print("Phone: ")
  fmt.Scanln(&u.Phone)

  return
}

func loginView() map[string]string {
  var l, p string
  output := make(map[string]string)

  fmt.Print("Login: ")
  fmt.Scanln(&l)

  fmt.Print("Password: ")
  fmt.Scanln(&p)

  output["login"] = l
  output["password"] = p
  return output
}

func view_main_menu() (step int) {
  fmt.Println("Main menu: ")
  fmt.Println("1. View Profile")
  fmt.Println("2. Order Go Ride")
  fmt.Println("3. Order History")
  fmt.Println("4. Exit")
  fmt.Print("Enter your option: ")

  fmt.Scanln(&step)
  return
}

func view_profile(u User) (step int) {
  // fmt.Printf("Full Name: %s\n", u.Name)
  // fmt.Printf("Email: %s\n", u.Email)
  // fmt.Printf("Phone: %s\n", u.Phone)
  // fmt.Printf("Password: %s\n", u.Password)
  u.print()
  fmt.Println("")
  fmt.Println("1. Edit Profile")
  fmt.Println("2. Back")
  fmt.Print("Enter your option: ")
  
  fmt.Scanln(&step)
  return
}

func orderGorideView() (o string, d string){
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Printf("Origin: ")
  // fmt.Scanf("%s", &o)
  scanner.Scan()
  o = scanner.Text()

  fmt.Printf("Destination: ")
  scanner.Scan()
  d = scanner.Text()

  return
}

func orderGorideConfirmView(o Order) (step int) {
  fmt.Printf("Origin: %s\n", o.Origin)
  fmt.Printf("Destination: %s\n", o.Destination)
  fmt.Printf("Distance: %f km\n", o.Distance)
  fmt.Printf("Est. Price: %f\n", o.Price)

  fmt.Println("")
  fmt.Println("1. Save Order")
  fmt.Println("2. Re-Order")
  fmt.Println("3. Back to main menu")
  fmt.Print("Enter your option: ")
  
  fmt.Scanln(&step)
  return
}

// func orderResultView() (step int) {
//   // fmt.Println(m)
//   fmt.Println("")
//   fmt.Println("1. Re-Order")
//   fmt.Println("2. Back to main menu")
//   fmt.Print("Enter your option: ")
  
//   fmt.Scanln(&step)
//   return
// }

func orderHistoryView(ors []Order) (step int) {
  for _, o := range ors {
    // fmt.Println("Origin: ", o.Origin)
    // fmt.Println("Destination: ", o.Destination)
    // fmt.Println("Distance: ", o.Distance)
    // fmt.Println("Price: ", o.Price)
    o.print()
    fmt.Println("-----------------------------------------")
  }

  fmt.Println("")
  fmt.Println("1. Back to main menu")
  fmt.Print("Enter your option: ")
  
  fmt.Scanln(&step)
  return
}