package main

import (
  "fmt" //Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
  "os" //Package os provides a platform-independent interface to operating system functionality.
  "os/exec"
  "math"
)

func main() {
  clearScreen("Welcome to GoGO-CLI")
  var user User
  user = UserFromFile()

  if (user == User{}) {
    register("Please register your data")
  }else {
    login(user)
  }
}

func register(m string) {
  clearScreen(m)
  var u User
  u = registrationView()
  if u.Name == "" || u.Email == "" || u.Password == "" || u.Phone == ""{
    register("Please fill all fields")
  }else{
    u.Save()
    main_menu(u, "Welcome to GoGO-CLI")
  }
}

func login(u User) {
  clearScreen("Please enter your login detail")
  var output map[string]string
  output = loginView()

  switch cred_match(output, u) {
    case true:
      main_menu(u, "Welcome back to GoGO-CLI")
    case false:
      login(u)
  }
}

func cred_match(log map[string]string, u User) bool {
  if (log["login"] == u.Email || log["login"] == u.Phone) && (log["password"] == u.Password) {
    return true
  } else {
    return false
  }
}

func main_menu(u User, m string) {
  clearScreen(m)
  var steps int
  steps = view_main_menu()

  switch steps {
    case 1:
      profile(u)
    case 2:
      orderGoride(u)
    case 3:
      orderHistory(u)
    case 4:
      os.Exit(1)
    default:
      main_menu(u, m)
  }
}

func profile(u User) {
  clearScreen("Your Profile")
  var steps int
  steps = view_profile(u)

  switch steps {
    case 1:
      register("Edit your profile")
    case 2:
      main_menu(u, "")
    default:
      profile(u)
  }
}

func orderGoride(u User) {
  clearScreen("Order Go-Ride")
  var o, d = orderGorideView()
  fmt.Println(o)
  fmt.Println(d)
  origin := getLocation(o)
  destination := getLocation(d)
  distance := calcDistance(origin.Coord, destination.Coord)
  price := calcEstPrice(distance)

  order := Order{
    Origin: origin.Name,
    Destination: destination.Name,
    Distance: distance,
    Price: price,
  }
  if origin.Name == "" || destination.Name == ""{
    main_menu(u, "Failed: Location not found")
  }else if distance == 0{
    main_menu(u, "Failed: Origin and destination must be different")
  }else{
    orderGorideConfirm(u, order)
  }
}

func orderGorideConfirm(u User, o Order) {
  clearScreen("Confirm your order")
  steps := orderGorideConfirmView(o)

  switch steps {
    case 1:
      o.Save()
      main_menu(u, "Success: Order saved")
    case 2:
      orderGoride(u)
    case 3:
      main_menu(u, "")
    default:
      orderGorideConfirm(u, o)
  }
}

// func orderResult(u User, m string) {
//   clearScreen(m)
//   steps := orderResultView()

//   switch steps {
//     case 1:
//       orderGoride(u)
//     case 2:
//       main_menu(u)
//     default:
//       orderResult(u, m)
//   }
// }

func orderHistory(u User) {
  clearScreen("Your order history")
  orders :=  ordersFromFile()
  steps := orderHistoryView(orders)
  switch steps {
    case 1:
      main_menu(u, "")
    default:
      orderHistory(u)
  }
}

func clearScreen(m string) {
  c:= exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
  if m != "" {
    fmt.Println(m)
  }
}

func calcDistance (o [2]float64, d [2]float64) (dist float64) {
  dist = math.Pow(o[0] - d[0], 2) + math.Pow(o[1] - d[1], 2)
  dist = math.Sqrt(dist)
  return
}

func calcEstPrice(dist float64) (p float64) {
  p = dist * 1500
  return
}