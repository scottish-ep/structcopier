# structcopier

[![Build Status](https://secure.travis-ci.org/ulule/structcopier.svg?branch=master)](http://travis-ci.org/ulule/structcopier)

This package is copy from deepcopier of [Ulule team](https://github.com/ulule).
Due to deepcopier hasn't updated for a long time, I created this repo to fix some bugs and extend some new features. structcopier is meant to make copying of structs to/from others structs a bit easier.

## Installation

```bash
go get -u github.com/scottish-ep/structcopier
```

## Usage

```golang
// Deep copy instance1 into instance2
Copy(instance1).To(instance2)

// Deep copy instance1 into instance2 and passes the following context (which
// is basically a map[string]interface{}) as first argument
// to methods of instance2 that defined the struct tag "context".
Copy(instance1).WithContext(map[string]interface{}{"foo": "bar"}).To(instance2)

// Deep copy instance2 into instance1
Copy(instance1).From(instance2)

// Deep copy instance2 into instance1 and filter some fields
Copy(instance1).Filter(filters).To(instance2)

// Deep copy instance2 into instance1 and passes the following context (which
// is basically a map[string]interface{}) as first argument
// to methods of instance1 that defined the struct tag "context".
Copy(instance1).WithContext(map[string]interface{}{"foo": "bar"}).From(instance2)
```

Available options for `structcopier` struct tag:

| Option    | Description                                                          |
| --------- | -------------------------------------------------------------------- |
| `field`   | Field or method name in source instance                              |
| `context` | Takes a `map[string]interface{}` as first argument (for methods)     |
| `force`   | Set the value of a `sql.Null*` field (instead of copying the struct) |
| `filter`  | Ignore some fields when copy                                         |

**Options example:**

```golang
type Source struct {
    Name                         string
    SkipMe                       string
    SQLNullStringToSQLNullString sql.NullString
    SQLNullStringToString        sql.NullString

}

func (Source) MethodThatTakesContext(c map[string]interface{}) string {
    return "whatever"
}

type Destination struct {
    FieldWithAnotherNameInSource      string         `structcopier:"field:Name"`
    SkipMe                            string         `structcopier:"skip"`
    MethodThatTakesContext            string         `structcopier:"context"`
    SQLNullStringToSQLNullString      sql.NullString
    SQLNullStringToString             string         `structcopier:"force"`
}

```

Example:

```golang
package main

import (
    "fmt"

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
    // copy all fields from user to resource based on struct tag
    structcopier.Copy(user).To(resource)

    // ignore name when copy from user to resource
    filters := map[string]interface{}{
        "fields" : []string{"Name"},
    }
    structcopier.Copy(user).Filter(filters).To(resource)

    // skip
    fmt.Println(resource.DisplayName)
    fmt.Println(resource.Email)
}
```
