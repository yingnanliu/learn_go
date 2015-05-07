package main

import (
  "shared"
  "another"
)

func main() {
  shared.Set("Foo", "Bar")
  another.Print()

  shared.Set("Jane", "Doe")
  another.Print()
}