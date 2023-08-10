package main

func main() {
	var p1 person
	p1.id = 6004521
	p1.name = "nithish"
	p1.email = "sainithish79@gmail.com"
	p1.mobile = 8688582551

	P2 := person{id: 1, name: "sai", email: "gmail", mobile: 890, Addr: Address{line1: "4-367", street: "asd", state: "ap", country: "india", pincode: "asds"}}
	P3 := person{id: 1, name: "sai", email: "gmail", mobile: 890, Addr: Address{line1: "4-367", street: "asd", state: "ap", country: "india", pincode: "asds"}}

	as1:= struct{
		name,emailstring
	}
	{
		name: "Sai"
		mail:
	}
}

type person struct {
	id     int
	name   string
	email  string
	mobile int
	Addr   Address
}
type Address struct {
	line1, city, street, state, country, pincode string
}
type employee struct {
	id     int
	name   string
	email  string
	mobile int
	Address
}
