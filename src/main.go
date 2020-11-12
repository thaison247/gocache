package main

import (
	"my_gocache"
)

//35.247.157.146 -p 16379 -a scte1234
//rdcli -h 35.247.157.146 -p 16379 -a scte1234

func main() {

	var rd my_gocache.ICache = my_gocache.Redis{Host: "35.247.157.146", Port: "16379", Password: "scte1234"}
	rd.Connect()

	defer rd.Close()

}
