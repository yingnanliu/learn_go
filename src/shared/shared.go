package shared

type name struct {
  first, last string
}

var globalName name

func Set(first, last string) {
  globalName.first = first
  globalName.last = last
}

func FirstName() string {
  return globalName.first
}

func LastName() string {
  return globalName.last
}