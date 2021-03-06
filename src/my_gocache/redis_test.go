package my_gocache

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

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
			name: "TC1: set value for a new key without expire time",
			args: args{
				key:   "country",
				value: "Vietnam",
			},
			expectedRes: expectedRes{
				value:      "Vietnam",
				errorMsg:   nil,
				expireTime: -1,
			},
		},
		{
			name: "TC2: set value for a existing key without expire time",
			args: args{
				key:   "country",
				value: "Japan",
			},
			expectedRes: expectedRes{
				value:      "Japan",
				errorMsg:   nil,
				expireTime: -1,
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
			name: "TC4: set value for a new key with negative expire time",
			args: args{
				key:        "city_2",
				value:      "Ho Chi Mih City",
				expireTime: []int{-20},
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

		// check returned error
		if err == nil {
			if tc.expectedRes.errorMsg != nil {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		} else {
			// check returned error (error != nil)
			if err.Error() != tc.expectedRes.errorMsg.Error() {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		}

		// check returned value
		if val, err1 := RConn.Get(tc.args.key); val != tc.expectedRes.value {
			t.Errorf("Fail at [%s], expected value = %v, get value = %v\n", tc.name, tc.expectedRes.value, val)
			t.Errorf("Error message: %s\n", err1)
			continue
		}

		// check expire time
		if remainTime, err2 := RConn.GetRemainLifeTime(tc.args.key); int(remainTime) != tc.expectedRes.expireTime {
			t.Errorf("Fail at [%s], expected expire time = %v, get expire time = %v\n", tc.name, tc.expectedRes.expireTime, remainTime)
			t.Errorf("Error message: %s\n", err2)
			continue
		}

		// show PASS message if we pass all above check
		fmt.Printf("%s: PASS\n", tc.name)
	}
}

func TestRedisGet(t *testing.T) {

	// Connect to redis server
	var RConn ICache = Redis{Host: "35.247.157.146", Port: "16379", Password: "scte1234"}
	RConn.Connect()

	// Close connector
	defer RConn.Close()

	// set keys-values in advanced
	RConn.Set("testKey1", "value1")
	RConn.Set("testKey2", "value2", 1)
	time.Sleep(time.Second * 2) // wait testKey2 to get expired

	type args struct {
		key string
	}

	type expectedRes struct {
		value    interface{}
		errorMsg error
	}

	type testCase struct {
		name        string
		args        args
		expectedRes expectedRes
	}

	testCases := []testCase{
		{
			name: "TC1: get value by a valid key",
			args: args{
				key: "testKey1",
			},
			expectedRes: expectedRes{
				value:    "value1",
				errorMsg: nil,
			},
		},
		{
			name: "TC2: get value by an invalid key",
			args: args{
				key: "testKey0",
			},
			expectedRes: expectedRes{
				value:    nil,
				errorMsg: errors.New("redigo: nil returned"),
			},
		},
		{
			name: "TC3: get value by an empty key",
			args: args{
				key: "",
			},
			expectedRes: expectedRes{
				value:    nil,
				errorMsg: errors.New("redigo: nil returned"),
			},
		},
		{
			name: "TC4: get value by an expired key",
			args: args{
				key: "testKey2",
			},
			expectedRes: expectedRes{
				value:    nil,
				errorMsg: errors.New("redigo: nil returned"),
			},
		},
	}

	// iterate to execute all tes case
	for index, tc := range testCases {
		fmt.Printf("%d - ", index+1)

		// set key - value - expireTime
		val, err := RConn.Get(tc.args.key)

		// check returned value
		if val != tc.expectedRes.value {
			t.Errorf("Fail at [%s], expected value = %v, get value = %v\n", tc.name, tc.expectedRes.value, val)
			t.Errorf("Error message: %s\n", err)
			continue
		}

		// check returned error
		if err == nil {
			if tc.expectedRes.errorMsg != nil {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		} else {
			// check returned error (error != nil)
			if err.Error() != tc.expectedRes.errorMsg.Error() {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		}

		// show PASS message if we pass all above check
		fmt.Printf("%s: PASS\n", tc.name)
	}
}

func TestRedisDelete(t *testing.T) {

	// Connect to redis server
	var RConn ICache = Redis{Host: "35.247.157.146", Port: "16379", Password: "scte1234"}
	RConn.Connect()

	// Close connector
	defer RConn.Close()

	// set keys-values in advanced
	RConn.Set("testKey1", "value1")
	RConn.Set("testKey2", "value2", 1)
	time.Sleep(time.Second * 2) // wait testKey2 get expired

	type args struct {
		key string
	}

	type expectedRes struct {
		numberOfDeletedKey int64
		errorMsg           error
	}

	type testCase struct {
		name        string
		args        args
		expectedRes expectedRes
	}

	testCases := []testCase{
		{
			name: "TC1: delete by a valid key",
			args: args{
				key: "testKey1",
			},
			expectedRes: expectedRes{
				numberOfDeletedKey: 1,
				errorMsg:           nil,
			},
		},
		{
			name: "TC2: delete by an invalid key",
			args: args{
				key: "testKey0",
			},
			expectedRes: expectedRes{
				numberOfDeletedKey: 0,
				errorMsg:           nil,
			},
		},
		{
			name: "TC3: delete by an empty key",
			args: args{
				key: "",
			},
			expectedRes: expectedRes{
				numberOfDeletedKey: 0,
				errorMsg:           nil,
			},
		},
		{
			name: "TC4: get value by an expired key",
			args: args{
				key: "testKey2",
			},
			expectedRes: expectedRes{
				numberOfDeletedKey: 0,
				errorMsg:           nil,
			},
		},
	}

	// iterate to execute all tes case
	for index, tc := range testCases {
		fmt.Printf("%d - ", index+1)

		// set key - value - expireTime
		val, err := RConn.Delete(tc.args.key)

		// check returned value
		if val != tc.expectedRes.numberOfDeletedKey {
			t.Errorf("Fail at [%s], expected number of deleted rows = %v, get number of deleted rows = %v\n", tc.name, tc.expectedRes.numberOfDeletedKey, val)
			t.Errorf("Error message: %s\n", err)
			continue
		}

		// check returned error
		if err == nil {
			if tc.expectedRes.errorMsg != nil {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		} else {
			// check returned error (error != nil)
			if err.Error() != tc.expectedRes.errorMsg.Error() {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		}

		// show PASS message if we pass all above check
		fmt.Printf("%s: PASS\n", tc.name)

		// del testing keys-values
		RConn.Delete("testKey1")
	}
}

func TestRedisGetRemainLifeTime(t *testing.T) {

	// Connect to redis server
	var RConn ICache = Redis{Host: "35.247.157.146", Port: "16379", Password: "scte1234"}
	RConn.Connect()

	// Close connector
	defer RConn.Close()

	// set keys-values in advanced
	RConn.Set("testKey1", "value1", 20)
	RConn.Set("testKey2", "value2", 25)
	RConn.Set("teseKey3", "value3")
	// time.Sleep(time.Second * 2) // wait testKey2 get expired

	type args struct {
		key string
	}

	type expectedRes struct {
		remainTime int64
		errorMsg   error
	}

	type testCase struct {
		name        string
		args        args
		expectedRes expectedRes
	}

	testCases := []testCase{
		{
			name: "TC1: get remaining life time of a valid key",
			args: args{
				key: "testKey1",
			},
			expectedRes: expectedRes{
				remainTime: 20,
				errorMsg:   nil,
			},
		},
		{
			name: "TC2: get remaining life time of an invalid key",
			args: args{
				key: "testKey0",
			},
			expectedRes: expectedRes{
				remainTime: -2,
				errorMsg:   nil,
			},
		},
		{
			name: "TC3: get remaining life time of an empty key",
			args: args{
				key: "",
			},
			expectedRes: expectedRes{
				remainTime: -2,
				errorMsg:   nil,
			},
		},
		{
			name: "TC4: get remaining life time of a key which has no expire time",
			args: args{
				key: "testKey3",
			},
			expectedRes: expectedRes{
				remainTime: -2,
				errorMsg:   nil,
			},
		},
	}

	// iterate to execute all tes case
	for index, tc := range testCases {
		fmt.Printf("%d - ", index+1)

		// set key - value - expireTime
		remainTime, err := RConn.GetRemainLifeTime(tc.args.key)

		// check returned remaining time
		if remainTime != tc.expectedRes.remainTime {
			t.Errorf("Fail at [%s], expected remaining time = %v, get remaining time = %v\n", tc.name, tc.expectedRes.remainTime, remainTime)
			t.Errorf("Error message: %s\n", err)
			continue
		}

		// check returned error
		if err == nil {
			if tc.expectedRes.errorMsg != nil {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		} else {
			// check returned error (error != nil)
			if err.Error() != tc.expectedRes.errorMsg.Error() {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		}

		// show PASS message if we pass all above check
		fmt.Printf("%s: PASS\n", tc.name)

		RConn.Delete("testKey3")
	}
}

func TestRedisExpire(t *testing.T) {

	// Connect to redis server
	var RConn ICache = Redis{Host: "35.247.157.146", Port: "16379", Password: "scte1234"}
	RConn.Connect()

	// Close connector
	defer RConn.Close()

	// set keys-values in advanced
	RConn.Set("testKey1", "value1")
	RConn.Set("testKey2", "value2", 1)
	time.Sleep(time.Second * 2) // wait testKey2 get expired

	type args struct {
		key        string
		expireTime int
	}

	type expectedRes struct {
		numberOfAffectedKey int64
		errorMsg            error
	}

	type testCase struct {
		name        string
		args        args
		expectedRes expectedRes
	}

	testCases := []testCase{
		{
			name: "TC1: set expire time for a valid key",
			args: args{
				key:        "testKey1",
				expireTime: 20,
			},
			expectedRes: expectedRes{
				numberOfAffectedKey: 1,
				errorMsg:            nil,
			},
		},
		{
			name: "TC2: set expire time for an invalid key",
			args: args{
				key:        "testKey0",
				expireTime: 20,
			},
			expectedRes: expectedRes{
				numberOfAffectedKey: 0,
				errorMsg:            nil,
			},
		},
		{
			name: "TC3: set expire time for a empty key",
			args: args{
				key:        "",
				expireTime: 20,
			},
			expectedRes: expectedRes{
				numberOfAffectedKey: 0,
				errorMsg:            nil,
			},
		},
		{
			name: "TC4: set expire time for an expired key",
			args: args{
				key:        "testKey2",
				expireTime: 20,
			},
			expectedRes: expectedRes{
				numberOfAffectedKey: 0,
				errorMsg:            nil,
			},
		},
	}

	// iterate to execute all tes case
	for index, tc := range testCases {
		fmt.Printf("%d - ", index+1)

		// set key - value - expireTime
		returnNumber, err := RConn.Expire(tc.args.key, tc.args.expireTime)

		// check returned value
		if returnNumber != tc.expectedRes.numberOfAffectedKey {
			t.Errorf("Fail at [%s], expected number of affected rows = %v, get number of affected rows = %v\n", tc.name, tc.expectedRes.numberOfAffectedKey, returnNumber)
			t.Errorf("Error message: %s\n", err)
			continue
		}

		// check returned error
		if err == nil {
			if tc.expectedRes.errorMsg != nil {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		} else {
			// check returned error (error != nil)
			if err.Error() != tc.expectedRes.errorMsg.Error() {
				t.Errorf("Fail at [%s], expected error = %v, get error = %v\n", tc.name, tc.expectedRes.errorMsg, err)
				continue
			}
		}

		// check expire time
		if returnNumber != 0 {
			if expireTime, _ := RConn.GetRemainLifeTime(tc.args.key); int(expireTime) != tc.args.expireTime {
				t.Errorf("Fail at [%s], expected expire time = %v, get expire time = %v\n", tc.name, tc.args.expireTime, expireTime)
				continue
			}
		}

		// show PASS message if we pass all above check
		fmt.Printf("%s: PASS\n", tc.name)

		RConn.Delete("testKey1")
	}
}
