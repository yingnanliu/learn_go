package main

import (
  "storage"
  "printer"
)

func main() {
  storage.Set("Foo", "Bar")
  printer.Print()

  storage.Set("Jane", "Doe")
  printer.Print()
}