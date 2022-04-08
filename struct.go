package forum

import (
	"fmt"
	"net/http"
	_ "github.com/google/uuid"
)

var Posts []Post
var Users []User
var Coment []Comment
var Data Send
var Category []string

type Send struct {
	Post []Post
	User User
	PostCategory []string
}

type User struct {
	Id       int
	Uid      string
	Email    string
	Passwd   string
	Username string
}

type Post struct {
	Pid      int
	Category string
	Content  string
	Comments []string
	Id       int
}

type Comment struct {
	Cid     int
	Content string
	Id      int
}

type Cookie struct {
	UUid  string
	Name  string
	Value string
}

func DataToStruct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API")
}
