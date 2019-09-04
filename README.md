# JSONull

This library allow JSON Unmarshalling and Marshalling `null` value.

## Why this library was created

### nil

```json
{
  "first_name": "thanakrit",
  "last_name": "lee"
}
```

```go
type Person struct {
  FirstName string  `json:"first_name"`
  LastName  string  `json:"last_name"`
  Age       int     `json:"age"`
}
```

When unmarshalling a JSON into a struct that doesn't include one of the fields in the struct, the unincluded fields are set to the `nil` value of the specified type when it is unmarshalled. In the above case, because `age` was not included, it would be set to `0` because `0` is the `nil` value of the `int` type.

These default nil values of each type can cause a problem when one wants to check whether a recieved JSON property is null/nil or not.

### null

```json
{
  "first_name": "thanakrit",
  "last_name": null
}
```

```go
type Person struct {
  FirstName string  `json:"first_name"`
  LastName  string  `json:"last_name"`
  Age       int     `json:"age"`
}
```

In the case where the marshalled JSON has a field that is explicitly set to `null`, unmarshalling it into your struct would result in the `nil` value of the specified type as before.

## Solution

This library provides new types which implements the `encoding/json` `Marshaler` and `Unmarshaler` interface to check whether a marshalled JSON field is `null` implicitly or explicitly.

The new types are struct containing the following 3 fields:
- Value - The value of the property. If the property is set to `null`, then the value will be the `nil` value of the specified type e.g. `0 (int)`, `"" (string)`, `false (bool)`.
- Valid - A boolean indicating whether the property is set to `null` both implicitly or explicitly.
- Set - A boolean indicating whether the property is set to `null` explicitly. `true` to indicate that it was set to `null` explicitly, and `false` if otherwise.


## Installation

Use `go get` to install the library.
```sh
go get github.com/thanakritlee/jsonull
```


## Usage

```go
package main

import (
  jFloat "github.com/thanakritlee/jsonull/lib/float"
  jInt "github.com/thanakritlee/jsonull/lib/int"
  jString "github.com/thanakritlee/jsonull/lib/string"

  "encoding/json"
  "fmt"  
)

type Person struct {
  FirstName jString.String  `json:"first_name"`
  LastName  jString.String  `json:"last_name"`
  Age       jInt.Int        `json:"age"`
  Height    jFloat.Float32  `json:"height"`
}

func main() {
  personMarshalledString := `{"first_name":"thanakrit","last_name":null}`
  
  var personStruct Person

  err := json.Unmarshal([]byte(personMarshalled), &personStruct)
  if err != nil {
    fmt.Errorf(err.Error())
  }

  // Check FirstName.
  if personStruct.FirstName.Valid {
    fmt.Println("FirstName: " + personStruct.FirstName.Value)
  } else if personStruct.FirstName.Set {
    fmt.Println("FirstName was set as null")
  } else {
    fmt.Println("FirstName wasn't given")
  }

  // Check LastName.
  if personStruct.LastName.Valid {
    fmt.Println("LastName: " + personStruct.LastName.Value)
  } else if personStruct.LastName.Set {
    fmt.Println("LastName was set as null")
  } else {
    fmt.Println("LastName wasn't given")
  }

  // Check Age.
  if personStruct.Age.Valid {
    fmt.Println("Age: " + fmt.Sprintf("%d", personStruct.Age.Value))
  } else if personStruct.Age.Set {
    fmt.Println("Age was set as null")
  } else {
    fmt.Println("Age wasn't given")
  }

  // Check Height.
  if personStruct.Height.Valid {
    fmt.Println("Height: " + fmt.Sprint("%f", personStruct.Height.Value))
  } else if personStruct.Height.Set {
    fmt.Println("Height was set as null")
  } else {
    fmt.Println("Height wasn't given")
  }

  // Marshal the struct.
  personMarshalledBytes, err := json.Marshal(personStruct)
  if err != nil {
    fmt.Errorf(err.Error())
  }

  // Any struct fields that is null, or wasn't given will be set as null,
  // when it is marshalled.
  fmt.Println(string(personMarshalledBytes))
  
}
```


## Inspiration

[How to determine if a JSON key has been set to null or not provided - Jon Calhoun](https://www.calhoun.io/how-to-determine-if-a-json-key-has-been-set-to-null-or-not-provided/)