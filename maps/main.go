package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func delete(hash map[string]string, key string) (bool, error) {
	if hash == nil {
		return false, errors.New("map doesnot exsits")
	}
	_, bool := hash[key]
	if bool {
		return true, nil
	} else {
		return false, nil
	}

}
func delete2(hash map[string]string, key string, val string) (bool, error) {
	if hash == nil {
		return false, errors.New("map doesnot exsits")
	}
	_, bool := hash[key]
	if bool {
		hash[key] = val
		return true, nil
	} else {
		return false, nil
	}

}

func main() {
	//arr:=[] int{1,2,3,1,4,5}
	var slices2 []int
	slices2 = make([]int, 2, 10)
	slices2[0] = 1
	fmt.Println(slices2)
	slice1 := make([]int, 100)
	for i, _ := range slice1 {
		slice1[i] = rand.Intn(5)
	}
	var mymap map[string]string
	mymap = make(map[string]string)
	mymap["1234"] = "Nithish"
	mymap["1233"] = "sai"
	mymap["1232"] = "Boya"
	for key, value := range mymap {
		fmt.Println(key, " ", value)
	}
	val, bool := mymap["1234"] //for key exist or not
	if bool {
		fmt.Println("Key Exists", val)
	} else {
		fmt.Println("no")
	}
	println(delete(mymap, "12345"))
	println(delete2(mymap, "1234", "Sai Nithish"))

}
