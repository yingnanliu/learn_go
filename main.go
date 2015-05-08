package main

import (
  "storage"
  "printer"
  "fmt"
)

func main() {
  storage.Set("Foo", "Bar")
  printer.Print()

  storage.Set("Jane", "Doe")
  printer.Print()

  n := storage.NameInMap()
  fmt.Printf("%s %s\n", n["first"], n["last"])

  n["first"] = "boom"

  changedName := storage.NameInMap()
  fmt.Printf("%s %s\n", changedName["first"], changedName["last"])
}