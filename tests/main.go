package main

import (
    "fmt"

    "database/sql"
    "github.com/scottish-ep/deepcopier"
)

// Model
type User struct {
    // Basic string field
    Name  string
    // Deepcopier supports https://golang.org/pkg/database/sql/driver/#Valuer
    Email sql.NullString
}

func (u *User) MethodThatTakesContext(ctx map[string]interface{}) string {
    // do whatever you want
    return "hello from this method"
}

// Resource
type UserResource struct {
    DisplayName            string `deepcopier:"field:Name"`
    SkipMe                 string `deepcopier:"skip"`
    MethodThatTakesContext string `deepcopier:"context"`
    Email                  string `deepcopier:"force"`

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
    deepcopier.Copy(user).Filter(filters).To(resource)
    // deepcopier.Copy(user).To(resource)

    fmt.Println(resource.DisplayName)
    fmt.Println(resource.Email)
}
