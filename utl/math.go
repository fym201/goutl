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

import "math"

// PowInt is int type of math.Pow function.
func PowInt(x int, y int) int {
	num := 1
	for i := 0; i < y; i++ {
		num *= x
	}
	return num
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func RoundFloor(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc(f*pow10_n) / pow10_n
}

func RoundCeil(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

//给定一个数组和数字，获取数组内若干个元素使得其和等于或最接近于给定数字
func GetSumNum(keys []float64, kill float64) []float64 {
	for _, i := range keys {
		if i == kill {
			return []float64{i}
		}
	}

	ln := len(keys)
	nbit := 1 << uint64(ln)

	var inNear float64 //最接近的数
	var near []float64 //最接近的数对应列表
	for i := 0; i < nbit; i++ {
		var in float64
		var ls []float64
		for j := 0; j < ln; j++ {
			tmp := 1 << uint64(j) // 由0到n右移位
			if tmp&i != 0 { // 与运算，同为1时才会是1
				in += keys[j]
				ls = append(ls, keys[j])
			}
		}
		if in == kill {
			return ls
		} else if in > inNear && in < kill {
			near = ls
		}
	}
	return near
}
