package storage

type name struct {
  first, last string
}

var globalName name

var nameInMap = map[string]string{
  "first": "hello",
  "last": "world",
}

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

func NameInMap() map[string]string {
  return nameInMap
}