package math

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func Test_ld_arg_minus(t *testing.T) {
	res := Ld(-10)
	if res != 0 {
		t.Error("返り値が不正: 正常値 0 - 実際の返り値 ", res)
	}
}

func Test_ld_arg_zero(t *testing.T) {
	res := Ld(0)
	if res != 0 {
		t.Error("返り値が不正: 正常値 0 - 実際の返り値 ", res)
	}
}

func Test_ld_arg_one(t *testing.T) {
	res := Ld(1)
	if res != 1 {
		t.Error("返り値が不正: 正常値 1 - 実際の返り値 ", res)
	}
}

func Test_ld_arg_two(t *testing.T) {
	res := Ld(2)
	if res != 2 {
		t.Error("返り値が不正: 正常値 2 - 実際の返り値 ", res)
	}
}

func Test_ld_arg_three_five_seven(t *testing.T) {
	x := []int{3, 5, 7}
	for _, v := range x {
		res := Ld(v)
		if res != v {
			t.Error("返り値が不正: 正常値", v, " - 実際の返り値 ", res)
		}
	}
}

func Test_ld_arg_four_six_eight(t *testing.T) {
	x := []int{4, 6, 8}
	for _, v := range x {
		res := Ld(v)
		if res != 2 {
			t.Error("返り値が不正: 正常値 2 - 実際の返り値 ", res)
		}
	}
}

func Test_ld_arg_greater_even_number(t *testing.T) {
	res := Ld(math.MaxInt - 1)
	if res != 2 {
		t.Error("返り値が不正: 正常値 2 - 実際の返り値 ", res)
	}
}
func Test_ld_arg_nine_to_ninenine(t *testing.T) {
	n := []int{3, 2, 11, 2, 13, 2, 3, 2, 17, 2, 19, 2, 3, 2, 23, 2, 5, 2, 3, 2,
		29, 2, 31, 2, 3, 2, 5, 2, 37, 2, 3, 2, 41, 2, 43, 2, 3, 2, 47, 2,
		7, 2, 3, 2, 53, 2, 5, 2, 3, 2, 59, 2, 61, 2, 3, 2, 5, 2, 67, 2,
		3, 2, 71, 2, 73, 2, 3, 2, 7, 2, 79, 2, 3, 2, 83, 2, 5, 2, 3, 2,
		89, 2, 7, 2, 3, 2, 5, 2, 97, 2, 3}
	i := 9
	for _, v := range n {
		res := Ld(i)
		if res != v {
			t.Error("返り値が不正: 正常値 ", v, " - 実際の返り値 ", res)
		}
		i++
	}
}

func Test_ld_arg_big_prime(t *testing.T) {
	var x int
	if strconv.IntSize == 64 {
		x = 1000000005721
	} else {
		x = 2147483647
	}
	res := Ld(x)
	if res != x {
		t.Error("返り値が不正: 正常値 ", x, " - 実際の返り値 ", res)
	}
}

func Test_ld_arg_big_composite(t *testing.T) {
	var x, y int
	if strconv.IntSize == 64 {
		x = 9999823000778041
		y = 99999043
	} else {
		x = 2146654199
		y = 46327
	}
	res := Ld(x)
	if res != y {
		t.Error("返り値が不正: 正常値 ", y, " - 実際の返り値 ", res)
	}
}

