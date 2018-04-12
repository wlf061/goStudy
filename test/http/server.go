package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age int `json:age`
}

func index(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("Form: ", r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println(r.Form["a"])
	fmt.Println(r.Form["b"])

	for k, v := range r.Form{
		fmt.Println(k,"=>",v,strings.Join(v, "-"))
	}
	fmt.Fprint(w,"It works!")

}

func test(w http.ResponseWriter, r *http.Request){
	body, _:=ioutil.ReadAll(r.Body)
	body_str := string(body)
	fmt.Println(body_str)
	var user User
	err := json.Unmarshal(body,&user)
	if err == nil {
		fmt.Println(user)
		user.Age += 100
		fmt.Println(user)
		ret, _ := json.Marshal(user)
		fmt.Println("ret="+string(ret))
		fmt.Fprint(w, string(ret))
	} else {
		fmt.Println(err)
	}



}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test/", test)

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
