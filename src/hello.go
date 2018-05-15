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
	const1 int = 'A'
	const2     = "2"
	const3     = iota // only works in this group
	const4
)

var name = "golbal name"

type newType int // a new name for newType

type gopher struct{} // declare structure

type golang interface{} // declare interface

func main() {
	a, _, c, d := 1, 2, 3, 4 // use _ to ignore sth
	fmt.Println("Hello world!")
	fmt.Println(a, c, d)
	e := float32(c) // type convert
	fmt.Println(e)

	var f = 65
	g := string(f)
	h := strconv.Itoa(f)
	fmt.Println(g, h)
	fmt.Println(const1, const2, const3, const4)

	fmt.Println(1 << 10)
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
