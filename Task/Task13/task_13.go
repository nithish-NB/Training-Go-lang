package main

import (
	"errors"
	"fmt"
)

func main() {
	var mymap map[string]string
	mymap = make(map[string]string)
	mymap["1234"] = "Nithish"
	mymap["1233"] = "sai"
	mymap["1232"] = "Boya"

	println(Update(mymap, "1234", "Sai Nithish"))
	val, err := Delete(mymap, "1233")
	if err != nil {
		fmt.Println(err)
		fmt.Println(mymap)
	} else if val {
		fmt.Println("Record deleted succesfully!!")
		fmt.Println(mymap)
	}

	val, err = Update(mymap, "1234", "Apple")
	if err != nil {
		fmt.Println(err)
		fmt.Println(mymap)
	} else if val {
		fmt.Println("Record updatd succesfully!!")
		fmt.Println(mymap)
	}
}

func Delete(hash map[string]string, key string) (bool, error) {
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
func Update(hash map[string]string, key string, val string) (bool, error) {
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
