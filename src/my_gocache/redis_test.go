package my_gocache

import (
	"fmt"
	"testing"
)

// var (
// 	//RConn: Redis connector
// 	RConn ICache
// )

// func setUp() {
// 	var RConn = Redis{
// 		Host:     "35.247.157.146",
// 		Port:     "16379",
// 		Password: "scte134",
// 	}

// 	RConn.Connect()
// }

func TestRedisSet(t *testing.T) {

	// Connect to redis server
	var RConn ICache = Redis{Host: "35.247.157.146", Port: "16379", Password: "scte1234"}
	RConn.Connect()

	// Close connector
	defer RConn.Close()

	type args struct {
		key        string
		value      interface{}
		expireTime []int
	}

	type expectedRes struct {
		value      interface{}
		errorMsg   error
		expireTime int
	}

	type testCase struct {
		name        string
		args        args
		expectedRes expectedRes
	}

	testCases := []testCase{
		{
			name: "TC1:set value for a new key without expire time",
			args: args{
				key:   "country",
				value: "Vietnam",
			},
			expectedRes: expectedRes{
				value:    "Vietnam",
				errorMsg: nil,
			},
		},
		{
			name: "TC2: set value for a existing key without expire time",
			args: args{
				key:   "country",
				value: "Japan",
			},
			expectedRes: expectedRes{
				value:    "Japan",
				errorMsg: nil,
			},
		},
		{
			name: "TC3: set value for a new key with positive expire time",
			args: args{
				key:        "city",
				value:      "Hai Phong City",
				expireTime: []int{20},
			},
			expectedRes: expectedRes{
				value:      "Hai Phong City",
				errorMsg:   nil,
				expireTime: 20,
			},
		},
		{
			name: "TC4:set value for a new key with negative expire time",
			args: args{
				key:        "city_2",
				value:      "Ho Chi Mih City",
				expireTime: []int{-2},
			},
			expectedRes: expectedRes{
				value:      nil,
				errorMsg:   nil,
				expireTime: -2,
			},
		},
	}

	// iterate to execute all tes case
	for index, tc := range testCases {
		fmt.Printf("%d - ", index+1)

		// set key - value - expireTime
		err := RConn.Set(tc.args.key, tc.args.value, tc.args.expireTime...)

		// check set key error
		if err != tc.expectedRes.errorMsg {
			t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
			continue
		}

		// check wrong value
		if val, err := RConn.Get(tc.args.key); val != tc.expectedRes.value {
			t.Errorf("Fail at [%s], expected value = %v, get value = %v\n", tc.name, tc.expectedRes.value, val)
			t.Errorf("Error message: %s\n", err)
			continue
		}

		// check expire time

		// show PASS message if we pass all above check
		fmt.Printf("%s: PASS\n", tc.name)

	}
}
