package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age,omitempty" validate:"gte=0"`
}

func main() {
	u := User{}
	t := reflect.TypeOf(u)
	f, _ := t.FieldByName("Name")
	tag := f.Tag
	jsonTag := tag.Get("json")
	fmt.Println("Name's json tag:", jsonTag)
}
