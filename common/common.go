package common

import "fmt"

/*
Hello with field world
*/
type Hello struct {
	World string `json:"world"`
}

/*
Hello with method Print
*/
func (h Hello) Print() string {
	return "Golang"
}

/*
Hello with method Println
*/
func (h Hello) Println(value string) string {
	return fmt.Sprintf("%s", value)
}
