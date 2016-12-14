package main

import (
	"reflect"
	"fmt"
)

type T struct {
	A int
	B string
}

func main(){
	t := T{203, "hello"}
	s:= reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
