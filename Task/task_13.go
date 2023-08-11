package main

import (
	"errors"
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
	var mymap map[string]string
	mymap = make(map[string]string)
	mymap["1234"] = "Nithish"
	mymap["1233"] = "sai"
	mymap["1232"] = "Boya"
	println(delete(mymap, "12345"))
	println(delete2(mymap, "1234", "Sai Nithish"))
}
