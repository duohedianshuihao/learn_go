package main

/*
- note the package import but not used would be deleted when saving the file, so don't save to often
- use import std "fmt" to change the name of package to std
*/
import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"sync"
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
	BIT float64 = 1 << (iota * 10)
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

// ASDF 1234
type ASDF int

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

	// slice is reference not array
	var s1 []int // without define the length, that is slice
	fmt.Println(s1)
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

	for i, v := range s7 {
		fmt.Println(i, v-10) // this would NOT change value in s7, because it use reference
	} // or use s7[i] to modify value in s7

	// map
	var m1 map[int]string // key type: int	value type: string
	m1 = map[int]string{}
	m2 := make(map[int]string)
	m1[1] = "OJBK" // m1[1] already defined type
	fmt.Println(m1[1], m2[1])
	delete(m1, 1)
	fmt.Println(m1[1])

	var complexMap map[int]map[int]string
	complexMap = make(map[int]map[int]string)
	// complexMap[1][1] = "OK"  this is wrong, make only initialize the first map not the map inside
	complexMap[1] = make(map[int]string)
	complexMap[1][1] = "OJBK"
	fmt.Println(complexMap)
	content, ok := complexMap[2][1] // multi-value return content and if the map exist; if the value not exist, content use default value
	fmt.Println(content, ok)

	m1[1] = "OJBK"
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	ms := make([]map[int]string, 5) // map[int]string here is to define the content type of array
	for _, v := range ms {
		v = make(map[int]string) // this to prove v is only COPY, and won't change anything in ms
		v[1] = "OJBK"
		fmt.Println(v)
	}
	fmt.Println(ms)

	for i := range ms {
		ms[i] = make(map[int]string)
		ms[i][1] = "OJBK"
	}
	fmt.Println(ms)

	m1 = map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	sortSlice := make([]int, len(m1))
	i := 0
	for k := range m1 {
		sortSlice[i] = k
		i++
	}
	sort.Ints(sortSlice) // put the keys of map into a slice and sort the slice, which equally sort the map
	fmt.Println(sortSlice)

	learn3('a', 'b', 'c')
	num1 := 1
	num2 := 2
	learn4(num1, &num2)
	fmt.Println("original:", num1, num2)

	uname := func() { // this is anonymous function, like lambda in python
		fmt.Println("this is anonymous function")
	}
	uname()

	learnClosure := closure(5)
	fmt.Println("returning function run continuesly: ", learnClosure(1))

	// defer   call in reverse order
	for i := 0; i < 3; i++ {
		//this will be called last, because defer run in reverse order
		defer fmt.Println(i) // i here passed value in function
		defer func() {
			/*
				 i in anonymous func is a copy to local var here, so i is a reference, when the loop end,
				it point to value 3
			*/
			fmt.Println(i)
		}()
	}

	fs := [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i:", i) // defer here is registered but not run until main func return
		defer func() {
			fmt.Println("defer closure i:", i)
		}()
		fs[i] = func() {
			fmt.Println("running func i:", i)
		}
	}
	for _, v := range fs {
		v()
	}
	/*
		output:
			running func i: 4
			running func i: 4
			running func i: 4
			running func i: 4
			defer closure i: 4
			defer i: 3
			defer closure i: 4
			defer i: 2
			defer closure i: 4
			defer i: 1
			defer closure i: 4
			defer i: 0
	*/

	// panic and recover
	aa()
	bb()
	cc()

	//struct
	struct1 := &person{} // good habit to use &certainStruct
	struct1.Age = 10     // struct1 is a pointer pointing to a heap address, and don't have to use *struct1.Age
	struct1.Name = "czy"

	struct2 := person{
		Name: "czy",
		Age:  10, // even the last one should have comma
	}
	fmt.Println(struct2)

	learnStruct(struct2) // path value to function, so struct is value pass by default
	fmt.Println(struct2)
	learnStruct1(&struct2)
	fmt.Println(struct2)

	struct3 := &struct { // annonymous struct
		Name string
		Age  int
	}{
		Name: "test",
		Age:  123,
	}
	fmt.Println(struct3)

	// struct4 := &student{
	// 	Name:          "czy",
	// 	Contact.Phone: "123",
	// 	Contact.City:  "Beijing",
	// }

	// struct in another struct could only be inilizalied by this method, method above is not allowed
	struct4 := &student{
		Name: "czy",
	}
	struct4.Contact.City = "Beijing"
	struct4.Contact.Phone = "123"
	fmt.Println(struct4)

	struct5 := &man{ // go way to make struct man inherit from human
		Name:  "czy",
		human: human{1}, // hunam in man is a annonymous fields, so its name is human by default
	}
	fmt.Println(struct5.Sex) // it could directly use attributes in human without call struct5.human.Sex

	// method
	student1 := student{Name: "ads"}
	student1.Print()
	var asdf ASDF
	asdf.Print()         // method value
	(*ASDF).Print(&asdf) // method expression

	// interface
	/*
		structual typing if a struct applied all the method from an interface, it automatically implement this interface,
		and no need to obviously announce it
	*/

	var usb USB
	usb = PhoneConnecter{"iphone"} // PhoneConnector implements USB
	fmt.Println(usb.Name())
	usb.Connect()
	Disconnect(usb)
	LearnInterface(usb)

	pc := PhoneConnecter{"name 1"}
	var usbsample USB
	usbsample = USB(pc) // only pass a copy to the interface, no pointer operations here
	usbsample.Connect()
	pc.name = "brand new name" // this would not change the usbsample's attribute
	usbsample.Connect()

	// reflect
	user := User{"the way to go", 0, 0}
	info(user)

	manager := Manager{User: User{"the way to go", 0, 0}, Title: "manager"}

	tm := reflect.TypeOf(manager)
	fmt.Printf("%#v\n", tm.Field(0)) // User is an anonymous field for manager
	fmt.Printf("%#v\n", tm.Field(1))

	fmt.Printf("%#v\n", tm.FieldByIndex([]int{0, 0})) // get the first fields in User, which is also first for Manager

	intVal := 123
	val := reflect.ValueOf(&intVal)
	val.Elem().SetInt(999)
	fmt.Println(intVal)

	Set(&user)
	fmt.Println(user)

	// concurrency!!!!!
	channel := make(chan bool)
	go func() {
		fmt.Println("Go Go Go")
		channel <- true
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}

	/*
		this will not print something for sure, because channel with cache will NOT wait for someone to read from it
		so the main function would quit after putting true to the channel

		annonymous func and cc <- true would run at the same time, so if cc <- finish first, the print will disappear
	*/
	cc := make(chan bool) // channel with size, is channel with cache
	go func() {
		fmt.Println("channel with cache is async, without cache is sync")
		<-cc
	}()
	cc <- true

	runtime.GOMAXPROCS(runtime.NumCPU())
	ccc := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(ii int) {
			fmt.Println(ii + 10000000000)
			ccc <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-ccc // ensure there is value for 10 times, which means the announymous func above put 10 values into channel
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup, ii int) {
			fmt.Println(9990000 + ii)
			wg.Done()
		}(&wg, i)
	}

	wg.Wait()

	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1: // even if c1 is closed, it will still read from c1, getting default 0 and false
				if !ok {
					o <- true
					break
				} else {
					fmt.Println("c1", v)
				}
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				} else {
					fmt.Println("c2", v)
				}
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 2
	c2 <- "hello"

	close(c1) // when c1 close, the select struct would get ok is false, so there would be a value put to channel o
	// and main func would output o, then the main function exit
	close(c2) // even if u dont close c2, the whole function would also run properly
	// also its useless to ensure two channel are both closed
	<-o
}

