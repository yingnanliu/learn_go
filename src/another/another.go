package another

import (
  "fmt"
  "shared"
)

func Print() {
  fmt.Printf("%s %s\n", shared.FirstName(), shared.LastName())
}