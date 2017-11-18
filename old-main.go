package main

import (
    // "bufio" //Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
    "fmt" //Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
    // "os" //Package os provides a platform-independent interface to operating system functionality.
)

func old_main() {
  ajeng := person{
    name: "Ajeng",
    age: 25,
    contact: contact{
      phone: "08118955088",
      email: "ajeng.sugiarti@gmail.com",
    },
  }
  ajeng.print()

  x := 5
  fmt.Println(&x)
  calculate(&x)
  fmt.Println(x)
}

func calculate(x *int){
  fmt.Println(x)
  *x = 7
}
