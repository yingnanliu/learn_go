package printer

import (
  "fmt"
  "storage"
)

func Print() {
  fmt.Printf("%s %s\n", storage.FirstName(), storage.LastName())
}