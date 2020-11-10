package main

import (
	"fmt"
	"log"
	"my_gocache"
)

//35.247.157.146 -p 16379 -a scte1234
//rdcli -h 35.247.157.146 -p 16379 -a scte1234

func main() {

	var rd my_gocache.ICache = my_gocache.Redis{Host: "35.247.157.146", Port: "16379", Password: "scte1234"}
	rd.Connect()

	defer rd.Close()

	// myMap := make(map[string]string)

	// myMap["country"] = "Vietnam"
	// myMap["district"] = "Phu Nhuan"
	// myMap["city"] = "Ho Chi Minh City"

	// err := rd.Set("address", myMap)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dat, err1 := rd.Get("address")
	// if err1 != nil {
	// 	log.Fatal(err1)
	// } else {
	// 	fmt.Println(dat)
	// }

	val, err2 := rd.Expire("address", 20)
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println(val)
	}
}
