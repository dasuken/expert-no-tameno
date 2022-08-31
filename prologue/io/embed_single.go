package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed sample.json
var userBytes []byte

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func main() {
	u := User{}
	if err := json.Unmarshal(userBytes, &u); err != nil {
		panic(err)
	}


	fmt.Printf("%+v\n", u)
}
