/*

Create Company struct Id,Name, Website, Telehone, Fax, Address // Address must be a promoted filed

Create Address struct City, Line1, Street, State, Country, PinCode , Map // Map must be a promoted field

Create Map struct Lan, Lat*/

package main

import "fmt"

func main() {
	var mp map[Address]add
	mp = make(map[Address]add)
	obj := Company{Id: 123, Name: "VSCO", Website: "VSCO.COM", Telehone: 98676767789, Fax: 21234}
	obj.add.lat = 123445
	obj.add.lon = 23445
	mp[obj.Address] = obj.add
	a := mp[obj.Address]
	fmt.Println(a)

}

type Company struct {
	Id       int
	Name     string
	Website  string
	Telehone int
	Fax      int
	Address
}

type Address struct {
	city    string
	line1   string
	street  string
	state   string
	country string
	pincode int
	add
}

type add struct {
	lat int
	lon int
}
