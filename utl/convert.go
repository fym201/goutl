// Copyright 2014 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package utl

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"reflect"
	"strings"
)

// Convert string to specify type.
type Str string

func (f Str) Exist() bool {
	return string(f) != string(0x1E)
}

func (f Str) Uint8() (uint8, error) {
	if strings.LastIndex(string(f), ".") != -1 {
		ff, err := f.Float32()
		return uint8(ff), err
	}
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

func (f Str) Int() (int, error) {
	if strings.LastIndex(string(f), ".") != -1 {
		ff, err := f.Float32()
		return int(ff), err
	}
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

func (f Str) Int64() (int64, error) {
	if strings.LastIndex(string(f), ".") != -1 {
		ff, err := f.Float32()
		return int64(ff), err
	}
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

func (f Str) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

func (f Str) Float64() (float64, error) {
	v, err := strconv.ParseFloat(f.String(), 64)
	return v, err
}

func (f Str) MustUint8() uint8 {
	v, _ := f.Uint8()
	return v
}

func (f Str) MustBool() bool {
	bf := strings.ToLower(f.String())
	if bf == "true" {
		return true
	} else if bf == "false" {
		return false
	}

	v, _ := f.Uint8()
	return v > 0
}

func (f Str) MustInt() int {
	v, _ := f.Int()
	return v
}

func (f Str) MustUInt() uint {
	v, _ := f.Int64()
	return uint(v)
}

func (f Str) MustInt64() int64 {
	v, _ := f.Int64()
	return v
}

func (f Str) MustUint64() uint64 {
	v, _ := strconv.ParseUint(string(f), 10, 64)
	return v
}

func (f Str) MustFloat32() float32 {
	v, _ := f.Float32()
	return v
}

func (f Str) MustFloat64() float64 {
	v, _ := f.Float64()
	return v
}

func (f Str) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

type StrArr []string

func (sa StrArr) MustIntArr() []int {
	ret := make([]int, len(sa))
	for i, v := range sa {
		ret[i] = Str(v).MustInt()
	}
	return ret
}

func (sa StrArr) MustInt64Arr() []int64 {
	ret := make([]int64, len(sa))
	for i, v := range sa {
		ret[i] = Str(v).MustInt64()
	}
	return ret
}

func (sa StrArr) MustUIntArr() []uint {
	ret := make([]uint, len(sa))
	for i, v := range sa {
		ret[i] = Str(v).MustUInt()
	}
	return ret
}

func (sa StrArr) MustFloat32Arr() []float32 {
	ret := make([]float32, len(sa))
	for i, v := range sa {
		ret[i] = Str(v).MustFloat32()
	}
	return ret
}

func (sa StrArr) MustFloat64Arr() []float64 {
	ret := make([]float64, len(sa))
	for i, v := range sa {
		ret[i] = Str(v).MustFloat64()
	}
	return ret
}

// Convert any type to string.
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	case nil:
		s = ""
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func ToInt(value interface{}) (int64, error) {
	var ret int64 = 0
	switch v := value.(type) {
	case bool:
		if true {
			ret = 1
		}
	case float32:
		ret = int64(v)
	case float64:
		ret = int64(v)
	case int:
		ret = int64(v)
	case int8:
		ret = int64(v)
	case int16:
		ret = int64(v)
	case int32:
		ret = int64(v)
	case int64:
		ret = v
	case uint:
		ret = int64(v)
	case uint8:
		ret = int64(v)
	case uint16:
		ret = int64(v)
	case uint32:
		ret = int64(v)
	case uint64:
		ret = int64(v)
	case string:
		ret, _ = strconv.ParseInt(v, 10, 64)
	default:
		return 0, errors.New("Can not convert to int")
	}
	return ret, nil
}

func MustToInt(value interface{}) int64 {
	v, _ := ToInt(value)
	return v
}

func ToFloat(value interface{}) (float64, error) {
	var ret float64 = 0
	switch v := value.(type) {
	case bool:
		if true {
			ret = 1
		}
	case float32:
		ret = float64(v)
	case float64:
		ret = v
	case int:
		ret = float64(v)
	case int8:
		ret = float64(v)
	case int16:
		ret = float64(v)
	case int32:
		ret = float64(v)
	case int64:
		ret = float64(v)
	case uint:
		ret = float64(v)
	case uint8:
		ret = float64(v)
	case uint16:
		ret = float64(v)
	case uint32:
		ret = float64(v)
	case uint64:
		ret = float64(v)
	case string:
		ret, _ = strconv.ParseFloat(v, 64)
	default:
		return 0, errors.New("Can not convert to int")
	}
	return ret, nil
}

func ToBool(value interface{}) (bool, error) {

	switch v := value.(type) {
	case bool:
		return v, nil
	case float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v != 0, nil
	case string:
		return v != "", nil
	default:
		vv := reflect.ValueOf(value)
		tp := vv.Type().Kind()
		if tp == reflect.Ptr || tp == reflect.Map || tp == reflect.Slice {
			return !vv.IsNil(), nil
		}
	}
	return false, errors.New(fmt.Sprintf("Can not convert %v to bool", value))
}

func MustToBool(value interface{}) bool {
	v, _ := ToBool(value)
	return v
}

func MustToFloat(value interface{}) float64 {
	v, _ := ToFloat(value)
	return v
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

// HexStr2int converts hex format string to decimal number.
func HexStr2int(hexStr string) (int, error) {
	num := 0
	length := len(hexStr)
	for i := 0; i < length; i++ {
		char := hexStr[length-i-1]
		factor := -1

		switch {
		case char >= '0' && char <= '9':
			factor = int(char) - '0'
		case char >= 'a' && char <= 'f':
			factor = int(char) - 'a' + 10
		default:
			return -1, fmt.Errorf("invalid hex: %s", string(char))
		}

		num += factor * PowInt(16, i)
	}
	return num, nil
}

// Int2HexStr converts decimal number to hex format string.
func Int2HexStr(num int) (hex string) {
	if num == 0 {
		return "0"
	}

	for num > 0 {
		r := num % 16

		c := "?"
		if r >= 0 && r <= 9 {
			c = string(r + '0')
		} else {
			c = string(r + 'a' - 10)
		}
		hex = c + hex
		num = num / 16
	}
	return hex
}

func ToJson(v interface{}) ([]byte, error) {
	d, e := json.Marshal(v)
	if e != nil {
		return nil, e
	}
	return d, e
}

func MustToJson(v interface{}) []byte {
	d, _ := ToJson(v)
	return d
}

func ToJsonString(v interface{}) (string, error) {
	d, e := json.Marshal(v)
	if e != nil {
		return "", e
	}
	return string(d), e
}

func MustToJsonString(v interface{}) string {
	s, _ := ToJsonString(v)
	return s
}

//Convert millisecond time stamp or date time string to time.Time
func ToTime(value interface{}) (time.Time, error) {
	var timestamp int64 = 0
	switch v := value.(type) {
	case time.Time:
		return v, nil
	case int:
		timestamp = int64(v)
	case int32:
		timestamp = int64(v)
	case int64:
		timestamp = v
	case uint:
		timestamp = int64(v)
	case uint32:
		timestamp = int64(v)
	case uint64:
		timestamp = int64(v)
	case string:
		return time.Parse(time.RFC3339Nano, v)
	default:
		return time.Unix(0, 0), errors.New("Can not convert to time")
	}

	sec := int64(timestamp / 1000)
	msec := timestamp - sec
	return time.Unix(sec, msec*int64(time.Millisecond)), nil

}

func MustToTime(val interface{}) time.Time {
	t, _ := ToTime(val)
	return t
}
