package common

import "fmt"

type Hello struct {
	World string `json:"world"`
}

func (h Hello) Print() string {
	return "Golang"
}

func (h Hello) Println(value string) string {
	return fmt.Sprintf("%s", value)
}
