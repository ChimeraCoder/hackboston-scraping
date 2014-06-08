package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Id   int64
}

var AliceJson = `{"name" : "Alice", "id" : 121312}`

func main() {
	var result User
	err := json.Unmarshal([]byte(AliceJson), &result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", result)
}

//END OMIT
