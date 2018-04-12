package main

import (
	"encoding/json"
	"fmt"
)
  /* `json:"user_id" bson:"user_id"` 用来标记为 tag*/
type User struct {
	UserId   int    `json:"user_id" bson:"user_id"`
	UserName string `json:"user_name" bson:"user_name"`
}

func main() {

	u := User{UserId: 1, UserName: "tony"}
	body, _ :=json.Marshal(u)
	fmt.Println(string(body))

}
