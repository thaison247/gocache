package my_gocache

import (
	"testing"
)

// test redis Connect()
func TestConnect(t *testing.T) {

}

// test redis Set()
func TestSet(t *testing.T) {
	// connect to cache server
	var red ICache = Redis{Host: "35.247.157.146", Port: "16379", Password: "scte1234"}
	red.Connect()

	// set with valid key and value
	err := red.Set("city", "Ho Chi Minh city")
	if err != nil {
		t.Error(`Set key: "city" and value: "Ho Chi Minh city" fail!`)
		return
	}

	// get value by key
	val, err1 := red.Get("country")
	if err1 != nil {
		t.Errorf("Error: Get value by Get(\"country\") fail, expected %v, get %v", "Vietnam", val)
		return
	} else {
		t.Logf("Get value by Get(\"country\") success, expected %v, get %v", "Vietnam", val)
	}

	// delete by key
	err2 := red.Delete("country")
	if err2 != nil {
		t.Errorf("Error: Delete by Delete(\"country\") fail!")
	} else {
		t.Logf("Delete by Delete(\"country\") success")
	}

}
