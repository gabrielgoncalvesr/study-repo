package main

import (
	"fmt"
	"learngo/finance/simpleinterest"
	"log"
	//_ "learnpackage/simpleinterest"
)

var p, r, t = 5000.0, 10.0, 1.0

var _ = simpleinterest.Calculate

func init() {
	fmt.Println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {
	//fmt.Println("Simple interest calculation")
	//p := 5000.0
	//r := 10.0
	//t := 1.0
	//si := simpleinterest.Calculate(p, r, t)
	//fmt.Println("Simple interest is", si)

	fmt.Println("Simple interest calculation")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
