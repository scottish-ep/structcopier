package main

import (
    "fmt"

    "database/sql"
    "github.com/scottish-ep/structcopier"
)

// Model
type User struct {
    // Basic string field
    Name  string
    // structcopier supports https://golang.org/pkg/database/sql/driver/#Valuer
    Email sql.NullString
}

func (u *User) MethodThatTakesContext(ctx map[string]interface{}) string {
    // do whatever you want
    return "hello from this method"
}

// Resource
type UserResource struct {
    DisplayName            string `structcopier:"field:Name"`
    SkipMe                 string `structcopier:"skip"`
    MethodThatTakesContext string `structcopier:"context"`
    Email                  string `structcopier:"force"`

}

func main() {
    user := &User{
        Name: "gilles",
        Email: sql.NullString{
            Valid: true,
            String: "gilles@example.com",
        },
    }

    resource := &UserResource{}
    filters := map[string]interface{}{
        "fields" : []string{"Name"},
    }
    structcopier.Copy(user).Filter(filters).To(resource)
    // structcopier.Copy(user).To(resource)

    fmt.Println(resource.DisplayName)
    fmt.Println(resource.Email)
}
