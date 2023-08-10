package main

func main() {
	var p1 person
	p1.id = 6004521
	p1.name = "nithish"
	p1.email = "sainithish79@gmail.com"
	p1.mobile = 8688582551

	 P2:=person{id:1,name:"sai",email:"gmail",mobile:890}
	 
}

type person struct {
	id     int
	name   string
	email  string
	mobile int
}
