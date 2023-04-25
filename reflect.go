package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type OtherSample struct {
	Name string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"Wawan"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	requiredTag := structField.Tag.Get("required")
	maxTag := structField.Tag.Get("max")

	fmt.Println(sampleType.NumField())
	fmt.Println(structField.Name)
	fmt.Println(requiredTag)
	fmt.Println(maxTag)

	// sample.Name = ""
	fmt.Println(IsValid(sample))

	otherSample := OtherSample{""}
	fmt.Println(IsValid(otherSample))
}