// function
func learn(a int, b string) (res1 int, res2 string) { // if no return value, just left it empty
	res1, res2 = 1, "asdf"
	return // if res1, res2 already named in return type above, there is no need to write it in return
}

func learn2(a, b int) (int, int) {
	res1, res2 := 1, 2
	return res1, res2 // if return type is not defined above, return must contain values
}

// if there is a variable-length input parameters, it should locate at the last position
func learn3(a ...int) { // it will make a to slice to receive variable-length input parameters
	a[0] = 100 // note: the variable passed in is copy (or to say passed the value), so changing in the function would not change the value in main
	/*
		if we take a slice as input parameters rather than changing the variable-length parameters into slice,
		if passed its reference(address), so change in function would influence itself
	*/
	fmt.Println(a) // note: variable length parameters like a in this case, should locate in the last position
}

func learn4(a int, b *int) {
	a = 111
	*b = 111
	fmt.Println("changed value", a, *b)
}

func closure(x int) func(y int) int { // return an anonymous function
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

func aa() {
	fmt.Println("func A")
}

func bb() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("B is panic")
}
func cc() {
	fmt.Println("func C")
}

//struct
type person struct {
	Name string // case sensitive
	Age  int    // uppercase means public fields, lowercase means private fields
	// Note: private and public based on package, in the same package, all fields could be visited
}

type student struct {
	Name    string
	Contact struct {
		Phone, City string
	}
}

type human struct {
	Sex int
}

type man struct {
	human
	Name string
}

func learnStruct(p person) {
	p.Age = 13
	fmt.Println(p)
}

func learnStruct1(p *person) {
	p.Age = 13
	fmt.Println(p)
}

// method
func (s student) Print() { // based on the receiver, method would automatically match to certain struct
	fmt.Println(s.Name)
}

// Print go could band method to any type
func (asdf ASDF) Print() {
	fmt.Println("ASDF")
}

// interface

// USB is a method to learn interface
type USB interface {
	Name() string
	Connect()
}

// PhoneConnecter is a struct for USB
type PhoneConnecter struct {
	name string
}

// Name for PhoneConnecter to implement USB
func (pc PhoneConnecter) Name() string {
	return pc.name
}

// Connect for PhoneConnector to implement USB
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect: ", pc.name)
}

// Disconnect requires an USB interface
func Disconnect(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("its a ", pc.name, " Connector")
		return
	}
	fmt.Println("Unknown Connector")
}

// LearnInterface interface{} kindda like a base interface for everything
func LearnInterface(item interface{}) {
	switch t := item.(type) {
	case PhoneConnecter:
		fmt.Println("its a PhoneConnector", t.name)
	default:
		fmt.Println("Unknown type")
	}
}

// User for reflect
type User struct {
	Name string
	ID   int
	Age  int
}

// Manager for reflect
type Manager struct {
	User
	Title string
}

// Hello for reflect
func (user User) Hello() {
	fmt.Println("hello", user.Name)
}

func info(o interface{}) { // pointer parameter is NOT working here
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("only value passed here")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for index := 0; index < t.NumField(); index++ {
		f := t.Field(index)
		val := v.Field(index).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
}

// Set used to change struct
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("this could not be set")
		return
	}

	v = v.Elem()

	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("BYEbye")
	}
}

// concurrency

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
