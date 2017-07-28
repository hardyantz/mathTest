package main

import (
	"sort"
	"fmt"
)

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

type Test1 interface {
	Ins1(string) string
}

type TestStruct1 struct {
	Name string
	Address string
}

type Test2 interface {
	Ins2(string) string
}

type TestStruct2 struct {
	Job string
	Title string
}

func (ts TestStruct1) Ins1(s string) string {
	x := "Bio:" +
	"\nName : " + ts.Name +
	"\nAddress: " + ts.Address +
	"\nAdditional: " + s
	return x
}

func (ts TestStruct2) Ins2(s string) string {
	x := "Career:" +
	"\nJob: " + ts.Job +
	"\nTitle: " + ts.Title +
	"\nRemarks: " + s
	return x
}

// create method implement struct interface 1
func imp1(test1 Test1)  {
	fmt.Println(test1.Ins1("sedap"))
}

// create method implement struct interface 2
func imp2 (test2 Test2) {
	fmt.Println(test2.Ins2("Killers"))
}

type StrStrong string

type Strong interface {
	As1() string
}

func (s StrStrong) As1() string {
	fmt.Printf("%#v",s)
	fmt.Println()
	return "Test String"
}

func getString(s Strong) {
	fmt.Println(s.As1())
}

func main (){
	s := Sequence{1,2,3,2,1,2}

	b := sort.Interface(s)
	x := sort.Interface(s).Len()
	j := s.String()
	fmt.Println(b, x, j)

	q := sort.IntSlice{4,5,4,1}

	fmt.Println(q)

	// call interface struct 1
	ts := TestStruct1{Name: "Hardy", Address:"Bogor"}
	y := ts.Ins1("Groot")
	fmt.Println(y)

	imp1(TestStruct1{Name:"Buyung", Address:"Jakarta"})

	// call interface struct 2
	ts2 := TestStruct2{Job:"Programmer", Title: "Staff"}
	y2 := ts2.Ins2("Cool")
	fmt.Println(y2)

	imp2(TestStruct2{Job:"Mechanic", Title:"Leader"})

	// implement test string
	getString(StrStrong("Get String"))

}
