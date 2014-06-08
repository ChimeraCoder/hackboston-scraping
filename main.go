package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Id   int64
}

var Bob = User{"Bob", 12312}

func main() {
	bts, err := json.Marshal(Bob)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(bts))
}

//END OMIT
