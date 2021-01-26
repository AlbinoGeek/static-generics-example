package main

//go:generate genny -in=foo/template.go -out=List.go gen "Item=string,int,int8,int16,int32,int64,bool,uint,uint8,uint16,uint32,uint64,float32,float64,complex64,complex128"
//go:generate genny -in=foo/template_test.go -out=List_test.go gen "Item=string,int,int8,int16,int32,int64,bool,uint,uint8,uint16,uint32,uint64,float32,float64,complex64,complex128"

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
