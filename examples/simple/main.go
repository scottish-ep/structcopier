package main

import (
	"fmt"
	"github.com/scottish-ep/structcopier"
)

// Model
type User struct {
	Name string
}

func (u *User) MethodThatTakesContext(ctx map[string]interface{}) string {
	// do whatever you want
	return ""
}

// Resource
type UserResource struct {
	DisplayName            string `structcopier:"field:Name"`
	SkipMe                 string `structcopier:"skip"`
	MethodThatTakesContext string `structcopier:"context"`
}

func main() {
	user := &User{
		Name: "gilles",
	}

	resource := &UserResource{}

	structcopier.Copy(user).To(resource)

	fmt.Println(resource.DisplayName)
}
