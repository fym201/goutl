// Copyright 2013 com authors
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
	"strconv"
	"strings"
	"reflect"
)

// AppendStr appends string to slice with no duplicates.
func AppendStr(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}
	return append(strs, str)
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements and order are both the same.
func CompareSliceStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements are the same, and ignores the order.
func CompareSliceStrU(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				s2 = append(s2[:j], s2[j+1:]...)
				break
			}
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}

// IsSliceContainsStr returns true if the string exists in given slice, ignore case.
func IsSliceContainsStr(sl []string, str string) bool {
	str = strings.ToLower(str)
	for _, s := range sl {
		if strings.ToLower(s) == str {
			return true
		}
	}
	return false
}

// IsSliceContainsInt64 returns true if the int64 exists in given slice.
func IsSliceContainsInt64(sl []int64, i int64) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

func IsSliceContainsFloat64(sl []float64, i float64) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

func IsSliceContainsInt(sl []int, i int) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

func IsSliceContainsItem(slice interface{}, item interface{}) bool {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == item {
				return true
			}
		}
	default:
		return false
	}

	return false
}

func StringSliceToInt64Slice(sl []string) []int64 {
	var ret []int64
	for _, v := range sl {
		i, _ := strconv.ParseInt(v, 10, 64)
		ret = append(ret, i)
	}
	return ret
}

func StringSliceToUint64Slice(sl []string) []uint64 {
	var ret []uint64
	for _, v := range sl {
		i, _ := strconv.ParseInt(v, 10, 64)
		ret = append(ret, uint64(i))
	}
	return ret
}

func Int64SliceToStringSlice(sl []int64) []string {
	var ret []string
	for _, v := range sl {
		ret = append(ret, strconv.FormatInt(v, 10))
	}
	return ret
}

