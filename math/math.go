// 数学関連
package math

import (
	"errors"
	"math"
	"strconv"
)

// least divisor
// 引数の 2 以上で最小の約数を返す（引数と同じ数が返ってくる場合、引数は素数）。
// 引数は自然数とする。
// 引数が 1 の場合は 1 を返す。
// 引数が（自然数定義から外れる）1 未満の場合は 0 を返す。
func Ld(x int) int {
	// 引数が偶数及び 9 以下。
	if x < 1 {
		return 0
	}
	if x%2 == 0 {
		return 2
	}
	if x%3 == 0 {
		return 3
	}
	if x%5 == 0 {
		return 5
	}
	if x%7 == 0 {
		return 7
	}
	// 引数が 11 以上の奇数は計算量圧縮のため、引数の平方までを探索対象とし、3, 5, 7 の倍数は無視する。
	sqx := int(math.Sqrt(float64(x)))
	for i := 11; i <= sqx; i += 2 {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			continue
		}
		if x%i == 0 {
			return i
		}
	}
	return x
}

// greatest common divisor
// ２つの自然数から最大公約数を求める。
// ゼロ以下の数を引数にするとエラーを返す。
func Gcd(x, y int) (int, error) {

	// ゼロ以下の数が引数。
	if x < 1 || y < 1 {
		return 0, errors.New("引数が1より小さい")
	}

	// ユークリッドの互除法による判定。
	var f func(int, int) int
	f = func(a, b int) int {
		if b == 0 {
			return a
		}
		return f(b, a%b)
	}
	return f(x, y), nil
}

// all divisors list
// 自然数から約数を列挙する。
// 約数は昇順にスライスに格納される。
// ゼロ以下の数を引数にするとエラーを返す。
func Adl(x int) ([]int, error) {

	// ゼロ以下の数が引数。
	if x < 1 {
		return nil, errors.New("引数が1より小さい")
	}

	// Int 以下の数の約数の数は常に IntSize 以下。
	// よって、約数を格納するスライスのキャパシティは IntSize で設定。
	d := make([]int, 1, strconv.IntSize)

	// 1 は常に約数
	d[0] = 1

	// 約数判定を利用して小さい約数から順にスライスへ格納する。
	if x > 1 {
		for {
			i := Ld(x)
			if i == x {
				d = append(d, x)
				break
			} else {
				d = append(d, i)
				x = x / i
			}
		}
	}

	return d, nil
}

// least common multiple
// ２つの自然数から最小公倍数を求める。
// ゼロ以下の数を引数にするとエラーを返す。
// 最小公倍数が math.MaxInt を超える場合はエラーを返す。
func Lcm(x, y int) (int, error) {

	// 引数の最大公約数を求める。ゼロ以下の数を引数にしてエラーとなった場合は、そのエラーを返す
	d, err := Gcd(x, y)
	if err != nil {
		return 0, err
	}

	// 最小公倍数が math.MaxInt を超える場合はエラーを返す。
	if math.MaxInt/(x/d) < y {
		return 0, errors.New("最小公倍数が" + strconv.Itoa(math.MaxInt) + "を超える")
	}

	return x / d * y, nil
}

// imos法による累積和
// 複数の集計対象のうち、ある範囲に同一値を加算することを繰り返す（範囲、加算値は回ごとに異なる）場合の
// 累積和計算高速化の手法。
// 引数は、
// x : 集計対象の数
// y : 以下の3つの要素を持つスライスを要素とする2次元のスライス
//
//	[0] 対象範囲の先頭のインデックス
//	[1] 対象範囲の最後のインデックス
//	[2] 加算値
//
// 対象ごとの集計結果をスライスで返す。
// 引数のスライスに要素数が2以下のスライスが存在する場合はエラーを返す。
// 対象範囲の先頭のインデックス > 対象範囲の最後のインデックス　となる場合はエラーを返す。
// 対象範囲の最後のインデックス >= 集計対象の数　となる場合はエラーを返す。
func Imos(x int, y [][]int) ([]int, error) {
	data := make([]int, x)
	res := make([]int, x)
	for i := range y {
		if len(y[i]) <= 2 {
			return nil, errors.New("[" + strconv.Itoa(i) + "] の要素数が" + strconv.Itoa(len(y[i])))
		}
		if y[i][0] > y[i][1] {
			return nil, errors.New(
				"対象範囲の先頭のインデックス" + strconv.Itoa(y[i][0]) + " > 対象範囲の最後のインデックス" + strconv.Itoa(y[i][1]),
			)
		}
		if y[i][1] >= x {
			return nil, errors.New(
				"対象範囲の最後のインデックス" + strconv.Itoa(y[i][1]) + " >= 集計対象の数" + strconv.Itoa(x),
			)
		}
		data[y[i][0]] += y[i][2]
		if y[i][1] < x-1 {
			data[y[i][1]+1] -= y[i][2]
		}
	}

	var tmp int
	for i := range data {
		tmp += data[i]
		res[i] += tmp
	}
	return res, nil
}
