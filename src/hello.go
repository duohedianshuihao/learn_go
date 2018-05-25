package main

/*
- note the package import but not used would be deleted when saving the file, so don't save to often
- use import std "fmt" to change the name of package to std
*/
import (
	"fmt"
	"strconv"
)

// PI is const
const PI = 3.14 // const comment should use same format

// no comma here
const (
	const1 int    = 'A'
	const2        = "2"
	const5        // could use default value in up line, which is "2"
	const3 = iota // only works in this group
	const4
)

// save unit
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	ZB
	YB
)

var name = "golbal name"

type newType int // a new name for newType

type gopher struct{} // declare structure

type golang interface{} // declare interface

func main() {
	a, _, c, d := 1, 2, 3, 4 // use _ to ignore sth
	fmt.Println("Hello world!")
	fmt.Println(a, c, d)
	e := float32(c) // type convert, always in obvious way
	fmt.Println(e)

	var f = 65
	g := string(f)
	h := strconv.Itoa(f)
	fmt.Println(g, h)
	fmt.Println(const1, const2, const3, const4)

	fmt.Println(1 << 10)

	var p = &a
	fmt.Println(p)
	fmt.Println(*p)

	a++ // this is good
	// a = a++  this is bad ++ is a sentence not a expression

	// only for in golang, there is no while loop, so for loop could work just like while loop
	for index := 0; index < 2; index++ { // left { must in the same line with for statement
		fmt.Println(index)
	}
	if index := 0; index < 2 { // left { must in the same line with if statement
		index++ // no ++index
		fmt.Println(index)
	}

	// if switch variable is provided, case statement judge if the result of that expression == case value
	// switch get no break, it break when find first satisfied result
	// if wanna continue, use fallthrough
	switch a {
	case 0:
		fmt.Println("a = 1")
		fallthrough
	case 1:
		fmt.Println("a = 2")
	default:
		fmt.Println("None")
	}

	// if switch expression is absent, case statement would take the expression which return true
	switch {
	case a > 0:
		fmt.Println("a > 0")
	case a < 0:
		fmt.Println("a < 0")
	default:
		fmt.Println("None")
	}

	// use expression
	switch a := 1; {
	case a > 0:
		fmt.Println("a > 0")
	case a < 0:
		fmt.Println("a < 0")
	default:
		fmt.Println("None")
	}

	// break, could use with label
LABEL1:
	for {
		for index := 0; index < 4; index++ {
			if index > 2 {
				break LABEL1 // if use goto or continue  here, still be a infinite loop
			}
		}
	}
	fmt.Println("finish loop")

	// array
	var a1 [2]int // length is part of type, a1 a2 is different
	var a2 [1]int // if a2 [2]int, then a1 = a2 is good
	fmt.Println(a1, a2)

	a3 := [2]int{1}      // assign number for array, if values provided less than length, use default value for the rest
	a4 := [20]int{19: 1} // assign particular number using index
	fmt.Println(a3, a4)

	fmt.Println(a1 == a3) // print false
	// fmt.Println(a1 == a2)  this is error, because a1 a3 are not same type

	a5 := [...]int{1, 2, 3, 4, 5} // if array is provided and would not change, could use ... for length
	a6 := [...]int{19: 1}         // length of a6 is 20
	fmt.Println(a5)
	var pr = &a6
	fmt.Println(pr)

	x, y := 1, 2
	prL := [...]*int{&x, &y}
	fmt.Println(prL) // array in go use value rather than reference

	a7 := [10]int{} // {} could not ignore
	a7[1] = 2
	fmt.Println(a7)

	a8 := new([10]int) // a8 is a reference to array
	a8[1] = 2
	fmt.Println(a8)

	a9 := [2][3]int{ // 2 position could use ..., but 3 position could not use ...
		{1, 2, 3},
		{2, 3, 4}} // } should in the same line with {2, 3, 4}
	fmt.Println(a9)

	// slice  is reference not array
	// var s1 []int // without define the length, that is slice
	s2 := a4[5:8]                     // take 5 to 8 from array a4
	fmt.Println(s2, len(s2), cap(s2)) // len is 3, cap is 5
	s3 := make([]int, 3, 10)          // 3 is the length of slice, and 10 is the capacity
	fmt.Println(len(s3), cap(s3))     // is there is more than capacity item in s3, it will automatically double the capacity
	s4 := s2[:2]                      //return the first two item in s2, if take more than len(s2), return value from a4
	fmt.Println(s4)                   // if take more item than original array, it will crash rather than resize

	s5 := make([]int, 3, 6)
	fmt.Printf("%p\n", s5)

	s5 = append(s5, 1, 2, 3) // if after appending size smaller than capacity, nothing happen, if not, resize automatically
	fmt.Printf("%v %p\n", s5, s5)

	s6 := []int{1, 2, 3, 4, 5}
	s7 := []int{7, 8, 9}
	copy(s6, s7) // copy s7 to s6
	fmt.Println(s6)
}

/*
函数名首字母小写： 私有函数
函数名首字母大写： 公有函数
*/

/*
- bool
true, false
0, None, [], {} is not false
*/

/*
- int/uint	u for unsigned
based on platform could be 32 bit or 64 bit
use int8/uint8 to specify -128~127/0~255
byte is uint8
*/

/*
float32/float64
complex64/complex128
uintprt
interface
func
*/

/*
  6: 0110
 11: 1011
 ---------
 &   0010
 |   1111
 ^   1101
 &^  0100

 && is 'and', lasy partten
*/
