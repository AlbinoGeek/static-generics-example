package main

//go:generate genny -in=foo/template.go -out=List.go gen "Item=string,int,int32,int64,bool"
//go:generate genny -in=foo/template_test.go -out=List_test.go gen "Item=string,int,int32,int64,bool"

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
