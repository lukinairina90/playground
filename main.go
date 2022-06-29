package main

import (
	"fmt"
	"sort"
)

func main() {
	//var mobile []int = []int{3, 2, 5, 1, 4}
	//sort.Slice(mobile, func(i, j int) bool { return mobile[i] < mobile[j] })
	//fmt.Println("\n\n######## Sort By Brand [ascending] ###########\n")
	//for _, v := range mobile {
	//	fmt.Println(v)
	//}

	//arr := IntSlice{1, 4, 6, 3, -1, -11}
	//sort.Sort(arr)
	//
	//fmt.Println(arr)

	// ---------------------------------------
	//persons := PersonCollection{
	//	{
	//		Name: "Valera",
	//		Age:  32,
	//	},
	//	{
	//		Name: "Ira",
	//		Age:  18,
	//	},
	//	{
	//		Name: "Petya",
	//		Age:  44,
	//	},
	//	{
	//		Name: "Vasya",
	//		Age:  25,
	//	},
	//	{
	//		Name: "Gena",
	//		Age:  28,
	//	},
	//}
	//
	//persons.SortByAge()
	//
	//for _, person := range persons {
	//	fmt.Printf("%s: %d\n", person.Name, person.Age)
	//}
	//
	//fmt.Println()
	//
	//persons.SortByName()
	//
	//for _, person := range persons {
	//	fmt.Printf("%s: %d\n", person.Name, person.Age)
	//}
	// ---------------------------------------

	frosya := Cat{Age: 4}
	rex := Dog{Age: 6}
	irka := Human{Age: 31}
	kesha := Parrot{Age: 84}

	fmt.Println()
	fmt.Println()

	agerCollection := []Ager{
		frosya,
		rex,
		irka,
		kesha,
	}

	PrintAgeCollection(agerCollection)
	PrintAVGAge(agerCollection)

	humanAgerCollection := []HumanAger{
		frosya,
		rex,
		irka,
		kesha,
	}

	PrintHumanAgeCollection(humanAgerCollection)
}

type IntSlice []int

func (intSlice IntSlice) Len() int {
	return len(intSlice)
}

func (intSlice IntSlice) Less(i, j int) bool {
	return intSlice[i] < intSlice[j]
}

func (intSlice IntSlice) Swap(i, j int) {
	intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
}

type Person struct {
	Name string
	Age  int
}

type PersonCollection []Person

func (p PersonCollection) SortByAge() {
	sort.Sort(PersonsCollectionAgeSorter(p))
}

func (p PersonCollection) SortByName() {
	sort.Sort(PersonsCollectionNameSorter(p))
}

type PersonsCollectionAgeSorter PersonCollection

func (p PersonsCollectionAgeSorter) Len() int {
	return len(p)
}

func (p PersonsCollectionAgeSorter) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p PersonsCollectionAgeSorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type PersonsCollectionNameSorter PersonCollection

func (p PersonsCollectionNameSorter) Len() int {
	return len(p)
}

func (p PersonsCollectionNameSorter) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p PersonsCollectionNameSorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Cat struct {
	Age uint
}

func (c Cat) GetHumanAge() uint {
	return c.Age * 8
}

func (c Cat) GetAge() uint {
	return c.Age
}

type Dog struct {
	Age uint
}

func (d Dog) GetHumanAge() uint {
	return d.Age * 6
}

func (d Dog) GetAge() uint {
	return d.Age
}

type Human struct {
	Age uint
}

func (h Human) GetHumanAge() uint {
	return h.Age
}

func (h Human) GetAge() uint {
	return h.Age
}

type Parrot struct {
	Age uint
}

func (p Parrot) GetHumanAge() uint {
	return p.Age * 12
}

func (p Parrot) GetAge() uint {
	return p.Age
}

type Ager interface {
	GetAge() uint
}

type HumanAger interface {
	GetHumanAge() uint
}

func PrintAgeCollection(a []Ager) {
	for _, ager := range a {
		fmt.Printf("Age: %d\n", ager.GetAge())
	}
}

func PrintHumanAgeCollection(a []HumanAger) {
	for _, ager := range a {
		fmt.Printf("Human age: %d\n", ager.GetHumanAge())
	}
}

func PrintAVGAge(a []Ager) {
	total := uint(0)
	for _, ager := range a {
		total += ager.GetAge()
	}

	fmt.Println(fmt.Sprintf("Average age is: %d", total/uint(len(a))))
}
