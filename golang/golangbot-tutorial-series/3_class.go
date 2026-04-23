package main

//import (
//	"fmt"
//	"unsafe"
//)
//
//func main() {
//	a := true
//	b := false
//	fmt.Println("a:", a, "b:", b)
//	c := a && b
//	fmt.Println("c:", c)
//	d := a || b
//	fmt.Println("d:", d)
//
//	//int8	8 bit signed integers	8 bits	-128 to 127
//	//int16	16 bit signed integers	16 bits	-32768 to 32767
//	//int32	32 bit signed integers	32 bits	-2147483648 to 2147483647
//	//int64	64 bit signed integers	64 bits	-9223372036854775808 to 9223372036854775807
//	//int	represents 32 or 64 bit integers depending on the underlying architecture. You should generally be using int to represent integers unless there is a need to use a specific sized integer.	32 bits in 32 bit systems and 64 bit in 64 bit systems.	-2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems
//
//	var a1 int = 89
//	b1 := 95
//	fmt.Println("value of a is", a1, "and b is", b1)
//
//	var a2 = 89
//	b2 := 95
//	fmt.Println("value of a is", a2, "and b is", b2)
//	fmt.Printf("type of a is %T, size of a is %d bytes", a2, unsafe.Sizeof(a2))   //type and size of a
//	fmt.Printf("\ntype of b is %T, size of b is %d bytes", b2, unsafe.Sizeof(b2)) //type and size of b
//
//	//uint8	8 bit unsigned integers	8 bits	0 to 255
//	//uint16	16 bit unsigned integers	16 bits	0 to 65535
//	//uint32	32 bit unsigned integers	32 bits	0 to 4294967295
//	//uint64	64 bit unsigned integers	64 bits	0 to 18446744073709551615
//	//uint	32 or 64 bit unsigned integers depending on the underlying architecture	32 bits in 32 bit systems and 64 bits in 64 bit systems	0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
//
//	var a3 uint = 60
//	var b3 uint = 30
//	c3 := a3 * b3
//	fmt.Println("c =", c3)
//	fmt.Printf("Data type of variable c is %T", c3)
//
//	//float32	32 bit floating point numbers
//	//float64	64 bit floating point numbers
//	a, b := 5.67, 8.97
//	fmt.Printf("type of a %T b %T\n", a, b)
//	sum := a + b
//	diff := a - b
//	fmt.Printf("sum of %f and %f is %f, diff is %f\n", a, b, sum, diff)
//
//	no1, no2 := 56, 89
//	fmt.Printf("type of no1 %T no2 %T\n", no1, no2)
//	fmt.Printf("sum of %d and %d is %d, diff is %d", no1, no2, no1+no2, no1-no2)
//
//	c1 := complex(5, 7)
//	c2 := 8 + 27i
//	cadd := c1 + c2
//	fmt.Println("sum:", cadd)
//	cmul := c1 * c2
//	fmt.Println("product:", cmul)
//
//	first := "Naveen"
//	last := "Ramanathan"
//	name := first + " " + last
//	fmt.Println("My name is", name)
//
//	a := 80
//	b := 91.8
//	sum := a + int(b)
//	fmt.Println(sum)
//
//	i := 10
//	var j float64 = float64(i)
//	fmt.Println("j =", j)
//}
