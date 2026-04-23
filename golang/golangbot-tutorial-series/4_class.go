package main

//import (
//	"fmt"
//)
//
//func main() {
//	const a = 50
//	fmt.Println(a)
//
//	const (
//		retryLimit = 4
//		httpMethod = "GET"
//	)
//
//	fmt.Println(retryLimit)
//	fmt.Println(httpMethod)
//
//	var a = math.Sqrt(4) //allowed
//	fmt.Println(a)
//	const b = math.Sqrt(4) //not allowed
//	fmt.Println(b)
//
//	const n = "Sam"
//	var name = n
//	fmt.Printf("type %T value %v", name, name)
//
//	var defaultName = "Sam" //allowed
//	type myString string
//	var customName myString = "Sam" //allowed
//	customName = defaultName        //not allowed
//	fmt.Println(customName)
//
//	var defaultName = "Sam" //allowed
//	type myString string
//	var customName myString = "Sam"    //allowed
//	customName = myString(defaultName) //allowed
//	fmt.Println(customName)
//
//	const trueConst = true
//	type myBool bool
//	var defaultBool = trueConst //allowed
//	var customBool myBool = trueConst //allowed
//	defaultBool = customBool //not allowed
//
//	const c = 5
//	var intVar int = c
//	var int32Var int32 = c
//	var float64Var float64 = c
//	var complex64Var complex64 = c
//	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
//
//	var i = 5
//	var f = 5.6
//	var c = 5 + 6i
//	fmt.Printf("i's type is %T, f's type is %T, c's type is %T", i, f, c)
//
//	const c = 5
//	var intVar int = c
//	var int32Var int32 = c
//	var float64Var float64 = c
//	var complex64Var complex64 = c
//	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
//
//	var a = 5.9 / 8
//	fmt.Printf("a's type is %T and value is %v", a, a)
//}