func Test_gcd_argzero_1(t *testing.T) {
	res, err := Gcd(0, 1)
	if res != 0 {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_gcd_argzero_2(t *testing.T) {
	res, err := Gcd(1, 0)
	if res != 0 {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_gcd_same_args(t *testing.T) {
	arg := 10
	res, err := Gcd(arg, arg)
	if res != arg {
		t.Error("返り値が不正: 正常値 ", arg, " - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_gcd_arg_big_small(t *testing.T) {
	big, small, correct := 120, 100, 20
	res, err := Gcd(big, small)
	if res != correct {
		t.Error("返り値が不正: 正常値 ", correct, " - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_gcd_arg_small_big(t *testing.T) {
	big, small, correct := 120, 100, 20
	res, err := Gcd(small, big)
	if res != correct {
		t.Error("返り値が不正: 正常値 ", correct, " - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_gcd_arg_small_equal_correct(t *testing.T) {
	big, small := 100, 10
	res, err := Gcd(big, small)
	if res != small {
		t.Error("返り値が不正: 正常値 ", small, " - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_gcd_arg_big_primes(t *testing.T) {
	big, small := 6467477, 3679061
	if Ld(big) != big {
		t.Error("テスト値が不正: 素数でない ", big)
	}
	if Ld(small) != small {
		t.Error("テスト値が不正: 素数でない ", small)
	}
	res, err := Gcd(big, small)
	if res != 1 {
		t.Error("返り値が不正: 正常値 1 - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_adl_argzero(t *testing.T) {
	res, err := Adl(0)
	if res != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_adl_argone(t *testing.T) {
	res, err := Adl(1)
	if len(res) != 1 || res[0] != 1 {
		t.Error("返り値が不正: 正常値 []int{1} - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_adl_argproductofprimes(t *testing.T) {
	data := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	x := 1
	for _, v := range data {
		x = x * v
	}
	res, err := Adl(x)
	if len(res) != len(data) || !reflect.DeepEqual(res, data) {
		t.Error(
			"返り値が不正: 正常値 ",
			data,
			"要素数",
			len(data),
			"- 実際の返り値 ",
			res,
			"要素数",
			len(res),
		)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_adl_argmaxhalf(t *testing.T) {
	res, err := Adl(math.MaxInt/2 + 1)
	if len(res) != strconv.IntSize-1 || res[strconv.IntSize-2] != 2 {
		t.Error(
			"返り値が不正: 正常値 []int{1, 2, ... 2} 要素数",
			strconv.IntSize-1,
			"- 実際の返り値 ",
			res,
			"要素数",
			len(res),
		)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_lcm_argzero_1(t *testing.T) {
	res, err := Lcm(0, 1)
	if res != 0 {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_lcm_argzero_2(t *testing.T) {
	res, err := Lcm(1, 0)
	if res != 0 {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_lcm_big_1(t *testing.T) {
	var x, y, z int
	if strconv.IntSize == 64 {
		x = 577545073
		y = 15969960559
		z = 1317624576693539401
	} else {
		x = 28237
		y = 305591
		z = 784452097
	}

	res, err := Lcm(x, y)
	if res != z {
		t.Error("返り値が不正: 正常値 ", z, " - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_lcm_big_2(t *testing.T) {
	var x, y int
	if strconv.IntSize == 64 {
		x = 3037000507
		y = 3037000537
	} else {
		x = 46349
		y = 46351
	}

	res, err := Lcm(x, y)
	if res != 0 {
		t.Error("返り値が不正: 正常値 0 - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_imos_row_one(t *testing.T) {
	res, err := Imos(10, [][]int{{3, 7, 100}})
	if len(res) != 10 || !reflect.DeepEqual(res, []int{0, 0, 0, 100, 100, 100, 100, 100, 0, 0}) {
		t.Error("返り値が不正: 正常値 []int{0, 0, 0, 100, 100, 100, 100, 100, 0, 0} - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_imos_row_five(t *testing.T) {
	data := [][]int{
		{3, 7, 100},
		{0, 4, 50},
		{6, 9, 10},
		{5, 5, 30},
		{8, 9, 20},
	}
	res, err := Imos(10, data)
	if len(res) != 10 || !reflect.DeepEqual(res, []int{50, 50, 50, 150, 150, 130, 110, 110, 30, 30}) {
		t.Error("返り値が不正: 正常値 []int{50, 50, 50, 150, 150, 130, 110, 110, 30, 30} - 実際の返り値 ", res)
	}
	if err != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", err)
	}
}

func Test_imos_row_short(t *testing.T) {
	data := [][]int{
		{3, 7, 100},
		{0, 4, 50},
		{6, 9},
		{5, 5, 30},
		{8, 9, 20},
	}
	res, err := Imos(10, data)
	if res != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_imos_row_badindex(t *testing.T) {
	data := [][]int{
		{3, 7, 100},
		{0, 4, 50},
		{6, 9, 10},
		{6, 5, 30},
		{8, 9, 20},
	}
	res, err := Imos(10, data)
	if res != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}

func Test_imos_row_badlength(t *testing.T) {
	data := [][]int{
		{3, 7, 100},
		{0, 4, 50},
		{6, 9, 10},
		{5, 5, 30},
		{8, 9, 20},
	}
	res, err := Imos(9, data)
	if res != nil {
		t.Error("返り値が不正: 正常値 nil - 実際の返り値 ", res)
	}
	if err == nil {
		t.Error("返り値が不正: 正常値 error - 実際の返り値 ", err)
	}
}
