package main

import (
	"log"

	"github.com/astaxie/beego/validation"
)

type User struct {
	ID     int
	Name   string //`valid:"Required; MaxSize(255)"`
	Age    int    //`valid:"Range(1, 140)`
	Email  string //`valid:"Email; MaxSize(100)"`
	Mobile string //`valid:"Mobile"`
	IP     string //`valid:"IP"`
}

func main() {
	u := User{
		Name:   "Beego",
		Age:    200,
		Email:  "dev@beego.me",
		Mobile: "15960389468",
	}

	valid := validation.Validation{}
	valid.MaxSize(u.Name, 15, "nameMax")
	valid.Range(u.Age, 0, 18, "age")
	valid.IP(u.IP, "ip")
	valid.Email(u.Email, "email")
	valid.Mobile(u.Mobile, "mobile")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	if v := valid.Max(u.Age, 140, "age"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
	}
}
